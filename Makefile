# Copyright 2018 Google LLC.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not
# use this file except in compliance with the License. You may obtain a copy of
# the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations under
# the License.

GOOS ?= $(shell go env GOOS)
GOARCH = amd64
BUILD_DIR ?= ./out

ORG := github.com/coollog
PROJECT := gitcd
REPOPATH ?= $(ORG)/$(PROJECT)

SUPPORTED_PLATFORMS := linux-$(GOARCH) darwin-$(GOARCH) windows-$(GOARCH).exe
BUILD_PACKAGE = $(REPOPATH)/cmd/gitcd

GO_FILES := $(shell find . -type f -name '*.go' -not -path "./vendor/*")

gitcd: $(GO_FILES) $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(PROJECT) $(BUILD_PACKAGE)

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

fmt:
	find . -name "*.go" | grep -v vendor/ | xargs gofmt -l -s

install: $(GO_FILES) $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go install $(BUILD_PACKAGE)


# RELEASE >>
$(BUILD_DIR)/$(PROJECT)-%-$(GOARCH): $(GO_FILES) $(BUILD_DIR)
	GOOS=$* GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o $@ $(BUILD_PACKAGE)

%.exe: %
	cp $< $@

cross: $(foreach platform, $(SUPPORTED_PLATFORMS), $(BUILD_DIR)/$(PROJECT)-$(platform))
# << RELEASE

clean:
	rm -rf $(BUILD_DIR)

test:
	go test -v `go list ./... | grep -v vendor`
