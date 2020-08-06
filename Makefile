GOOS ?= linux
GOARCH ?= amd64
GIT_COMMIT=$(shell git rev-list -1 HEAD)
GIT_URL=$(shell git config --get remote.origin.url)
GIT_TAG=$(shell git describe --tags $(shell git rev-list --tags --max-count=1))

.PHONY: build run test startam stopam zip clean

build:
	GOARCH=$(GOARCH) GOOS=$(GOOS) go build ./...

run:
	go run ./...

test: build
	go test -v -cover -race

startam:
	docker run -d --name amsilenceschedulertest -p 9093:9093 prom/alertmanager:v0.21.0

stopam:
	docker stop amsilenceschedulertest
	docker rm amsilenceschedulertest

zip: build
	 zip amsilencescheduler-$(GIT_TAG).zip amsilencescheduler

clean:
	go clean
	rm -f amsilencescheduler*.zip
