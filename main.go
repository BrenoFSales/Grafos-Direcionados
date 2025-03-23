package mainstring

import (
	"errors"
	"slices"

	"github.com/google/uuid"
)

type conjunto []*Node

type Node struct {
	conjunto conjunto
	rotulo   string
	id       uuid.UUID
	filhos   []*Node
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
	for index, node := range a.filhos {
		if node.id == id {
			a.filhos = slices.Delete(a.filhos, index, index+1)
			return
		}
	}
	panic("nó não encontrado!")
}

// cria um novo nó sem conexões com tal valor.
func (c conjunto) NovoNode(rotulo string) *Node {
	node := &Node{
		conjunto: c,
		rotulo:   rotulo,
		filhos:   make([]*Node, 0),
		id:       uuid.New(),
	}
	c = append(c, node)
	return node
}

// cria um novo nó filho pertencendo ao mesmo conjunto que o objeto sendo chamado.
func (a *Node) NovoNode(rotulo string) *Node {
	node := a.conjunto.NovoNode(rotulo)
	a.Conectar(node)
	return node
}

// cria um novo conjunto vazio de nós.
func NovoConjunto() conjunto {
	return make(conjunto, 0)
}

// cria uma matriz de adjacência onde todos os nós são ordenados pelo rótulo em colunas e linhas
// em ordem crescente
func (c conjunto) MatrizAdjacencia() [][]int {
	panic("não implementado!")
}

// cria uma lista de adjacência onde todos os conjuntos são representados pelos seus rótulos, em ordem alfabética,
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
