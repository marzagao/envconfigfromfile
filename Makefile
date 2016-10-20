.PHONY: all testdeps lint test build run

all: test

testdeps:
	go get -d -t ./...

lint: testdeps
	go get github.com/alecthomas/gometalinter honnef.co/go/unused/cmd/unused
	gometalinter --install --vendored-linters
	go install ./...
	go list -f '{{.TestImports}}' ./... | sed -e 's/\[\(.*\)\]/\1/' | tr ' ' '\n' | grep '^.*\..*/.*$$' | xargs go install
	gometalinter -j 4 --enable=gofmt --enable=unused --disable=dupl --disable=errcheck --disable=gas --disable=interfacer --disable=gocyclo --deadline=10m --tests ./...

gotest: testdeps
	go test ./...

test: lint gotest

vet: testdeps
	go vet ./...
