# .PHONY: build clean deploy gomodgen

build: clean
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -o bin/bootstrap cmd/lambda/main.go
	cd bin && zip bootstrap.zip bootstrap

deploy_prod: build
	serverless deploy --stage prod

dev:
	sam local start-api

clean:
	go clean
	rm -rf ./bin

# deploy: clean build
# 	sls deploy --verbose


# gomodgen:
# 	chmod u+x gomod.sh
# 	./gomod.sh
