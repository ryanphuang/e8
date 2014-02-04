.PHONY: all fmt tags

all:
	go install ./...

fmt:
	gofmt -s -w -l .

tags:
	gotags `find . -name "*.go"` > tags

test:
	go test ./...
