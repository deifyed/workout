fmt:
	@gofmt -w .

test:
	@go test ./...

check: fmt test

build:
	go build -o dist/workout .

install:
	cp ./dist/workout ~/.local/bin/workout

clean:
	rm -r ./dist/