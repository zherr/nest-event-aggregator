all: fmt vet lint test build

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golint ./... | grep -v vendor || true

test:
	go test -v ./...

build:
	go build

run:
	./gocapi
