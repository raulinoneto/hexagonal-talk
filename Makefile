.PHONY: build clean deploy

build:
	env GO111MODULE=on GOOS=linux go build -ldflags="-s -w" -o bin/votes cmd/vote/*

test: 
	env GO111MODULE=on go test ./... -cover
	
clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
