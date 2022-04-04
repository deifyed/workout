fmt:
	@gofmt -w .

test:
	@go test ./...

check: fmt test
