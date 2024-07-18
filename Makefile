.PHONY: build clean deploy deploy_prod dev

LAMBDA_NAME = snsLambda
GOOS = linux
GOARCH = amd64
PUBLISHER_BINARY_NAME = bootstrap
CONSUMER_BINARY_NAME = bootstrap
PUBLISHER_ZIP_NAME = main.zip
CONSUMER_ZIP_NAME = consumer.zip
PUBLISHER_MAIN_PATH = cmd/app/main.go
CONSUMER_MAIN_PATH = cmd/consumer/main.go
BIN_DIR = bin

build: clean
	mkdir -p $(BIN_DIR)
	export GO111MODULE=on
	# Build publisher
	env GOARCH=$(GOARCH) GOOS=$(GOOS) go build -ldflags="-s -w" -o $(BIN_DIR)/$(PUBLISHER_BINARY_NAME) $(PUBLISHER_MAIN_PATH)
	cd $(BIN_DIR) && zip ../$(PUBLISHER_ZIP_NAME) $(PUBLISHER_BINARY_NAME)
	# Build consumer
	env GOARCH=$(GOARCH) GOOS=$(GOOS) go build -ldflags="-s -w" -o $(BIN_DIR)/$(CONSUMER_BINARY_NAME) $(CONSUMER_MAIN_PATH)
	cd $(BIN_DIR) && zip ../$(CONSUMER_ZIP_NAME) $(CONSUMER_BINARY_NAME)

deploy_prod: build
	serverless deploy --stage prod

dev:
	sam local start-api

clean:
	go clean
	rm -rf $(BIN_DIR) $(PUBLISHER_ZIP_NAME) $(CONSUMER_ZIP_NAME)