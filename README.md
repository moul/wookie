# wookie
:game_die: Wookie genome solver

[![Build Status](https://travis-ci.org/moul/wookie.svg?branch=master)](https://travis-ci.org/moul/wookie)
[![GoDoc](https://godoc.org/github.com/moul/wookie?status.svg)](https://godoc.org/github.com/moul/wookie)
[![Coverage Status](https://coveralls.io/repos/moul/wookie/badge.svg?branch=master&service=github)](https://coveralls.io/github/moul/wookie?branch=master)

![Wookie logo](https://raw.githubusercontent.com/moul/wookie/master/cmd/appspot/static/wookie.jpg)

Solution for [wookie challenge](https://github.com/jeannedhack/programmingChallenges/tree/master/wookie%20à%20poil%20dur)

## Demo

http://wookie-genome.appspot.com

## Bonus

- super fast ([1000 runs in 1.15s](https://github.com/moul/wookie/blob/master/data/wookie-1000.txt))
- web version (API + form) ([link](#demo))
- export graphs to graphviz ([link](https://github.com/moul/wookie/blob/master/data/wookie.svg))
- automatically detects join length
- flexible parsing of the sequences
- unit tests with nice coverage score

## Install

```console
$ go get github.com/moul/wookie/cmd/wookie
```

## Performances

```console
$ time ./wookie -count=1000 ./data/wookie.txt
AAAAACAAAAGAAAATAAACCAAACGAAACTAAAGCAAAGGAAAGTAAATCAAATGAAATTAACACAACAGAACATAACCCAACCGAACCTAACGCAACGGAACGTAACTCAACTGAACTTAAGACAAGAGAAGATAAGCCAAGCGAAGCTAAGGCAAGGGAAGGTAAGTCAAGTGAAGTTAATACAATAGAATATAATCCAATCGAATCTAATGCAATGGAATGTAATTCAATTGAATTTACACCACACGACACTACAGCACAGGACAGTACATCACATGACATTACCAGACCATACCCCACCCGACCCTACCGCACCGGACCGTACCTCACCTGACCTTACGAGACGATACGCCACGCGACGCTACGGCACGGGACGGTACGTCACGTGACGTTACTAGACTATACTCCACTCGACTCTACTGCACTGGACTGTACTTCACTTGACTTTAGAGCAGAGGAGAGTAGATCAGATGAGATTAGCATAGCCCAGCCGAGCCTAGCGCAGCGGAGCGTAGCTCAGCTGAGCTTAGGATAGGCCAGGCGAGGCTAGGGCAGGGGAGGGTAGGTCAGGTGAGGTTAGTATAGTCCAGTCGAGTCTAGTGCAGTGGAGTGTAGTTCAGTTGAGTTTATATCATATGATATTATCCCATCCGATCCTATCGCATCGGATCGTATCTCATCTGATCTTATGCCATGCGATGCTATGGCATGGGATGGTATGTCATGTGATGTTATTCCATTCGATTCTATTGCATTGGATTGTATTTCATTTGATTTTCCCCCGCCCCTCCCGGCCCGTCCCTGCCCTTCCGCGCCGCTCCGGGCCGGTCCGTGCCGTTCCTCGCCTCTCCTGGCCTGTCCTTGCCTTTCGCGGCGCGTCGCTGCGCTTCGGCTCGGGGCGGGTCGGTGCGGTTCGTCTCGTGGCGTGTCGTTGCGTTTCTCTGCTCTTCTGGGCTGGTCTGTGCTGTTCTTGGCTTGTCTTTGCTTTTGGGGGTGGGTTGGTGTGGTTTGTGTTGTTTT
./wookie -count=1000 ./data/wookie.txt  1.40s user 0.02s system 101% cpu 1.401 total
# =~1.4ms per iteration
```

## License

MIT
