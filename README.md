# Nest Event Aggregator

A microservice that aggregates Nest cam activity and provides an HTTP REST interface to query this data.

# Configuration
## Environment Variables

`NEST_TOKEN`
- Your Nest access token
For more information on how to generate an access token, see: https://developers.nest.com/documentation/cloud/how-to-auth

`NEST_CAMERA_ID`
- The ID of the Nest Cam you want to monitor

`NEST_DB_ENDPOINT`
- The fully qualified endpoint of the mysql database.
- Ex: `root:root@tcp(localhost:3306)/nest?parseTime=true`

# Build

```bash
go build
```

# Test
Ensure `nest_test` database exists, then:

```bash
go test
```

# Run

```bash
./nest-event-aggregator
```


