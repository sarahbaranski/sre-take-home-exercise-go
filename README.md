# Uptime Status Checker

## Overview
A GoLang App that regularly checks the health of an endpoint

## Prerequisites
- Go v1.24 installed

## Getting Started

To get started with this app, you first need to clone the repository.

Run `go run main.go sample.yaml` to run the app with sample input.

To exit from the app use `Ctrl+C`

## Fixes
- The first thing I identified that needed to be fixed was to ensure that the file input is a `.yaml` file and log and error if the file is not. This is important to have to ensure that the user is passing in the correct file type as input and if they aren't allowing them to know what file type they should be passing in.

- The next fix was to make sure that the endpoint request responded in 500 milliseconds or less. This was accomplished by adding a timeout to the client object.

- Looking at how the domain was being split, I noticed that it was not ignoring ports. In order to fix this, I added logic to parse through the url and remove the port from the domain.

- Thinking about additional checks I might want to have in place, I added a check for the endpoint to verify if it was an empty string and remove it if it was. This way the app will not waste time printing out 0% availability for an endpoint that is not there. I also added a message to let the user know that that endpoint had been removed.

## Changelog

To view the history and recent changes to this repository, see the [CHANGELOG](./CHANGELOG.md)
