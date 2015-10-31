package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"

	"github.com/moul/wookie"
)

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
var count = flag.Int("count", 1, "amount of iterations")
var graphviz = flag.Bool("graphviz", false, "generate graphviz output")

func main() {
	flag.Parse()

	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	stream, _ := os.Open(flag.Arg(0))
	input, _ := ioutil.ReadAll(stream)

	var output string
	var w wookie.Wookie
	for i := 0; i < *count; i++ {
		w = wookie.NewWookie(string(input))
		err := w.Compute()
		if err != nil {
			log.Fatal(err)
		}
		output = w.Genome.String()
	}

	// only print latest output
	if *graphviz {
		fmt.Println(w.Graphviz())
	} else {
		fmt.Println(output)
	}
}
