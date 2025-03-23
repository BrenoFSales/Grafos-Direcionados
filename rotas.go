package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type D3Node struct {
	Id string `json:"id"`
}

type D3Link struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type D3NodeLink struct {
	Nodes []D3Node `json:"nodes"`
	Links []D3Link `json:"links"`
}

func paraD3(conjunto *conjunto) D3NodeLink {

	var saida D3NodeLink = D3NodeLink{
		Nodes: make([]D3Node, 0),
		Links: make([]D3Link, 0),
	}
	for _, node := range *conjunto {
		saida.Nodes = append(saida.Nodes, D3Node{
			Id: node.rotulo,
		})
		for _, filho := range node.filhos {
			saida.Links = append(saida.Links, D3Link{
				Source: node.rotulo,
				Target: filho.rotulo},
			)
		}
	}
	return saida
}

func main() {
	principal := NovoConjunto()
	a := principal.NovoNode("a")
	b := principal.NovoNode("b")
	c := principal.NovoNode("c")

	a.Conectar(b)
	b.Conectar(c)
	c.Conectar(a)

	exemplos := map[string]*conjunto{
		"principal": principal,
	}

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("GET /node", func(w http.ResponseWriter, r *http.Request) {

		saida := paraD3(principal)

		bytes, err := json.Marshal(saida)

		if err != nil {
			panic(err)
		}

		_, _ = w.Write(bytes)
	})

	http.HandleFunc("POST /link/{conjunto}", func(w http.ResponseWriter, r *http.Request) {

		conjuntoSelectionado := r.PathValue("conjunto")
		if conjuntoSelectionado == "" {
			panic("vazio")
		}

		conjunto := exemplos[conjuntoSelectionado]

		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var entrada D3Link

		if err = json.Unmarshal(bytes, &entrada); err != nil {
			panic(err)
		}

		if entrada.Source == "" || entrada.Target == "" {
			panic("vazio")
		}

		source := conjunto.Get(entrada.Source)
		target := conjunto.Get(entrada.Target)

		source.Conectar(target)

		w.WriteHeader(204)
	})

	http.HandleFunc("POST /node", func(w http.ResponseWriter, r *http.Request) {

		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		var entrada D3Node

		if err = json.Unmarshal(bytes, &entrada); err != nil {
			panic(err)
		}

		if entrada.Id == "" {
			panic("vazio")
		}

		principal.NovoNode(entrada.Id)

		w.WriteHeader(204)
	})

	fmt.Println("http://localhost:7373")
	log.Fatal(http.ListenAndServe(":7373", nil))
}
