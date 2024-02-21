TESTABLE=$$(go list ./...)

run:
	@go run main.go
.PHONY: run

test:
	@'go test -v -race $(TESTABLE)'
.PHONY: test
