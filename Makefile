# Makefile for the Docker image stevesloka/microservice-sample
# MAINTAINER: Steve Sloka <steve@stevesloka.com>

.PHONY: all build container push clean test

TAG ?= 0.0.1
PREFIX ?= stevesloka

all: container

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o _output/bin/microservice --ldflags '-w' ./main.go

container: build
	docker build -t $(PREFIX)/microservice-sample:$(TAG) .

push:
	docker push $(PREFIX)/microservice-sample:$(TAG)

clean:
	rm -f _output

test: clean
	go test $$(go list ./... | grep -v /vendor/)