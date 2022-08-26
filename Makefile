BINDIR=./bin
BLDFLAGS=-ldflags="-s -w"

.PHONY: all build

all: build

test:
	go test -v ./asty

build:
	go build ${BLDFLAGS} -o ${BINDIR}/asty
