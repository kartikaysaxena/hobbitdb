build:
	@go build -o bin/hobbit cmd/main.go

run: build
	@./bin/hobbit

test:
	@go test -v ./...
