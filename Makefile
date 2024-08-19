.PHONY: default run build test docs clean

BUILD_FOLDER=build/run
MAIN=cmd/main.go

default: run

run:
	@go run $(MAIN)
build:
	@go build -o $(BUILD_FOLDER) $(MAIN)
docker:
	@docker-compose build
	@docker-compose up
test:
	@go test ./ ...
docs:
	@echo "developing"
	# @swag init -g $(MAIN)
clean:
	@rm -f $(BUILD_FOLDER)
	@rm -rf ./docs