# notification
![notification-tests](https://github.com/4thel00z/notification/workflows/Test/badge.svg)
[![GPLv3 license](https://img.shields.io/badge/License-GPLv3-blue.svg)](http://perso.crans.org/besson/LICENSE.html)

## Motivation

A simple cross plattform cli for creating desktop notification.
Useful tool if used in congration with other tools, to show a process has finished executing etc.

## Installation

Via [gobinaries](https://gobinaries.com):
```shell script
curl https://gobinaries.com/4thel00z/notification | sh 
```

Alternatively via github:

```shell script
git clone git@github.com:4thel00z/notification.git 
go build ./...
```

## Usage
```
Usage of ./notification:
  -appIcon string
    	app icon
  -message string
    	notification message
  -title string
    	notification title (default "Notification")
```

### Example

```
notification -message "Process xyz is ready" -title "Ready"
```
## License

This project is licensed under the GPL-3 license.
