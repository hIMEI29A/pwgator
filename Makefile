IMPORT_PATH := github.com/hIMEI29A/pwgator
BUILDDIR := $(CURDIR)/build/amd64/pwgator

# Source files
OBJECTS := $(wildcart *.go)

# comment this line out for quieter things
V := 1 # When V is set, print commands and build progress.

# Space separated patterns of packages to skip in list, test, format.
IGNORED_PACKAGES := $(IMPORT_PATH)/build $(IMPORT_PATH)/_tests

.PHONY: all
all: test build

.PHONY: build tags
build:
	@echo "Building..."
	$Q go build $(if $V,-v) $(IMPORT_PATH)/...

.PHONY: tags
tags:
	@echo "Listing tags..."
	$Q @git tag


##### =====> Utility targets <===== #####

.PHONY: test format

test:
	@echo "Testing..."
	$Q go test $(if $V,-v) -i $(allpackages)
ifndef CI
	@echo "Testing Outside CI..."
#	$Q go vet $(allpackages)
	$Q GODEBUG=cgocheck=2 go test $(allpackages)
else
	@echo "Testing in CI..."
	$Q mkdir -p test
	$Q ( go vet $(allpackages); echo $$? ) | \
       tee test/vet.txt | sed '$$ d'; exit $$(tail -1 test/vet.txt)
	$Q ( GODEBUG=cgocheck=2 go test -v -race $(allpackages); echo $$? ) | \
       tee test/output.txt | sed '$$ d'; exit $$(tail -1 test/output.txt)
endif

format: $(GOPATH)/bin/goimports
	@echo "Formatting..."
	$Q find . -iname \*.go | grep -v \
        -e "^$$" $(addprefix -e ,$(IGNORED_PACKAGES)) | xargs goimports -w
		
##### =====> Internals <===== #####

VERSION          := $(shell git describe --tags --always --dirty="-dev")
DATE             := $(shell date -u '+%Y-%m-%d-%H:%M UTC')

# assuming go 1.9 here!!
_allpackages = $(shell go list ./...) 

# memoize allpackages, so that it's executed only once and only if used
allpackages = $(if $(__allpackages),,$(eval __allpackages := $$(_allpackages)))$(__allpackages)


Q := $(if $V,,@)

$(GOPATH)/bin/gocovmerge:
	@echo "Checking Coverage Tool Installation..."
	@test -d $(GOPATH)/src/github.com/wadey/gocovmerge || \
        { echo "Vendored gocovmerge not found, try running 'make setup'..."; exit 1; }
	$Q go install github.com/wadey/gocovmerge
$(GOPATH)/bin/goimports:
	@echo "Checking Import Tool Installation..."
	@test -d $(GOPATH)/src/golang.org/x/tools/cmd/goimports || \
        { echo "Vendored goimports not found, try running 'make setup'..."; exit 1; }
	$Q go install golang.org/x/tools/cmd/goimports

