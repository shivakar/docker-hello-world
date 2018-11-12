package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", ":8080", "Address to serve on")
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!!"))
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
