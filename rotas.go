package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/adicionar/node", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<div>foobar</div>`))
	})

	fmt.Println("http://localhost:7373")
	log.Fatal(http.ListenAndServe(":7373", nil))
}
