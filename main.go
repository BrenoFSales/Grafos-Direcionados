package main

import (
	"errors"
	"slices"
	"strings"

	"github.com/google/uuid"
)

type conjunto []*Node

type Node struct {
	conjunto *conjunto // Conjunto de todos os vértices do grafo
	Rotulo   string    // "Valor" do nó
	Id       uuid.UUID // Identificador único do nó
	Filhos   []*Node   // Relações do nó
	x, y     float64   // Posição dos nós na interface ??????
}

// aresta unidirecional entre dois nós: a -> b.
func (a *Node) Conectar(b *Node) {
	_ = *b
	a.Filhos = append(a.Filhos, b)
}

// deve retornar o primeiro nó com o mesmo id caso haja mais de um.
func (a *Node) Get(id uuid.UUID) (*Node, error) {
	// panic("não implementado!")
	if a.Id == id {
		return a, nil
	} else {
		for _, node := range a.Filhos {
			if node.Id == id {
				return node, nil
			}
		}
		return nil, errors.New("nó não encontrado")
	}
}

// graus de saída e entrada de um vértice
func (a *Node) Grau() (grauEntradaTotal int, grauSaidaTotal int) {
	conjunto := *a.conjunto

	grauSaidaTotal = len(a.Filhos)

	for i := range conjunto {
		for _, f := range conjunto[i].Filhos {
			if f == a {
				grauEntradaTotal++
			}
		}
	}
	return grauEntradaTotal, grauSaidaTotal
}

// deve remover o primeiro nó caso o haja mais de um com o mesmo id.
func (a *Node) Remover(id uuid.UUID) {
	idx := slices.IndexFunc(a.Filhos, func(x *Node) bool {
		return x.Id == id
	})
	if idx != -1 {
		a.Filhos = slices.Delete(a.Filhos, idx, idx+1)
	} else {
		panic("O IndexFunc retornou -1 para o Node, esse índice não existe!")
	}
}

func (c *conjunto) Remover(rotulo string) {

	idx := slices.IndexFunc(*c, func(x *Node) bool {
		return x.Rotulo == rotulo
	})

	if idx != -1 { // Evita erro ao acessar índices inválidos
		*c = slices.Delete(*c, idx, idx+1)
	} else {
		panic("O IndexFunc retornou -1 para o conjunto, esse índice não existe!")
	}

	for _, node := range *c {
		idxFilho := slices.IndexFunc(node.Filhos, func(x *Node) bool {
			return x.Rotulo == rotulo
		})

		if idxFilho != -1 {
			node.Filhos = slices.Delete(node.Filhos, idxFilho, idxFilho+1)
		}
	}
}

// cria um novo nó sem conexões e com tal valor.
func (c *conjunto) NovoNode(rotulo string) *Node {
	node := &Node{
		conjunto: c,
		Rotulo:   rotulo,
		Filhos:   make([]*Node, 0),
		Id:       uuid.New(),
	}
	*c = append(*c, node)
	return node
}

// Retorna o primeiro nó encontrado
func (c *conjunto) Get(rotulo string) *Node {
	idx := slices.IndexFunc(*c, func(x *Node) bool {
		return x.Rotulo == rotulo
	})
	return (*c)[idx]
}

func (c conjunto) String() string {
	return c.String() // ???
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
	for _, n := range pai.Filhos {
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
		return strings.Compare(a.Rotulo, b.Rotulo)
	})
	var (
		rotulos = make([]string, len(clone))
		matriz  = make([][]int, len(clone))
	)
	for i := range matriz {
		matriz[i] = make([]int, len(clone))
	}
	for i := range clone {
		rotulos[i] = clone[i].Rotulo
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

func (c conjunto) VerificarArvore(raiz *Node, considerarSubgrafo bool) (
	GrafoArvore bool, ArvoreBinaria bool, ArvoreCheia bool, ArvoreCompleta bool,
) {
	// não fica claro no documento quais tipos de árvores ele quer.
	// talvez árvore binária seja redundante.

	GrafoArvore = true
	ArvoreBinaria = true
	ArvoreCheia = true
	ArvoreCompleta = true

	if len(c) == 0 {
		return
	}
	var (
		nivelProximoNodes []*Node
		nivelAtualNodes   = []*Node{raiz}
		visitados         []*Node
		niveis            [][]*Node
	)

	for len(nivelAtualNodes) > 0 {
		nivelProximoNodes = []*Node{}
		var graus []int
		for _, n := range nivelAtualNodes {
			// grafo é ciclico. não é uma árvore.
			if slices.Index(visitados, n) > -1 {
				return false, false, false, false
			}
			if len(n.Filhos) != 2 && len(n.Filhos) != 0 {
				ArvoreCheia = false
			}
			nivelProximoNodes = append(nivelProximoNodes, n.Filhos...)
			visitados = append(visitados, n)
			graus = append(graus, len(n.Filhos))
		}
		niveis = append(niveis, nivelAtualNodes)
		nivelAtualNodes = nivelProximoNodes
	}
	if !considerarSubgrafo {
		// existe algum nó que não é alcançável através do nó raiz. grafo não é uma árvore
		if len(c) != len(visitados) {
			return false, false, false, false
		}

	}

	nivelMaximo := len(niveis) - 1

	for i := range nivelMaximo {
		if len(niveis[i]) != 1<<i {
			ArvoreCompleta = false
			return
		}
	}

	// demorei um dia para fazer isso.
	for i := range niveis {
		contiguo := true
		ultimoNivel := i+1 == len(niveis)
		for j := range niveis[i] {
			l := len(niveis[i][j].Filhos)
			if contiguo {
				contiguo = contiguo && l == 2
			} else if l > 0 {
				if !ultimoNivel {
					ArvoreCompleta = false
					return
				}
			}
		}
	}

	return
}

// TODO: testar
// To chutando que essa função deve retornar a quantidade total de graus do digrafo
// retorna o número máximo do grau de todos os vértices
func (c conjunto) GrafoGrau() int {
	var valenciaTotalDigrafo = 0
	for i := 0; i < len(c); i++ {
		grauEntrada, grauSaida := c[i].Grau()
		valenciaTotalDigrafo += grauEntrada + grauSaida
	}
	return valenciaTotalDigrafo
}
func (a *Node) GrafoGrau() int {
	if a == nil {
		panic("O nó não pode ser nulo")
	}
	var grauEntradaTotal, grauSaidaTotal int = a.Grau()
	return grauEntradaTotal + grauSaidaTotal
}

// verifica se grafo é completo
func (c conjunto) VerificarCompleto() bool {
	for i := range c {
		for j := range c {
			if i == j {
				continue
			}
			if slices.Index(c[i].Filhos, c[j]) < 0 {
				return false
			}
		}
	}
	return true
}

// verifica se grafo possuí ao menos um vértice com um laço
func (c conjunto) VerificarLacos() bool {
	for _, n := range c {
		if slices.Index(n.Filhos, n) > -1 {
			return true
		}
	}
	return false
}

// TODO: testar
// verifica se grafo é um grafo simples
func (c conjunto) VerificarSimples() bool {
	for _, node := range c {
		contagem := map[*Node]int{node: 1}
		for _, filho := range node.Filhos {
			contagem[filho]++
		}
		for _, v := range contagem {
			if v > 1 {
				return false
			}
		}
	}
	return true
}
