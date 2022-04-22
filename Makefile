.PHONY: all release debug clean

all: release

LATEST_TAG 		:= $(shell git describe --abbrev=0 --tags )
LATEST_COMMIT_SHA1     := $(shell git rev-parse HEAD )

RELEASE_LDFLAGS	= -ldflags '-w -s -H windowsgui -X "main.releaseType=release" -X "main.latestTag=${LATEST_TAG}" -X "main.latestCommitSHA1=${LATEST_COMMIT_SHA1}"'
DEBUG_LDFLAGS = -ldflags '-X "main.releaseType=debug" -X "main.latestTag=${LATEST_TAG}" -X "main.latestCommitSHA1=${LATEST_COMMIT_SHA1}"'
TARGET = cabinet

release:
	go build ${RELEASE_LDFLAGS}

debug:
	go build ${DEBUG_LDFLAGS}

clean:
	rm -rf ${TARGET}