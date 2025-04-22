# Uptime Status Checker

## Installation
- Must have Go v1.24.2 installed

## Bugs and Fixes
- Updated main function to check if config file is a `.yaml` and return an error message if its not.

- Added Timeout to the client to ensure timeout was set 500 milliseconds and tested with a URL that I knew would trigger the timeout