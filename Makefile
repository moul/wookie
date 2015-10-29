SOURCES := $(shell find . -name "*.go")

all: wookie

.PHONY: build
build: wookie

wookie: $(SOURCES)
	go get ./...
	go build -o $@ ./cmd/$@

.PHONY: test
test:
	go get -t ./...
	go test -v .

.PHONY: cover
cover:
	rm -f profile.out
	go test -covermode=count -coverpkg=. -coverprofile=profile.out .

.PHONY: convey
convey:
	go get github.com/smartystreets/goconvey
	goconvey -cover -port=10042 -workDir="$(realpath .)" -depth=-1
