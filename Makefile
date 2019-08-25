.PHONY: deps clean build

deps:
	go mod tidy

clean: 
	rm -rf airlog

build:
	cd airlog && GOOS=linux GOARCH=amd64 go build -o airlog .
