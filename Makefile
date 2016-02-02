# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


test:
	@# -v means verbose, can see logs of t.Log
	@go test -v

run:
	@go run example/usage.go

bitsjs:
	chromium-browser reference/test.html

fmt:
	@#go fmt Bits.go
	@#go fmt Bits_test.go
	go fmt *.go

help:
	@go help

install:
	go get -u github.com/siongui/go-succinct-data-structure-trie
