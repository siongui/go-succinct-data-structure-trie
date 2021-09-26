# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
ifndef GOROOT
export GOROOT=$(realpath $(CURDIR)/../go)
export PATH := $(GOROOT)/bin:$(PATH)
endif


test:
	@# -v means verbose, can see logs of t.Log
	@go test -v -race

run_basic:
	@go run example/basic/usage.go

run_pali:
	@cd example/pali; go run encode.go
	@cd example/pali; go run decode.go

bitsjs:
	chromium-browser reference/test.html

fmt:
	@go fmt *.go
	@go fmt example/basic/*.go
	@go fmt example/pali/*.go

help:
	@go help

modinit:
	go mod init github.com/siongui/go-succinct-data-structure-trie

modtidy:
	go mod tidy
