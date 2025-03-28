package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
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

type exemploSelecao map[string]*conjunto

func inicializarExemplos() map[string]*conjunto {
	principal := NovoConjunto()
	a := principal.NovoNode("a")
	b := principal.NovoNode("b")
	c := principal.NovoNode("c")

	a.Conectar(b)
	b.Conectar(c)
	c.Conectar(a)

	grafoCompleto := NovoConjunto()
	grafoCompletoVertices := []*Node{
		grafoCompleto.NovoNode("1"), grafoCompleto.NovoNode("2"), grafoCompleto.NovoNode("3"),
		grafoCompleto.NovoNode("4"), grafoCompleto.NovoNode("5"), grafoCompleto.NovoNode("6"),
		grafoCompleto.NovoNode("7"),
	}
	for i := range grafoCompletoVertices {
		for j := range grafoCompletoVertices {
			if i == j {
				continue
			}
			a, b := grafoCompletoVertices[i], grafoCompletoVertices[j]
			a.Conectar(b)
			b.Conectar(a)
		}
	}

	arvoreBinaria := NovoConjunto()
	a = arvoreBinaria.NovoNode("a")
	b = a.NovoNode("b")
	c = a.NovoNode("c")

	d := b.NovoNode("d")
	e := b.NovoNode("e")

	f := c.NovoNode("f")
	g := c.NovoNode("g")

	d.NovoNode("h")
	d.NovoNode("i")
	e.NovoNode("j")
	e.NovoNode("k")

	f.NovoNode("l")
	f.NovoNode("m")
	g.NovoNode("n")
	g.NovoNode("o")

	exemplos := exemploSelecao{
		"principal": principal,
		"completo":  grafoCompleto,
		"binaria":   arvoreBinaria,
	}
	return exemplos
}

func buscarConjunto(exemplos exemploSelecao, r *http.Request) *conjunto {
	conjuntoSelectionado := r.PathValue("conjunto")
	if conjuntoSelectionado == "" {
		panic("vazio")
	}

	return exemplos[conjuntoSelectionado]
}

var exemplos = inicializarExemplos()

func buscarExemplos(_ http.ResponseWriter, _ *http.Request) exemploSelecao {
	return exemplos
}

func main() {

	http.HandleFunc("/link/{conjunto}", func(w http.ResponseWriter, r *http.Request) {

		exemplos := buscarExemplos(w, r)
		conjunto := buscarConjunto(exemplos, r)

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

		switch r.Method {
		case "POST":
			source.Conectar(target)
			w.WriteHeader(204)
		case "DELETE":
			source.Remover(target.id)
			w.WriteHeader(204)
		default:
			w.WriteHeader(405)
		}

	})

	http.Handle("/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/node/{conjunto}", func(w http.ResponseWriter, r *http.Request) {

		exemplos := buscarExemplos(w, r)
		conjunto := buscarConjunto(exemplos, r)

		switch r.Method {
		case "GET":
			saida := paraD3(conjunto)

			bytes, err := json.Marshal(saida)

			if err != nil {
				panic(err)
			}

			_, _ = w.Write(bytes)
		case "POST":
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

			conjunto.NovoNode(entrada.Id)

			w.WriteHeader(204)

		case "DELETE":
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

			conjunto.Remover(entrada.Id)

			w.WriteHeader(204)
		default:
			w.WriteHeader(405)
		}

	})

	http.HandleFunc("/matriz/{conjunto}", func(w http.ResponseWriter, r *http.Request) {

		exemplos := buscarExemplos(w, r)
		conjunto := buscarConjunto(exemplos, r)

		rotulos_ := r.URL.Query().Get("rotulo")
		if rotulos_ == "" {
			panic("vazio")
		}
		exibirRotulos := rotulos_ == "true"

		matriz, rotulos := conjunto.MatrizAdjacencia()

		for i := range matriz {
			var linha []string
			for j := range matriz[i] {
				var s string
				if exibirRotulos {
					s = fmt.Sprintf("%d_{%s,%s}", matriz[i][j], rotulos[i], rotulos[j])
				} else {
					s = fmt.Sprintf("%d", matriz[i][j])
				}
				linha = append(linha, s)
			}
			fmt.Fprintf(w, "%s \\\\\n", strings.Join(linha, " & "))
		}

	})

	http.HandleFunc("/lista/{conjunto}", func(w http.ResponseWriter, r *http.Request) {

		exemplos := buscarExemplos(w, r)
		conjunto := buscarConjunto(exemplos, r)

		// cria a representação em latex de uma lista de adjacência.
		// o resultado é sempre algo parecido com isso:
		//
		// \begin{array}{l}
		//	1: 2 \\
		// 	2: 1, 3 \\
		// 	3: 2 \\
		// \end{array}

		// tamanhoStringMaximo := 0
		// for _, node := range *conjunto {
		// 	tamanhoStringMaximo = max(tamanhoStringMaximo, len(node.rotulo))
		// }

		fmt.Fprintf(w, "\\begin{array}{l}\n")
		for _, node := range *conjunto {
			var s []string
			for _, filho := range node.filhos {
				s = append(s, filho.rotulo)
			}
			fmt.Fprintf(w, "%s: %s \\\\\n", node.rotulo, strings.Join(s, ", "))
		}
		fmt.Fprintf(w, "\\end{array}\n")
	})

	http.HandleFunc("/grau/{conjunto}", func(w http.ResponseWriter, r *http.Request) {

		exemplos := buscarExemplos(w, r)
		conjunto := buscarConjunto(exemplos, r)

		resultado := make(map[string]map[string]int)

		for _, node := range *conjunto {
			entrada, saida := node.Grau()
			resultado[node.rotulo] = map[string]int{
				"entrada": entrada,
				"saida":   saida,
			}
		}

		bytes, err := json.MarshalIndent(resultado, "", "  ")
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(bytes)
	})

	http.HandleFunc("/tipo/{conjunto}", func(w http.ResponseWriter, r *http.Request) {

		exemplos := buscarExemplos(w, r)
		conjunto := buscarConjunto(exemplos, r)

		type tipos struct {
			Arvore   bool `json:"arvore"`
			Binaria  bool `json:"binaria"`
			Cheia    bool `json:"cheia"`
			Completa bool `json:"completa"`
		}

		var t tipos

		nodeRaizRotulo := r.URL.Query().Get("raiz")
		if nodeRaizRotulo == "" {
			t = tipos{}
		} else {
			raiz := conjunto.Get(nodeRaizRotulo)
			GrafoArvore, ArvoreBinaria, ArvoreCheia, ArvoreCompleta := conjunto.VerificarArvore(raiz)
			t = tipos{
				Arvore:   GrafoArvore,
				Binaria:  ArvoreBinaria,
				Cheia:    ArvoreCheia,
				Completa: ArvoreCompleta,
			}
		}
		bytes, err := json.Marshal(t)
		if err != nil {
			panic(err)
		}
		w.Write(bytes)
		return

	})

	fmt.Println("http://0.0.0.0:7373")
	log.Fatal(http.ListenAndServe("0.0.0.0:7373", logging(http.DefaultServeMux)))
}
