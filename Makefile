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
	@printf "%s\n" '$(OK)'
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

$(PRIMECALC_OUTPUT): $(GOFILES_BUILD)
    @echo -n ">> BUILD, version = $(PRIMECALC_VERSION)/$(PRIMECALC_REVISION), output = $@)"
    @$(GO) build -o $@ $(BUILDFLAGS)
    @printf '%s\n' '$(OK)'

install: all
    @echo -n ">> INSTALL, version = $(PRIMECALC_VERSION)"
    @install -m 0755 -d $(DESTDIR)$(BINDIR)
    @install -m 0755 -d $(PRIMECALC_OUTPUT) $(DESTDIR)$(BINDIR)/primecalc
    @printf '%s\n' '$(OK)'

test: $(PRIMECALC_OUTPUT)
    @echo -n ">> TEST \"fast-mode\": race detector off"
    @echo "mode: count" > coverage-all.out
    @$(foreach pkg, $(PKGS),\
        echo -n "     ";\
        $(GO) test $(BUILDFLAGS) -coverprofile=coverage.out -covermode=count $(pkg)||exit 1;\
        tail -n +2 coverage.out >> coverage-all.out;)
    @$(GO) tool cover -html=coverage-all.out -o coverage-all.html

crosscompile:
	@echo -n ">> CROSSCOMPILE linux/amd64"
	@GOOS=linux GOARCH=amd64 $(GO) build -o $(PRIMECALC_OUTPUT)-linux-amd64
	@printf '%s\n' '$(OK)'
	@echo -n ">> CROSSCOMPILE darwin/amd64"
	@GOOS=darwin GOARCH=amd64 $(GO) build -o $(PRIMECALC_OUTPUT)-macOS-amd64
	@printf '%s\n' '$(OK)'
	@echo -n ">> CROSSCOMPILE windows/amd64"
	@GOOS=windows GOARCH=amd64 $(GO) build -o $(PRIMECALC_OUTPUT)-windows-amd64
	@printf '%s\n' '$(OK)'

check-release-env:
	ifndef GITHUB_TOKEN
		$(error GITHUB_TOKEN is undefined)
	endif

release: goreleaser

goreleaser: check-release-env travis clean
	@echo ">> RELEASE, goreleaser"
	@goreleaser

