# Nest Event Aggregator

A service that aggregates Nest cam activity and provides a GraphQL interface to query this data.

# Configuration
## Environment Variables

`NEST_TOKEN`
- Your Nest access token
For more information on how to generate an access token, see: https://developers.nest.com/documentation/cloud/how-to-auth

`NEST_CAMERA_ID`
- The ID of the Nest Cam you want to monitor

`NEST_DB_HOST`
- Database host

`NEST_DB_NAME`
- Name of the database

`NEST_DB_USER`
- User of the database with write permissions

`NEST_DB_PASSWORD`
- Password of the above user

## Docker Compose
Used to install dependent application services, like a Postgres database. See: https://docs.docker.com/compose/install/

# Development
## Do it all with `make`
By default, this runs:
```
go fmt
go vet
golint
go test
go build
```
See `Makefile` for more targets.

# Test

In another terminal, run:

```bash
docker-compose up
```

Then:

```bash
make test
```

# Build

```bash
make build
```

# Run

```bash
./nest-event-aggregator
```


