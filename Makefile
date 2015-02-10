# Copyright (c) Orange Applications for Business.
#
# This software is confidential and proprietary information of
# Orange Applications for Business. You shall not disclose such Confidential
# Information and shall use it only in accordance with the terms of the
# agreement you entecolors.red into. Unauthorized copying of this file, via any
# medium is strictly prohibited.

APP="go-onlinelabs"
EXE="onlinelabs"

SHELL = /bin/bash

DIR = $(shell pwd)
GO_PATH = $(DIR)/Godeps/_workspace:$(DIR)

DOCKER = docker
GODEP= $(DIR)/Godeps/_workspace/bin/godep
GOLINT= $(DIR)/Godeps/_workspace/bin/golint
ERRCHECK= $(DIR)/Godeps/_workspace/bin/errcheck
GOVER= $(DIR)/Godeps/_workspace/bin/gover
GOVERALLS= $(DIR)/Godeps/_workspace/bin/goveralls

NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m

VERSION=$(shell \
        grep "const Version" src/github.com/nlamirault/go-onlinelabs/version/version.go \
        |awk -F'=' '{print $$2}' \
        |sed -e "s/[^0-9.]//g" \
	|sed -e "s/ //g")

PACKAGE=$(APP)-$(VERSION)
ARCHIVE=$(PACKAGE).tar

all: help

help:
	@echo -e "$(OK_COLOR)==== $(APP) [$(VERSION)] ====$(NO_COLOR)"
	@echo -e "$(WARN_COLOR)init$(NO_COLOR)   :  Install requirements"
	@echo -e "$(WARN_COLOR)deps$(NO_COLOR)   :  Install dependencies"
	@echo -e "$(WARN_COLOR)build$(NO_COLOR)  :  Make all binaries"
	@echo -e "$(WARN_COLOR)clean$(NO_COLOR)  :  Cleanup"
	@echo -e "$(WARN_COLOR)reset$(NO_COLOR)  :  Remove all dependencies"

clean:
	@echo -e "$(OK_COLOR)[$(APP)] Cleanup$(NO_COLOR)"
	@rm -f $(EXE) $(EXE)_* $(APP)-*.tar.gz coverage.out gover.coverprofile

.PHONY: destroy
destroy:
	@echo -e "$(OK_COLOR)[$(APP)] Destruction environnement de developpement$(NO_COLOR)"
	@rm -fr $(VENV)

.PHONY: init
init:
	@echo -e "$(OK_COLOR)[$(APP)] Install requirements$(NO_COLOR)"
	@GOPATH=$(GO_PATH) go get github.com/golang/glog
	@GOPATH=$(GO_PATH) go get github.com/tools/godep
	@GOPATH=$(GO_PATH) go get -u github.com/golang/lint/golint
	@GOPATH=$(GO_PATH) go get -u github.com/kisielk/errcheck

deps:
	@echo -e "$(OK_COLOR)[$(APP)] Install dependancies$(NO_COLOR)"
	@GOPATH=$(GO_PATH) $(GODEP) restore

build:
	@echo -e "$(OK_COLOR)[$(APP)] Build $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go build -o $(EXE) github.com/nlamirault/$(APP)

doc:
	@godoc -http=:6060 -index

fmt:
	@echo -e "$(OK_COLOR)[$(APP)] Launch fmt $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go fmt github.com/nlamirault/$(APP)/...

errcheck:
	@echo -e "$(OK_COLOR)[$(APP)] Launch errcheck $(NO_COLOR)"
	@GOPATH=$(GO_PATH) $(ERRCHECK) github.com/nlamirault/$(APP)/...

vet:
	@echo -e "$(OK_COLOR)[$(APP)] Launch vet $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go vet github.com/nlamirault/$(APP)/...

lint:
	@echo -e "$(OK_COLOR)[$(APP)] Launch golint $(NO_COLOR)"
	@GOPATH=$(GO_PATH) $(GOLINT) github.com/nlamirault/$(APP)/...

style: fmt vet lint

test:
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test -v github.com/nlamirault/$(APP)/...

race:
	@echo -e "$(OK_COLOR)[$(APP)] Launc unit tests race $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test -race github.com/nlamirault/$(APP)...

coverage:
	@echo -e "$(OK_COLOR)[$(APP)] Launch code coverage $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test github.com/nlamirault/$(APP)/... -cover

covoutput:
	@echo -e "$(OK_COLOR)[$(APP)] Launch code coverage $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test ${pkg} -covermode=count -coverprofile=coverage.out
#	@GOPATH=$(GO_PATH) go tool cover -html=coverage.out
	@GOPATH=$(GO_PATH) go tool cover -func=coverage.out

coveralls:
	@GOPATH=$(GO_PATH) go get -u code.google.com/p/go.tools/cmd/cover || go get -u golang.org/x/tools/cmd/cover
	@GOPATH=$(GO_PATH) go get -u github.com/axw/gocov/gocov
	@GOPATH=$(GO_PATH) go get github.com/mattn/goveralls
	@PATH=$(PATH):$(DIR)/Godeps/_workspace/bin/ GOPATH=$(GO_PATH) $(GOVERALLS) -service drone.io -repotoken $$COVERALLS_TOKEN github.com/nlamirault/$(APP)

release: clean build
	@echo -e "$(OK_COLOR)[$(APP)] Make archive $(VERSION) $(NO_COLOR)"
	@rm -fr $(PACKAGE) && mkdir $(PACKAGE)
	@cp -r $(EXE) $(PACKAGE)
	@tar cf $(ARCHIVE) $(PACKAGE)
	@gzip $(ARCHIVE)
	@rm -fr $(PACKAGE)
	@addons/github.sh $(VERSION)

# for go-projectile
gopath:
	@echo ${GOPATH}
