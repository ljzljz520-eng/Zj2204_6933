package main

import (
	"hospitalforms/internal/api"
	"hospitalforms/internal/forms"
	"hospitalforms/internal/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	path := os.Getenv("FORM_DB")
	if path == "" {
		path = "forms.db"
	}
	s, e := storage.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	f := forms.New(s)
	a := api.New(f)
	mux := http.NewServeMux()
	mux.HandleFunc("/health", a.Health)
	mux.HandleFunc("/records", a.List)
	mux.HandleFunc("/records/create", a.Create)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
