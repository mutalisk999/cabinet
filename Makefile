LATEST_TAG 		:= $(shell git describe --abbrev=0 --tags )
LATEST_COMMIT_SHA1     := $(shell git rev-parse HEAD )

LD_GUI_FLAGS	= -ldflags '-w -s -H windowsgui -X "main.releaseType=release" -X "main.latestTag=${LATEST_TAG}" -X "main.latestCommitSHA1=${LATEST_COMMIT_SHA1}"'
LD_FLAGS	= -ldflags '-w -s -X "main.releaseType=release" -X "main.latestTag=${LATEST_TAG}" -X "main.latestCommitSHA1=${LATEST_COMMIT_SHA1}"'


all: cabinet file_server
.PHONY: all

clean: cabinet_clean file_server_clean
.PHONY: clean

cabinet:
	go build ${LD_GUI_FLAGS} ./cmd/cabinet
.PHONY: cabinet

file_server:
	go build ${LD_FLAGS} ./cmd/file_server
.PHONY: file_server

cabinet_clean:
	rm -f cabinet
.PHONY: cabinet_clean

file_server_clean:
	rm -rf file_server
.PHONY: file_server_clean