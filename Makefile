build:
	@go build -o bin/gobank
	@./set_container.sh
	@sleep 4

setup: build 
	@./bin/gobank

run:
	@go build -o bin/gobank
	@./bin/gobank

test:
	@go test -v ./...
