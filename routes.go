package main

import "net/http"

func (s *server) routes() {
	s.mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	s.mux.HandleFunc("/api/greet", s.handleGreeter())
	s.mux.HandleFunc("/", s.handleIndex())
}
