VERSION ?= 0.1.0

run-build: clean
	docker run --rm \
	-v "${PWD}":/go/src/myapp -w /go/src/myapp \
	golang:1.8 \
	make build

build:
	go get -v
	go build -v -o ./out/watch-reload

clean:
	rm -f ./out/*

build-image:
	docker build -t dost/watch-reload:$(VERSION) .