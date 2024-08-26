.PHONY: default run build test docs clean

BUILD_FOLDER=build/run
MAIN=cmd/link.in/main.go

default: run

run:
	@go run $(MAIN)
build:
	@go build -o $(BUILD_FOLDER) $(MAIN)
docker:
	@docker-compose build
	@docker-compose up
test:
	@go test ./test/...
docs:
	@echo "developing"
	# @swag init -g $(MAIN)
clean:
	@rm -f $(BUILD_FOLDER)
	@rm -rf ./docs
install_lint:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.60.3
	@golangci-lint --version
lint:
	@golangci-lint run