package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

// preguiçoso mas funciona
//
//go:embed index.html
var index string

func main() {
	http.HandleFunc("/adicionar/node", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<div>foobar</div>`))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(index))
	})
	fmt.Println("http://localhost:7373")
	log.Fatal(http.ListenAndServe(":7373", nil))
}
