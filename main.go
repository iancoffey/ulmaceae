package main

import (
	"fmt"
	"log"
	"net/http"
)

type tyHandler struct{}

func (h tyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("at=request path=%s", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ty!")
}

func main() {
	err := http.ListenAndServe(":8080", tyHandler{})
	log.Fatal(err)
}
