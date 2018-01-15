MAKEDIR:=$(strip $(shell dirname "$(realpath $(lastword $(MAKEFILE_LIST)))"))

.PHONY: build
build:
	docker run -v ${GOPATH}/src:/go/src -v ${MAKEDIR}:/go/src/github.com/almariah/ik-mysql golang:1.8 make -C /go/src/github.com/almariah/ik-mysql build-plugin

.PHONY: build-plugin
build-plugin:
	go build -buildmode=plugin -o .build/ik-mysql.so
