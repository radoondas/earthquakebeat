# Earthquakebeat

Welcome to Earthquakebeat.

Earthquakebeat is a beat which periodically pulls data from [USGS earthquake API](https://earthquake.usgs.gov/fdsnws/event/1/). There are 2 api calls done each `Period` which request new and updated earthquakes.

`New` earthquakes call will request data in GeoJSON format and use attribute `starttime` set to `Now-Period`. That means beat will pull data from past X Period of time you define.
Example `https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2019-08-13T09%3A18%3A18`

`Updated` earthquakes does the same, except it uses attribute `updatedafter` to pull last updated data.
Example: `https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&updatedafter=2019-08-13T09%3A18%3A18`

All other attributes are default and earthquakes from all over the world are being pulled.

Note: Beat preserves earthquake original ID to not duplicate data in index.

## Compatibility

- **Elasticsearch**: 9.x
- **Go**: 1.24+
- **Beats framework**: libbeat 9.3.0

## Installation
Download and install appropriate package for your system. Check release [page](https://github.com/radoondas/earthquakebeat/releases) for latest packages.

You also can use Docker image `docker pull radoondas/earthquakebeat:<version>`


## Configuration

To run Earthquakebeat you have to define `Period` for data pull. 5m should be sufficient and beat will pull new and updated earthquakes from last 5 minutes.

```yaml
  period: 5m
```

Optionally define the path to a CA file for TLS verification of the USGS API. If omitted, system CAs are used.

```yaml
  ssl:
    cafile: ""
```

### Elasticsearch output

Configure the output in `earthquakebeat.yml`:
```yaml
output.elasticsearch:
  hosts: ["localhost:9200"]
  username: "elastic"
  password: "changeme"
```


## Run

```
./earthquakebeat -c earthquakebeat.yml -e
```

To debug run with `debug` flag enabled `./earthquakebeat -c earthquakebeat.yml -e -d "*"`

## Visualisations
This is an example of visualisation for measurements.

![Map](docs/images/map01.png)


## Build
If you want to build Earthquakebeat from scratch, follow [build](BUILD.md) documentation.
