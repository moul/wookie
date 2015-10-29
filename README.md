# wookie
:game_die: Wookie genome solver

[![Build Status](https://travis-ci.org/moul/wookie.svg?branch=master)](https://travis-ci.org/moul/wookie)
[![GoDoc](https://godoc.org/github.com/moul/wookie?status.svg)](https://godoc.org/github.com/moul/wookie)
[![Coverage Status](https://coveralls.io/repos/moul/wookie/badge.svg?branch=master&service=github)](https://coveralls.io/github/moul/wookie?branch=master)

## Install

```console
$ go get github.com/moul/wookie/cmd/wookie
```

## Performances

```console
$ time ./wookie -count=1000 ./data/wookie.txt > /dev/null
./wookie -count=1000 ./data/wookie.txt > /dev/null  5.65s user 0.03s system 99% cpu 5.698
# <5.65ms per iteration
```

## License

MIT
