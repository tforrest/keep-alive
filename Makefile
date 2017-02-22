all: build-go build-docker
build-docker:
	docker build .
build-go:
	go build
