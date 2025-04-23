# Uptime Status Checker

## Overview
A GoLang App that regularly checks the health of endpoints and gives a cumulative total percentage per domain.

## Prerequisites
- Go v1.24 installed

## Getting Started

To get started with this app, you first need to clone this repository.

Additionally if you are on a mac and have homebrew installed, you will want to run `brew install go`

You can also install go for windows using this [installation](https://go.dev/doc/install) link

Check to see which version of Go you have installed by running `go version` in terminal.

Run `go run main.go sample.yaml` to run the app with sample input.
or use the go [build](https://go.dev/doc/tutorial/compile-install) command to generate a binary executable.

To exit from the app use `Ctrl+C` in the terminal.

## Fixes
- The first thing I identified that needed to be fixed was to ensure that the file input is a `.yaml` file and log and error if the file is not. This is important to have to ensure that the user is passing in the correct file type as input and if they aren't allowing them to know what file type they should be passing in.
```
sre-take-home-exercise-go git:(main) ✗ go run main.go sample.txt
2025/04/21 16:58:59 Error config is not yaml file
exit status 1
```

- The next fix was to make sure that the endpoint request responded in 500 milliseconds or less. This was accomplished by adding a timeout to the client object.

- Looking at how the domain was being split, I noticed that it was not ignoring ports. In order to fix this, I added logic to parse through the url and remove the port from the domain. Using `www.google.com:81` endpoint to add to the sample.yaml file, I was able to verify that this did remove the port:
```
sre-take-home-exercise-go git:(main) ✗ go run main.go sample.yaml
www.google.com has 0.000000% availability
dev-sre-take-home-exercise-rubric.us-east-1.recruiting-public.fetchrewards.com has 25.000000% availability
```

- Thinking about additional checks I might want to have in place, I added a check for the endpoint to verify if it was an empty string and remove it if it was. This way the app will not waste time printing out 0% availability for an endpoint that is not there. I also added a message to let the user know that that endpoint had been removed.
```
sre-take-home-exercise-go git:(main) ✗ go run main.go sample.yaml
Endpoint 4 does not contain a URL. Removing from list of endpoints to check.
dev-sre-take-home-exercise-rubric.us-east-1.recruiting-public.fetchrewards.com has 25.000000% availability
^Csignal: interrupt
```

## Changelog

To view the history and recent changes to this repository, see the [CHANGELOG](./CHANGELOG.md)

## Additional thoughts
- I'm not totally sure how helpful this app is while only showing cumulative percentages of availability per domain. I would assume seeing logs or metrics displaying each endpoint individually with their availability and status code would give someone more information in order to solve any issues.
- Using a tool such as [Kuma](https://github.com/louislam/uptime-kuma) and having it trigger notifications through Slack or another communication platform and/or a paging system would help the SRE team (and additional teams) know if there was an issue with and endpoint not being reachable.