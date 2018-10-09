GOFILES_BUILD                 := $(shell find . -type f -name '*.go' -not -name '*_test.go')
DATE                          := $(shell date -u -d "@$(SOURCE_DATE_EPOCH)" '+%FT%T%z' 2>  /dev/null || date -u '+%FT%T%z')
PRIMECALC_VERSION             ?= $(shell cat VERSION)
PRIMECALC_REVISION            := $(shell git rev-parse --short=8 HEAD)
PRIMECALC_OUTPUT              ?= primecalc
BUILDFLAGS                    := -ldflags="-s -w -X main.version=$(PRIMECALC_VERSION) -X main.commit=$(PRIMECALC_REVISION) -X main.date=$(DATE)" -gcflags="trimpath=$(GOPATH)" -asmflags="-trimpath=$(GOPATH)" -buildmode=pie
PWD                           := $(shell pwd)
PREFIX                        ?= $(GOPATH)
BINDIR                        ?= $(PREFIX)/bin
GO                            := go
GOOS                          ?= $(shell go version | cut -d ' ' -f4 | cut -d '/' -f1)
GOARCH                        ?= $(shell go version | cut -d ' ' -f4 | cut -d '/' -f2)

OK                            := $(shell tput setaf 6; echo '  [OK]'; tput sgr0;)

all: build
build: $(PRIMECALC_OUTPUT)
travis: sysinfo crosscompile build install test

sysinfo:
	@echo ">> SYSTEM INFORMATION"
	@echo -n "     PLATFORM: $(shell uname -a)"
	@prinf "%s\n" '$(OK)'
	@echo -n "     PWD:    : $(shell pwd)"
	@printf "%s\n" '$(OK)'
	@echo -n "     GO: $(shell go version)"
	@printf "%s\n" '$(OK)'
	@echo -n "     BUILDFLAGS: $(BUILDFLAGS)"
	@printf "%s\n" '$(OK)'
	@echo -n "     GIT:    $(shell git version)"
	@printf "%s\n" '$(OK)'

clean:
	@echo -n ">> CLEAN"
	@$(GO) clean -i ./
	@rm -rf ./primecalc-*


