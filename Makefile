.PHONY: deps clean build

deps:
	go mod tidy

clean: 
	rm -rf textvid

build:
	cd textvid && GOOS=linux GOARCH=amd64 go build -o textvid .
