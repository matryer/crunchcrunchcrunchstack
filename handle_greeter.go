package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func (s *server) handleGreeter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(io.LimitReader(r.Body, 100000))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		name := string(b)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Well 'ello there, %s.", name)
	}
}
