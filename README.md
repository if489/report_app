# Reports app

## Description
**Reports app** lists all the reports with state open. There are 3 buttons for each report. Button **Block** changes the report state to *Blocked* and **Resolve** changes it to *Closed*.

## Requirements
To launch the demo application, Go (>= 1.13) and Docker are required to be installed.

## Setup
1. Run from the root of the project
```bash
make all
```
or alternatively
```bash
make db-init
make db-import
make build-and-run
```

2. Open `localhost:3000` in browser

## Tests
### Unit Tests
```bash
go test ./...
```

## Documentation
The code for the documentation about the endpoints cane be found in [the `schemas/` folder](schemas/) of this repository.