package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/moul/wookie"
)

func main() {
	stream, _ := os.Open(os.Args[1])
	input, _ := ioutil.ReadAll(stream)
	wookie := wookie.NewWookie(strings.Split(strings.TrimSpace(string(input)), "\n"))
	wookie.Compute()
	fmt.Println(wookie.Genome.String())
}
