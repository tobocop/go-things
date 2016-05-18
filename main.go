package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	log.Fatal(http.ListenAndServe(":12345", nil))
}

func outerHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello"))
}
