build:
	@go build -o bin/gobank
	@./set_container.sh
	@sleep 4

run: build 
	@./bin/gobank

test:
	@go test -v ./...
