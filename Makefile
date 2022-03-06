unit-tests:
	go test -v ./...

linter:
	golangci-lint run