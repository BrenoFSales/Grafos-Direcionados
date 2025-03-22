package main

import (
	"errors"
	"fmt"
)
type conjunto []*Node

type Node struct {
	conjunto conjunto
	id       string
	filhos   []*Node
}

// aresta unidirecional entre dois nós: a -> b.
func (a *Node) Conectar(b *Node) error {
	if a == nil || b == nil {
		return errors.New("O nó não pode ser nulo.")
	}
	a.filhos = append(a.filhos, b)
	return nil
}

// deve retornar o primeiro nó com o mesmo valor caso haja mais de um.
func (a *Node) Get(valor string) (*Node, error) {
	// panic("não implementado!")
	if a.id == valor {
		return a, nil
	} else {
		for index, node := range a.filhos {
			if node.id == valor {
				fmt.Printf("-- Seu nó está no índice: %d\n", index)
				fmt.Println("-- Informação do nó:")
				fmt.Printf("- Id: %s\n-- Filhos:\n", node.id)
				for _, filho := range node.filhos {
					fmt.Printf("- %s\n", filho.id)
				}
			}
		}
		return nil, errors.New("nó não encontrado")
	}
}

// grau de um vértice
func (a *Node) Grau() int {
	// panic("não implementado!")
	return len(a.filhos)
}

// deve remover o primeiro nó caso o haja mais de um.
// Por que isso aqui receberia um nó ao invés do valor que deve ser buscado wtf?
// Precisa retornar um erro?
// func (a *Node) Remover(b *Node) (*Node, error) ???
func (a *Node) Remover(id string) {
	// panic("não implementado!")
	nodePai, err := a.Get(id)
	if err != nil {
		fmt.Println(err)
	}
	for index, node := range a.filhos {
		if node.id == nodePai.id {
			// https://www.geeksforgeeks.org/delete-elements-in-a-slice-in-golang/
			a.filhos = append(a.filhos[:index], a.filhos[index+1:]...)
		}
	}
}

// cria um novo nó sem conexões com tal valor.
func (c conjunto) NovoNode(id string) *Node {
	node := &Node{
		conjunto: c,
		id:       id,
		filhos:   make([]*Node, 0),
	}
	c = append(c, node)
	return node
}

// cria um novo nó filho pertencendo ao mesmo conjunto que o objeto sendo chamado.
func (a *Node) NovoNode(id string) *Node {
	node := a.conjunto.NovoNode(id)
	a.Conectar(node)
	return node
}

// cria um novo conjunto vazio de nós.
func NovoConjunto() conjunto {
	return make(conjunto, 0)
}

// cria uma matriz de adjacência onde todos os nós são ordenados pelo id em colunas e linhas
// em ordem crescente
func (c conjunto) MatrizAdjacencia() [][]int {
	panic("não implementado!")
}

// cria uma lista de adjacência onde todos os conjuntos são representados pelos seus IDs, em ordem alfabética,
// onde cada índice aponta para uma array dos nós filhos
func (c conjunto) ListaAdjacencia() [][]*Node {
	panic("não implementado!")
}

func (c conjunto) VerificarArvore() (GrafoArvore bool, ArvoreBinaria bool, ArvoreCheia bool, ArvoreCompleta bool) {
	// não fica claro no documento quais tipos de árvores ele quer.
	// talvez árvore binária seja redundante.
	panic("não implementado!")
}

// retorna um map onde cada chave representa um nó do conjunto
// e cada valor o número de grau da nó chave.
func (c conjunto) VerticesGrau() map[*Node]int {
	panic("não implementado!")
}

// retorna o número máximo do grau de todos os vértices
func (c conjunto) GrafoGrau() int {
	panic("não implementado!")
}

// verifica se grafo é completo
func (c conjunto) VerificarCompleto() bool {
	panic("não implementado!")
}

// verifica se grafo possuí ao menos um vértice com um laço
func (c conjunto) VerificarLacos() bool {
	panic("não implementado!")
}

// verifica se grafo é um grafo simples
func (c conjunto) VerificarSimples() bool {
	panic("não implementado!")
}
