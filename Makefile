# Makefile for msfs_connect

# We use g++ because SimConnect.h uses C++ features
export CC = g++
export CGO_CFLAGS = -x c++ -I$(CURDIR)
export TMP := $(or $(TMP),$(TEMP),$(TMPDIR),/tmp)

GO = go

.PHONY: all generate test

all: generate test

# Run godefs to update defs_generated.go (committed, so consumers don't need g++ or SimConnect.h)
generate:
	@echo "Generating definitions..."
	$(GO) tool cgo -godefs simconnect/defs.go > simconnect/defs_generated.go

test:
	$(GO) build ./...
	$(GO) test -v ./...
