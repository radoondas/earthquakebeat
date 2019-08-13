# Earthquakebeat

Welcome to Earthquakebeat.

Earthquakebeat is a beat which periodically pulls data from [USGS earthquake API](https://earthquake.usgs.gov/fdsnws/event/1/). There are 2 api calls done eacr `Period` which request new and updated earthquakes. 

`New` earthquakes call will request data in GeoJSON format and use attribute `starttime` set to `Now-Period`. That meas beat will pull data from past X Period of time you define.
Example `https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2019-08-13T09%3A18%3A18
`   

`Updated ` earthquakes does the same, except is uses attribute `updatedafter` to pull last updated data. 
Example: `https://earthquake.usgs.gov/fdsnws/event/1/query?format=geojson&starttime=2019-08-13T09%3A18%3A18`

All other attributes are default and earthquakes from all over the world are being pulled.

Note: Beat preserve earthquake original ID to not to duplicate data in index.


## Installation
Download and install appropriate package for your system. Check release [page](https://github.com/radoondas/earthquakebeat/releases) for latest packages.

You also can use Docker image `docker pull radoondas/earthquakebeat:<version>`


## Configuration

To run Earthquakebeat you have to define `Period` for data pull. 5m should be sufficient and beat will pull new and updated earthquakes from last 5 minutes.

```yaml
  period: 5m
```

Define the path to CA file which requires TLS call. One CA is provided in the repository. Feel free to use it.


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
