# Nest Event Aggregator

A service that aggregates Nest cam activity and provides a GraphQL interface to query this data.

[![Build status](https://ci.appveyor.com/api/projects/status/fwqruxgdp5p41d33/branch/master?svg=true)](https://ci.appveyor.com/project/zherr/nest-event-aggregator/branch/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/zherr/nest-event-aggregator)](https://goreportcard.com/report/github.com/zherr/nest-event-aggregator)
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

## Query

### allEvents
```bash
curl -g 'http://localhost:8080/graphql?query={allEvents{id,has_sound,has_motion}}'
```
```json
{
  "data":{
    "allEvents":[
      {
        "has_motion":true,
        "has_sound":false,
        "id":1
      },
      {
        "has_motion":true,
        "has_sound":false,
        "id":2
      },
      {
        "has_motion":false,
        "has_sound":true,
        "id":3
      }
    ]
  }
}
```

### eventById
```bash
curl -g 'http://localhost:8080/graphql?query={eventById(id:4){start_time,end_time}}'
```
```json
{
  "data":{
    "eventById":{
      "start_time":"2017-12-28T20:37:14Z",
      "end_time":"2017-12-28T20:37:27.152Z"
    }
  }
}
```

### eventsBetween
```bash
curl -g 'http://localhost:8080/graphql?query={eventsBetween(start:"2017-12-28",end:"2017-12-30"){id}}'
```
```json
{
  "data":{
    "eventsBetween":[
      {
        "id":1
      },
      {
        "id":2
      },
      {
        "id":3
      },
      {
        "id":4
      }
    ]
  }
}
```

## Contributing
Have a feature you want to see in this project? Open an issue with a clear description of what you are looking for. Fork the project, write tests and code, submit a pull request, and have it peer reviewed.

See a bug? Open an issue with a clear description of the bug, prefereably with a test that reproduces it. If you have a fix, do the same as above.

## License
[MIT license](https://github.com/zherr/nest-event-aggregator/blob/master/LICENSE)
