package main

import (
	"errors"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type conjunto []*Node

type Node struct {
	conjunto *conjunto
	rotulo   string
	id       uuid.UUID
	filhos   []*Node
	x, y     float64
}

// aresta unidirecional entre dois nós: a -> b.
func (a *Node) Conectar(b *Node) {
	_ = *b
	a.filhos = append(a.filhos, b)
}

// deve retornar o primeiro nó com o mesmo id caso haja mais de um.
func (a *Node) Get(id uuid.UUID) (*Node, error) {
	// panic("não implementado!")
	if a.id == id {
		return a, nil
	} else {
		for _, node := range a.filhos {
			if node.id == id {
				return node, nil
			}
		}
		return nil, errors.New("nó não encontrado")
	}
}

// grau de um vértice
func (a *Node) Grau() int {
	return len(a.filhos)
}

// deve remover o primeiro nó caso o haja mais de um com o mesmo id.
func (a *Node) Remover(id uuid.UUID) {
	idx := slices.IndexFunc(a.filhos, func(x *Node) bool {
		return x.id == id
	})
	if idx != -1 {
		a.filhos = slices.Delete(a.filhos, idx, idx+1)
	} else {
		panic("O IndexFunc retornou -1 para o Node, esse índice não existe!")
	}
}

func (c *conjunto) Remover(rotulo string) {

	idx := slices.IndexFunc(*c, func(x *Node) bool {
		return x.rotulo == rotulo
	})

	if idx != -1 { // Evita erro ao acessar índices inválidos
		*c = slices.Delete(*c, idx, idx+1)
	} else {
		panic("O IndexFunc retornou -1 para o conjunto, esse índice não existe!")
	}

	for _, node := range *c {
		idxFilho := slices.IndexFunc(node.filhos, func(x *Node) bool {
			return x.rotulo == rotulo
		})

		if idxFilho != -1 {
			node.filhos = slices.Delete(node.filhos, idxFilho, idxFilho+1)
		}
	}
}

// cria um novo nó sem conexões e com tal valor.
func (c *conjunto) NovoNode(rotulo string) *Node {
	node := &Node{
		conjunto: c,
		rotulo:   rotulo,
		filhos:   make([]*Node, 0),
		id:       uuid.New(),
	}
	*c = append(*c, node)
	return node
}

func (c *conjunto) Get(rotulo string) *Node {
	idx := slices.IndexFunc(*c, func(x *Node) bool {
		return x.rotulo == rotulo
	})
	return (*c)[idx]
}

func (c conjunto) String() string {
	return c.String()
}

// cria um novo nó filho pertencendo ao mesmo conjunto que o objeto sendo chamado.
func (a *Node) NovoNode(rotulo string) *Node {
	node := a.conjunto.NovoNode(rotulo)
	a.Conectar(node)
	return node
}

// cria um novo conjunto vazio de nós.
func NovoConjunto() *conjunto {
	return new(conjunto)
}

func contarOcorrenciasDoMesmoNode(pai, filho *Node) int {
	retorno := 0
	for _, n := range pai.filhos {
		if n == filho {
			retorno++
		}
	}
	return retorno
}

// cria uma matriz de adjacência onde todos os nós são ordenados pelo rótulo em colunas e linhas
// em ordem alfabética
func (c conjunto) MatrizAdjacencia() ([][]int, []string) {
	clone := slices.Clone(c)
	slices.SortFunc(clone, func(a, b *Node) int {
		return strings.Compare(a.rotulo, b.rotulo)
	})
	var (
		rotulos = make([]string, len(clone))
		matriz  = make([][]int, len(clone))
	)
	for i := range matriz {
		matriz[i] = make([]int, len(clone))
	}
	for i := range clone {
		rotulos[i] = clone[i].rotulo
		for j := range clone {
			matriz[i][j] = contarOcorrenciasDoMesmoNode(clone[i], clone[j])
		}
	}
	return matriz, rotulos
}

// cria uma lista de adjacência onde todos os conjuntos são representados pelos seus rótulos, em ordem alfabética,
// onde cada índice aponta para uma array dos nós filhos
func (c conjunto) ListaAdjacencia() conjunto {
	return c
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
