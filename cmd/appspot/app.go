package wookieapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/moul/wookie"
)

func init() {
	http.HandleFunc("/", handler)
}

type ResolveRequest struct {
	Sequences string          `json:"Sequences,omitempty"`
	Options   map[string]bool `json:"Options,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var resolveRequest ResolveRequest
	err := decoder.Decode(&resolveRequest)

	if err != nil {
		fmt.Fprintf(w, "POST parsing error: %v\n", err)
		return
	}
	wook := wookie.NewWookie(resolveRequest.Sequences)

	// Timeout
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 3)
		timeout <- true
	}()

	computed := make(chan bool, 1)
	go func() {
		err = wook.Compute()
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		} else {
			fmt.Fprintf(w, "%s", wook.Genome.String())
		}
		computed <- true
	}()

	select {
	case <-computed:
		// compute finished
	case <-timeout:
		fmt.Fprintf(w, "%s", "Timeout, this webapp times out after 3 seconds.\nUse the CLI version if you need to compute bigger genomes.\n\n    https://github.com/moul/wookie\n")
	}
}
