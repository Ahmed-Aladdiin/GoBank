build:
	@go build -o bin/gobank
	@./set_container.sh

run: build
	@./bin/gobank

test:
	@go test -v ./...
