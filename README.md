# Uptime Status Checker
A GoLang App that regularly checks the health of an endpoint

## Installation
- Must have Go v1.24.2 installed

run `go run main.go sample.yaml` to run the app with sample input.

## Fixes
- Updated main function to check if config file is a `.yaml` and return an error message if its not.

- Added Timeout to the client to ensure timeout was set 500 milliseconds and tested with a URL that I knew would trigger the timeout

- Parsed through the url to remove any port prior to splitting off the domain