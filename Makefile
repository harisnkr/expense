build:
	@go build -o bin/expense

run: build
	@./bin/expense

test:
	@go test -v ./...