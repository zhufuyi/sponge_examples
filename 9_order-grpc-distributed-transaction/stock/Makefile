SHELL := /bin/bash

PROJECT_NAME := "stock"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ | grep -v /api/)



.PHONY: mod
# add missing and remove unused modules
mod:
	go mod tidy


.PHONY: fmt
# go format *.go files
fmt:
	gofmt -s -w .


.PHONY: ci-lint
# check the code specification against the rules in the .golangci.yml file
ci-lint: fmt
	golangci-lint run ./...


.PHONY: dep
# download dependencies to the directory vendor
dep:
	go mod download


.PHONY: test
# go test *_test.go files, the parameter -count=1 means that caching is disabled
test:
	go test -count=1 -short ${PKG_LIST}


.PHONY: cover
# generate test coverage
cover:
	go test -short -coverprofile=cover.out -covermode=atomic ${PKG_LIST}
	go tool cover -html=cover.out


.PHONY: graph
# generate interactive visual function dependency graphs
graph:
	@echo "generating graph ......"
	@cp -f cmd/stock/main.go .
	go-callvis -skipbrowser -format=svg -nostd -file=stock stock
	@rm -f main.go stock.gv


.PHONY: docs
# generate swagger docs, only for ⓵ Web services created based on sql, the host address can be changed via parameters, e.g. make docs HOST=192.168.3.37
docs: mod fmt
	@bash scripts/swag-docs.sh $(HOST)


.PHONY: proto
# generate *.go and template code by proto files, if you do not refer to the proto file, the default is all the proto files in the api directory. you can specify the proto file, multiple files are separated by commas, e.g. make proto FILES=api/user/v1/user.proto. only for ⓶ Microservices created based on sql, ⓷ Web services created based on protobuf, ⓸ Microservices created based on protobuf, ⓹ grpc gateway service created based on protobuf
proto: mod fmt
	@bash scripts/protoc.sh $(FILES)


.PHONY: proto-doc
# generate doc from *.proto files
proto-doc:
	@bash scripts/proto-doc.sh


.PHONY: build
# build stock for linux amd64 binary
build:
	@echo "building 'stock', linux binary file will output to 'cmd/stock'"
	@cd cmd/stock && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build



.PHONY: run
# run service
run:
	@bash scripts/run.sh


.PHONY: run-nohup
# run server with nohup to local, if you want to stop the server, pass the parameter stop, e.g. make run-nohup CMD=stop
run-nohup:
	@bash scripts/run-nohup.sh $(CMD)


.PHONY: binary-package
# packaged binary files
binary-package: build
	@bash scripts/binary-package.sh


.PHONY: deploy-binary
# deploy binary to remote linux server, e.g. make deploy-binary USER=root PWD=123456 IP=192.168.1.10
deploy-binary: binary-package
	@expect scripts/deploy-binary.sh $(USER) $(PWD) $(IP)


.PHONY: image-build-local
# build image for local docker, tag=latest, use binary files to build
image-build-local: build
	@bash scripts/image-build-local.sh


.PHONY: deploy-docker
# deploy service to local docker, if you want to update the service, run the make deploy-docker command again.
deploy-docker: image-build-local
	@bash scripts/deploy-docker.sh


.PHONY: image-build
# build image for remote repositories, use binary files to build, e.g. make image-build REPO_HOST=addr TAG=latest
image-build:
	@bash scripts/image-build.sh $(REPO_HOST) $(TAG)


.PHONY: image-build2
# build image for remote repositories, phase II build, e.g. make image-build2 REPO_HOST=addr TAG=latest
image-build2:
	@bash scripts/image-build2.sh $(REPO_HOST) $(TAG)


.PHONY: image-push
# push docker image to remote repositories, e.g. make image-push REPO_HOST=addr TAG=latest
image-push:
	@bash scripts/image-push.sh $(REPO_HOST) $(TAG)


.PHONY: deploy-k8s
# deploy service to k8s
deploy-k8s:
	@bash scripts/deploy-k8s.sh


.PHONY: image-build-rpc-test
# build grpc test image for remote repositories, e.g. make image-build-rpc-test REPO_HOST=addr TAG=latest
image-build-rpc-test:
	@bash scripts/image-rpc-test.sh $(REPO_HOST) $(TAG)


.PHONY: patch
# patch some dependent code, such as types.proto, mysql initialization code. e.g. make patch TYPE=types-pb , make patch TYPE=mysql-init
patch:
	@bash scripts/patch.sh $(TYPE)


.PHONY: copy-proto
# copy proto file from the grpc server directory, multiple directories separated by commas. e.g. make copy-proto SERVER=yourServerDir
copy-proto:
	@sponge patch copy-proto --server-dir=$(SERVER)


.PHONY: update-config
# update internal/config code base on yaml file
update-config:
	@sponge config --server-dir=.


.PHONY: clean
# clean binary file, cover.out, template file
clean:
	@rm -vrf cmd/stock/stock
	@rm -vrf cover.out
	@rm -vrf main.go stock.gv
	@rm -vrf internal/ecode/*.go.gen*
	@rm -vrf internal/routers/*.go.gen*
	@rm -vrf internal/handler/*.go.gen*
	@rm -vrf internal/service/*.go.gen*
	@rm -rf stock-binary.tar.gz
	@echo "clean finished"


# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m  %-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := all
