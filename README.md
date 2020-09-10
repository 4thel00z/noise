# noise

## Motivation

A simple cross plattform cli for creating noise in bash scripts etc.

## Installation

Via [gobinaries](https://gobinaries.com):
```shell script
curl https://gobinaries.com/4thel00z/noise | sh 
```

Alternatively via github:

```shell script
git clone git@github.com:4thel00z/noise.git 
go build ./...
```

## Usage
```
Usage of ./noise:
  -duration int
    	duration for the noise (default 200)
  -frequency float
    	frequency for the noise (default 440)
```

### Example

```
noise -duration 123 -frequency 123.4
```