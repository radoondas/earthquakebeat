package beater

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/radoondas/earthquakebeat/config"
)

// Earthquakebeat configuration.
type Earthquakebeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

type EarthquakeResponse struct {
	Type     string `json:"type"`
	Metadata struct {
		Generated int64  `json:"generated"`
		Url       string `json:"url"`
		Title     string `json:"title"`
		Status    int    `json:"status"`
		Api       string `json:"api"`
		Count     int    `json:"count"`
	}
	Features []Feature
}

//https://earthquake.usgs.gov/earthquakes/feed/v1.0/geojson.php
type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
	Id         string     `json:"id"`
}

//https://earthquake.usgs.gov/data/comcat/data-eventterms.php
type Properties struct {
	Mag     float64 `json:"mag"`
	Place   string  `json:"place"`
	Time    int64   `json:"time"`
	Updated int64   `json:"updated"`
	Tz      int     `json:"tz"`
	Url     string  `json:"url"`
	Detail  string  `json:"detail"`
	Felt    int     `json:"felt"`
	Cdi     float32 `json:"cdi"`
	Mmi     float32 `json:"mmi"`
	Alert   string  `json:"alert"`
	Status  string  `json:"status"`
	Tsunami int     `json:"tsunami"`
	Sig     int     `json:"sig"`
	Net     string  `json:"net"`
	Code    string  `json:"code"`
	Ids     string  `json:"ids"`
	Sources string  `json:"sources"`
	Types   string  `json:"types"`
	Nst     int     `json:"nst"`
	Dmin    float64 `json:"dmin"`
	Rms     float64 `json:"rms"`
	Gap     float64 `json:"gap"`
	MagType string  `json:"magType"`
	Type    string  `json:"type"`
	Title   string  `json:"title"`
}

type Geometry struct {
	Type string `json:"type"`
	// Longitude, Latitude, depth
	Coordinates []float64 `json:"coordinates"`
}

type EQResponse struct {
	tm      time.Time
	success bool
}

const (
	// https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2019-08-04T22:00:00
	earthquakeUrl  = "https://earthquake.usgs.gov"
	uriPath        = "/fdsnws/event/1/query"
	selector       = "earthquake"
	dateTimeFormat = "2006-01-02T15:04:05"
)

