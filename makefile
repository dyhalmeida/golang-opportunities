.PHONY: default run build test docs clean
# Variables
APP_NAME=golang-opportunities

# Tasks
default: run-with-docs

run:
	@go run cmd/main.go

run-with-docs:
	@swag init -g ./cmd/main.go -o ./docs
	@go run cmd/main.go

build:
	@go build -o ${APP_NAME} cmd/main.go

test:
	@go test ./ ...

docs:
	@swag init -g ./cmd/main.go -o ./docs

clean:
	@rm -f ${APP_NAME}
	@rm -rf ./docs