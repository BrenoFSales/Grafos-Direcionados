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
	Id string  `json:"id"`
	X  float64 `json:"x"`
	Y  float64 `json:"y"`
}

type D3Link struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type D3NodeLink struct {
	Nodes []D3Node `json:"nodes"`
	Links []D3Link `json:"links"`
}

func paraD3(conjunto conjunto) D3NodeLink {

	var saida D3NodeLink
	for _, node := range conjunto {
		saida.Nodes = append(saida.Nodes, D3Node{
			Id: node.rotulo,
			X:  node.x,
			Y:  node.y,
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

	http.Handle("/", http.FileServer(http.Dir("./")))
	http.HandleFunc("GET /node", func(w http.ResponseWriter, r *http.Request) {

		saida := paraD3(principal)

		bytes, err := json.Marshal(saida)

		if err != nil {
			panic(err)
		}

		_, _ = w.Write(bytes)
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

		salvo := principal.NovoNode(entrada.Id)
		salvo.x = entrada.X
		salvo.y = entrada.Y

		saida := paraD3(principal)
		bytes, err = json.Marshal(saida)
		if err != nil {
			panic(err)
		}

		_, _ = w.Write(bytes)
	})

	fmt.Println("http://localhost:7373")
	log.Fatal(http.ListenAndServe(":7373", nil))
}
