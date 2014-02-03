.PHONY: all fmt tags

all:
	go build

fmt:
	go fmt

tags:
	gotags `find . -name "*.go"` > tags
