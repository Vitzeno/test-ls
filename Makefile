TESTABLE=$$(go list ./...)

run:
	@go run main.go
.PHONY: run

build:
	@go build -o lspserver .
.PHONY: build

# make sure Content-Length is of correct value for given request
init: build
	@echo 'Content-Length: 51\r\n\r\n{"jsonrpc": "2.0", "method": "initialize", "id": 1}' | ./lspserver
	@ rm -f lspserver
.PHONY: init

test:
	@go test -v -race $(TESTABLE)
.PHONY: test
