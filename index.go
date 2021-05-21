package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Home Page"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting serve on Port 4000")
	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
