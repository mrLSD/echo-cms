run:
	@go run server.go

test:
	@./runtest.sh

build:
	@go build

fmt:
	@gofmt -w -l .
