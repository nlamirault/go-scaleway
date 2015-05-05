# Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http:#www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

APP="go-scaleway"
EXE="scaleway"

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
        grep "const Version" src/github.com/nlamirault/go-scaleway/version/version.go \
        |awk -F'=' '{print $$2}' \
        |sed -e "s/[^0-9.]//g" \
	|sed -e "s/ //g")

PACKAGE=$(APP)-$(VERSION)
ARCHIVE=$(PACKAGE).tar

all: help

help:
	@echo -e "$(OK_COLOR)==== $(APP) [$(VERSION)] ====$(NO_COLOR)"
	@echo -e "$(WARN_COLOR)init$(NO_COLOR)    :  Install requirements"
	@echo -e "$(WARN_COLOR)deps$(NO_COLOR)    :  Install dependencies"
	@echo -e "$(WARN_COLOR)build$(NO_COLOR)   :  Make all binaries"
	@echo -e "$(WARN_COLOR)clean$(NO_COLOR)   :  Cleanup"
	@echo -e "$(WARN_COLOR)reset$(NO_COLOR)   :  Remove all dependencies"
	@echo -e "$(WARN_COLOR)release$(NO_COLOR) :  Make a new release"

clean:
	@echo -e "$(OK_COLOR)[$(APP)] Cleanup$(NO_COLOR)"
	@rm -f $(EXE) $(EXE)_* $(APP)-*.tar.gz coverage.out gover.coverprofile

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
	@GOPATH=$(GO_PATH) godoc -http=:6060 -index

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
	@echo -e "$(OK_COLOR)[$(APP)] Launch unit tests race $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test -race github.com/nlamirault/$(APP)/...

coverage:
	@echo -e "$(OK_COLOR)[$(APP)] Launch code coverage $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test github.com/nlamirault/$(APP)/... -cover

covoutput:
	@echo -e "$(OK_COLOR)[$(APP)] Launch code coverage $(NO_COLOR)"
	@GOPATH=$(GO_PATH) go test ${pkg} -covermode=count -coverprofile=coverage.out
#	@GOPATH=$(GO_PATH) go tool cover -html=coverage.out
	@GOPATH=$(GO_PATH) go tool cover -func=coverage.out

coveralls:
	@GOPATH=$(GO_PATH) go get github.com/axw/gocov/gocov
	@GOPATH=$(GO_PATH) go get github.com/mattn/goveralls
	$(DIR)/Godeps/_workspace/bin/goveralls -service=travis-ci

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
