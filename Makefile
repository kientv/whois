# Program version
VERSION := $(shell grep "const Version " version.go | sed -E 's/.*"(.+)"$$/\1/')

# Binary name for beepi
BIN_NAME=$(shell basename $(abspath ./))

# Project owner for beepi
OWNER=kientv

# Project name for beepi
PROJECT_NAME=$(shell basename $(abspath ./))

# Project url used for builds
# examples: github.com, bitbucket.org
REPO_HOST_URL=bitbucket.org

# Grab the current commit
GIT_COMMIT="$(shell git rev-parse HEAD)"

# Check if there are uncommited changes
GIT_DIRTY="$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)"

# Add the godep path to the GOPATH
#GOPATH=$(shell godep path):$(shell echo $$GOPATH)

install:
	go get -v github.com/astaxie/beego
	go get -v github.com/astaxie/beego/orm
	go get -v github.com/astaxie/beego/validation
	go get -v github.com/astaxie/beego/logs
	go get -v github.com/astaxie/beego/config
	go get -v github.com/beego/i18n
	go get -v github.com/beego/bee
	go get -v github.com/go-sql-driver/mysql
	go get -v github.com/smartystreets/goconvey/convey
	go get -v github.com/Shopify/sarama

release:
	@echo "building ${OWNER} ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	#godep get && \
	go build -ldflags "-X main.GitCommit ${GIT_COMMIT}${GIT_DIRTY}" -o bin/${BIN_NAME}

build: clean
	@echo "building ${OWNER} ${BIN_NAME} ${VERSION}"
	@go build -o ./${BIN_NAME}

clean:
	@test ! -e ./${BIN_NAME} || rm ./${BIN_NAME}

test:
	go test -v ./...

run:
	@bee run watchall true

.PHONY: build dist clean test release run install

