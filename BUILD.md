# Earthquakebeat

Welcome to Earthquakebeat.

## Getting Started with Earthquakebeat

### Requirements

* [Golang](https://golang.org/dl/) 1.24+


### Clone

To clone Earthquakebeat from the git repository, run the following commands:

```
git clone https://github.com/radoondas/earthquakebeat
cd earthquakebeat
```


### Build

To build the binary for Earthquakebeat run the command below. This will generate a binary
in the same directory with the name earthquakebeat.

```
go build -o earthquakebeat .
```


### Run

To run Earthquakebeat with debugging output enabled, run:

```
./earthquakebeat -c earthquakebeat.yml -e -d "*"
```


### Test

To test Earthquakebeat, run the following command:

```
go test ./...
```


### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean up the build directory and generated artifacts, run:

```
make clean
```

## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/). To build packages of your beat, run the following command:

```
mage package
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
