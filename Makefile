GO_SRCS := $(shell find . -type f -name '*.go')
GO_DIRS := . ./urlinfo/...
BUILD_DIR := build

# urlinfo runs on linux even if its built on mac or windows
TARGET_OS ?= linux
GOARCH ?= $(shell go env GOARCH)
GOPATH ?= $(shell go env GOPATH)
BUILD_UID ?= $(shell id -u)
BUILD_GID ?= $(shell id -g)
BUILD_GO_FLAGS := -tags osusergo
BUILD_GO_CGO_ENABLED ?= 0
BUILD_GO_LDFLAGS_EXTRA :=
DEBUG ?= false

# VERSION ?= $(shell git describe --tags)
VERSION ?= test
ifeq ($(DEBUG), false)
LD_FLAGS ?= -w -s
endif

# https://reproducible-builds.org/docs/source-date-epoch/#makefile
# https://reproducible-builds.org/docs/source-date-epoch/#git
# https://stackoverflow.com/a/15103333
BUILD_DATE_FMT = %Y-%m-%dT%H:%M:%SZ
ifdef SOURCE_DATE_EPOCH
	BUILD_DATE ?= $(shell date -u -d "@$(SOURCE_DATE_EPOCH)" "+$(BUILD_DATE_FMT)" 2>/dev/null || date -u -r "$(SOURCE_DATE_EPOCH)" "+$(BUILD_DATE_FMT)" 2>/dev/null || date -u "+$(BUILD_DATE_FMT)")
else
	BUILD_DATE ?= $(shell TZ=UTC git log -1 --pretty=%cd --date='format-local:$(BUILD_DATE_FMT)' || date -u +$(BUILD_DATE_FMT))
endif

LD_FLAGS += -X build.Version=$(VERSION)
LD_FLAGS += $(BUILD_GO_LDFLAGS_EXTRA)


GOLANG_IMAGE = golang:1.17-alpine
GO ?= GOCACHE=/go/src/urlinfo/build/cache/go/build GOMODCACHE=/go/src/urlinfo/build/cache/go/mod docker run --rm \
	-v "$(CURDIR)":/go/src/urlinfo \
	-w /go/src/urlinfo \
	-e GOOS \
	-e CGO_ENABLED \
	-e GOARCH \
	-e GOCACHE \
	-e GOMODCACHE \
	--user $(BUILD_UID):$(BUILD_GID) \
	$(GOLANG_IMAGE) go

.PHONY: build
ifeq ($(TARGET_OS),windows)
build: urlinfo-bin.exe
else
build: urlinfo-bin
endif

.PHONY: all
all: urlinfo-bin urlinfo-bin.exe

go.sum: go.mod
	$(GO) mod tidy

urlinfo-bin: TARGET_OS = linux
urlinfo-bin: BUILD_GO_CGO_ENABLED = 0
urlinfo-bin: BUILD_GO_LDFLAGS_EXTRA = -extldflags=-static

urlinfo-bin.exe: TARGET_OS = windows
urlinfo-bin.exe: BUILD_GO_CGO_ENABLED = 0

urlinfo-bin.exe urlinfo-bin: $(GO_SRCS) go.sum
	CGO_ENABLED=$(BUILD_GO_CGO_ENABLED) GOOS=$(TARGET_OS) GOARCH=$(GOARCH) $(GO) build $(BUILD_GO_FLAGS) -ldflags='$(LD_FLAGS)' -o $@ urlinfo/urlinfo.go
	rm -rf "$(BUILD_DIR)/bin" && mkdir -p "$(BUILD_DIR)/bin" && mv $@ $(BUILD_DIR)/bin
		
clean-gocache: GO = \
  GOCACHE='$(CURDIR)/build/cache/go/build' \
  GOMODCACHE='$(CURDIR)/build/cache/go/mod' \
  go

.PHONY: clean-gocache
clean-gocache:
	$(GO) clean -cache -modcache

.PHONY: build_image
build_image: urlinfo-bin build/Dockerfile
	docker build --rm \
		-f build/Dockerfile \
		-t urlinfo build/