# Nest Event Aggregator

A service that aggregates Nest cam activity and provides a GraphQL interface to query this data.

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


