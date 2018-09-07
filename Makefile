VERSION ?= 0.1.0
APP_NAME=watch-reload
IMAGE_TAG=dost/$(APP_NAME):$(VERSION)

run-build: clean
	docker run --rm \
	-v "${PWD}":/go/src/$(APP_NAME) -w /go/src/$(APP_NAME) \
	golang:1.9 \
	make build

build:
	go get -v
	go build -v -o ./out/$(APP_NAME)

clean:
	rm -f ./out/*

build-image:
	docker build -t $(IMAGE_TAG) .