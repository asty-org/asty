BINDIR=./bin
BLDFLAGS=-ldflags="-s -w"

.PHONY: all build

all: build

test:
	go test -v -coverpkg=github.com/asty-org/asty/asty -covermode=set -coverprofile=coverage.cov ./asty

build:
	go build ${BLDFLAGS} -o ${BINDIR}/asty
