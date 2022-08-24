BINDIR=./bin
BLDFLAGS=-ldflags="-s -w"

.PHONY: all build

all: build

build:
	go build ${BLDFLAGS} -o ${BINDIR}/asty
