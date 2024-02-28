TESTABLE=$$(go list ./...)

run:
	@go run main.go
.PHONY: run

build:
	@go build -o test-ls .
.PHONY: build

# make sure Content-Length is of correct value for given request
init: build
	@echo 'Content-Length: 118\r\n\r\n{"jsonrpc":"2.0","method":"initialize","id":1,"params":{"processId":1,"clientInfo":{"name":"test","version":"0.0.1"}}}' | ./test-ls
	@ rm -f test-ls
.PHONY: init

inited: build
	@echo 'Content-Length: 40\r\n\r\n{"jsonrpc":"2.0","method":"initialized"}' | ./test-ls
	@ rm -f test-ls
.PHONY: inited

test:
	@go test -v -race $(TESTABLE)
.PHONY: test

path: build
	@ln -s -f $$(pwd)/test-ls /usr/local/bin/test-ls
.PHONY: link-path

rm-link:
	@unlink /usr/local/bin/test-ls
.PHONY: rm-link 
