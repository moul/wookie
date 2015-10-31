package wookieapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	sequences := strings.Split(strings.TrimSpace(string(resolveRequest.Sequences)), "\n")
	wook := wookie.NewWookie(sequences)
	wook.Compute()
	output := wook.Genome.String()

	fmt.Fprintf(w, "%v\n", output)
}
