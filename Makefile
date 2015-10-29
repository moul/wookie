SOURCES := $(shell find . -name "*.go")

all: wookie

.PHONY: build
build: wookie

.PHONY: graphviz
graphviz: data/wookie.svg

data/wookie.svg: data/wookie.gv
	dot -Tsvg -o $@ $<

data/wookie.gv: wookie
	rm -f data/wookie.svg
	./wookie -graphviz ./data/wookie.txt > data/wookie.gv

.PHONY: pprof
pprof: data/wookie-100.prof data/wookie-200.prof data/wookie-1000.prof
	rm -rf ./data/wookie-*.svg ./data/wookie-*.txt
	go tool pprof -svg -output=./data/wookie-100.svg wookie data/wookie-100.prof
	go tool pprof -text -output=./data/wookie-100.txt wookie data/wookie-100.prof
	go tool pprof -svg -output=./data/wookie-200.svg wookie data/wookie-200.prof
	go tool pprof -text -output=./data/wookie-200.txt wookie data/wookie-200.prof
	go tool pprof -svg -output=./data/wookie-1000.svg wookie data/wookie-1000.prof
	go tool pprof -text -output=./data/wookie-1000.txt wookie data/wookie-1000.prof

data/wookie-100.prof data/wookie-200.prof data/wookie-1000.prof: wookie
	rm -f $@
	$(eval COUNT := $(shell echo $@ | sed 's/.*-\([0-9][0-9]*\).*/\1/'))
	./wookie -cpuprofile=$@ -count=$(COUNT) ./data/wookie.txt >/dev/null

.PHONY: clean
clean:
	rm -f data/*.prof data/*.svg wookie

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