// New creates an instance of earthquakebeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	bt := &Earthquakebeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts earthquakebeat.
func (bt *Earthquakebeat) Run(b *beat.Beat) error {
	logp.NewLogger(selector).Info("earthquakebeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	logp.NewLogger(selector).Info("Configuration: ", bt.config)

	go func() {
		ticker := time.NewTicker(bt.config.Period)
		defer ticker.Stop()

		// UTC is required
		lastTimeNew := time.Now().UTC().Add(-bt.config.Period)

		for {

			// new earthquakes
			err, sfr := bt.GetEarthquakes(b, lastTimeNew, bt.config.SSL.CAfile, true)
			if err != nil {
				logp.NewLogger(selector).Error("Error while getting earthquakes data: %v", err)
			} else {
				if sfr.success {
					lastTimeNew = sfr.tm
				}
			}

			select {
			case <-bt.done:
				goto GotoFinish
			case <-ticker.C:
			}
		}
	GotoFinish:
	}()

	go func() {
		ticker := time.NewTicker(bt.config.Period)
		defer ticker.Stop()

		// UTC is required
		lastTimeUpdated := time.Now().UTC().Add(-bt.config.Period)

		for {

			// updated earthquakes
			err, sfr := bt.GetEarthquakes(b, lastTimeUpdated, bt.config.SSL.CAfile, false)
			if err != nil {
				logp.NewLogger(selector).Error("Error while getting earthquakes data: %v", err)
			} else {
				if sfr.success {
					lastTimeUpdated = sfr.tm
				}
			}

			select {
			case <-bt.done:
				goto GotoFinish
			case <-ticker.C:
			}
		}
	GotoFinish:
	}()

	<-bt.done

	return nil
}

// Stop stops earthquakebeat.
func (bt *Earthquakebeat) Stop() {
	bt.client.Close()
	close(bt.done)
}

func (bt *Earthquakebeat) GetEarthquakes(b *beat.Beat, lastRun time.Time, CAFile string, new bool) (error, EQResponse) {

	now := time.Now().UTC()
	eqr := EQResponse{tm: now, success: false}

	tlsConfig := &tls.Config{RootCAs: x509.NewCertPool()}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	var ParsedUrl *url.URL

	if CAFile != "" {
		// Load our trusted certificate path
		pemData, err := ioutil.ReadFile(CAFile)
		if err != nil {
			panic(err)
		}
		ok := tlsConfig.RootCAs.AppendCertsFromPEM(pemData)
		if !ok {
			logp.NewLogger(selector).Error("Unable to load CA file")
			panic("Couldn't load PEM data")
		}
	}

	client := &http.Client{Transport: transport}

	ParsedUrl, err := url.Parse(earthquakeUrl)
	if err != nil {
		logp.NewLogger(selector).Error("Unable to parse URL string")
		panic(err)
	}

	ParsedUrl.Path += uriPath
	parameters := url.Values{}

	parameters.Add("format", "geojson")
	if new {
		parameters.Add("starttime", lastRun.Format(dateTimeFormat))
	} else {
		parameters.Add("updatedafter", lastRun.Format(dateTimeFormat))
	}

	logp.NewLogger(selector).Debug("Last run at: ", lastRun.String())

	ParsedUrl.RawQuery = parameters.Encode()
	logp.NewLogger(selector).Debug("Requesting Earthquake data: ", ParsedUrl.String())

	req, err := http.NewRequest("GET", ParsedUrl.String(), nil)

	res, err := client.Do(req)

	if err != nil {
		return err, eqr
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		logp.NewLogger(selector).Debug("Status code: ", res.StatusCode)
		logp.NewLogger(selector).Debug("Status code: ", res.Body)
		return fmt.Errorf("HTTP %s", res), eqr
	}

	body, err := ioutil.ReadAll(res.Body)
	//logp.NewLogger(selector).Debug(body)
	if err != nil {
		log.Fatal(err)
		return err, eqr
	}

	eqdata := EarthquakeResponse{}
	err = json.Unmarshal(body, &eqdata)
	if err != nil {
		fmt.Printf("error: %v", err)
		return err, eqr
	}

	if eqdata.Metadata.Count > 0 {
		if new {
			logp.NewLogger(selector).Debug(eqdata.Metadata.Count, " new earthquakes detected.")
		} else {
			logp.NewLogger(selector).Debug(eqdata.Metadata.Count, " updated earthquakes detected.")
		}
		//logp.NewLogger(selector).Debug("Unmarshal-ed Earthquake data: ", eqdata)

		transformedData := bt.TransformStationData(eqdata)

		ts := time.Now()
		for _, d := range transformedData {

			event := beat.Event{
				Meta: common.MapStr{
					"id": d["id"].(string),
				},
				Timestamp: ts,
				Fields: common.MapStr{
					"type":       "earthquakebeat",
					"earthquake": d,
				},
			}

			logp.NewLogger(selector).Debug("Event: ", event)
			bt.client.Publish(event)
		}
		eqr.success = true

	} else {
		eqr.success = true
		if new {
			logp.NewLogger(selector).Debug("No new earthquakes.")
		} else {
			logp.NewLogger(selector).Debug("No updated earthquakes.")
		}
	}

	return nil, eqr
}

func (bt *Earthquakebeat) TransformStationData(data EarthquakeResponse) []common.MapStr {

	earthquakeData := []common.MapStr{}

	for _, f := range data.Features {
		eq := common.MapStr{
			"id": f.Id,
			"properties": common.MapStr{
				"mag":     f.Properties.Mag,
				"place":   f.Properties.Place,
				"time":    strconv.FormatInt(f.Properties.Time, 10),
				"updated": strconv.FormatInt(f.Properties.Updated, 10),
				"tz":      f.Properties.Tz,
				"url":     f.Properties.Url,
				"detail":  f.Properties.Detail,
				"felt":    f.Properties.Felt,
				"cdi":     f.Properties.Cdi,
				"mmi":     f.Properties.Mmi,
				"alert":   f.Properties.Alert,
				"status":  f.Properties.Status,
				"tsunami": f.Properties.Tsunami,
				"sig":     f.Properties.Sig,
				"net":     f.Properties.Net,
				"code":    f.Properties.Code,
				"ids":     f.Properties.Ids,
				"sources": f.Properties.Sources,
				"types":   f.Properties.Types,
				"nst":     f.Properties.Nst,
				"dmin":    f.Properties.Dmin,
				"rms":     f.Properties.Rms,
				"gap":     f.Properties.Gap,
				"magType": f.Properties.MagType,
				"type":    f.Properties.Type,
				"title":   f.Properties.Title,
			},
			"geometry": common.MapStr{
				"depth": f.Geometry.Coordinates[2],
				"location": common.MapStr{
					"lat": f.Geometry.Coordinates[1],
					"lon": f.Geometry.Coordinates[0],
				},
			},
		}
		logp.NewLogger(selector).Debug("EQ data: ", eq)

		earthquakeData = append(earthquakeData, eq)
	}

	return earthquakeData
}
