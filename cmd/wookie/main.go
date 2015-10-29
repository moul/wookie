package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime/pprof"
	"strings"

	"github.com/moul/wookie"
)

var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
var count = flag.Int("count", 1, "amount of iterations")

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
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	for i := 0; i < *count; i++ {
		wookie := wookie.NewWookie(lines)
		wookie.Compute()
		fmt.Println(wookie.Genome.String())
	}
}
