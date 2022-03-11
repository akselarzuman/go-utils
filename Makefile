unit-tests:
	go test -v ./...

linter:
	golangci-lint run

vet:
	go vet -v ./...

build:
	go build ./...