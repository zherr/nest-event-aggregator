# Nest Event Aggregator

A service that aggregates Nest cam activity and provides a GraphQL interface to query this data.

[![go report card](https://goreportcard.com/badge/github.com/zherr/nest-event-aggregator "go report card")](https://goreportcard.com/report/github.com/zherr/nest-event-aggregator)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/zherr/nest-event-aggregator?status.svg)](https://godoc.org/github.com/zherr/nest-event-aggregator)

## Configuration
### Environment Variables
See `.env.example` for application configurations.

### Docker Compose
Used to install dependent application services, like a Postgres database. See: https://docs.docker.com/compose/install/

## Development

###
In development, environment variables will be picked up automatically from `.env`:
```
cp .env.example .env
```

### Do it all with `make`
By default, this runs:
```
go fmt
go vet
golint
go test
go build
```
See `Makefile` for more targets.

## Test
In another terminal, run:
```bash
docker-compose up
```

Then:
```bash
make test
```

## Build
```bash
make build
```

## Run
```bash
./nest-event-aggregator
```

## Contributing
Have a feature you want to see in this project? Open an issue with a clear description of what you are looking for. Fork the project, write tests and code, submit a pull request, and have it peer reviewed.

See a bug? Open an issue with a clear description of the bug, prefereably with a test that reproduces it. If you have a fix, do the same as above.
