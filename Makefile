unit-tests:
	go test -v ./...

unit-tests-coverage:
	go test -coverprofile=coverage.out ./...

linter:
	golangci-lint run

vet:
	go vet -v ./...

build:
	go build ./...