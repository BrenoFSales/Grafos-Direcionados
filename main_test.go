package main

import (
	"reflect"
	"testing"
)

func TestNode(t *testing.T) {
	conjunto := NovoConjunto()
	a := conjunto.NovoNode("1")
	b := conjunto.NovoNode("2")
	c := conjunto.NovoNode("2")

	a.Conectar(b)
	a.Conectar(c)

	resultado, err := a.Get(b.id)

	if err != nil {
		t.Fatal(err)
	}

	esperado := b

	if resultado != esperado {
		t.Fatalf("nó resultado não é o esperado. resultado: %p, esperado: %p", resultado, esperado)
	}

	a.Remover(b.id)

	resultado, err = a.Get(c.id)
	if err != nil {
		t.Fatal(err)
	}

	esperado = c

	if resultado != esperado {
		t.Fatalf("nó resultado não é o esperado. resultado: %p, esperado: %p", resultado, esperado)
	}

	c.Conectar(a)
	a.Conectar(a)
	a.Conectar(a)

	grauEntradaResultado, grauSaidaResultado := a.Grau()
	grauEntradaEsperado, grauSaidaEsperado := 3, 3

	if grauEntradaEsperado != grauEntradaResultado {
		t.Log(a)
		t.Fatalf(
			"grauEntradaEsperado (%d) != grauEntradaResultado (%d)",
			grauEntradaEsperado, grauEntradaResultado,
		)
	}
	if grauSaidaEsperado != grauSaidaResultado {
		t.Log(a)
		t.Fatalf(
			"grauSaidaEsperado (%d) != grauSaidaResultado (%d)",
			grauSaidaEsperado, grauSaidaResultado,
		)
	}
}

func TestConjunto(t *testing.T) {
	conjunto := NovoConjunto()

	var (
		a = conjunto.NovoNode("1")
		b = conjunto.NovoNode("2")
		c = conjunto.NovoNode("3")
		d = conjunto.NovoNode("4")
		e = conjunto.NovoNode("5")
		f = conjunto.NovoNode("6")
	)

	a.Conectar(b)
	a.Conectar(c)
	a.Conectar(d)
	a.Conectar(e)
	a.Conectar(f)

	b.Conectar(c)
	b.Conectar(d)
	b.Conectar(e)
	b.Conectar(f)

	c.Conectar(d)
	c.Conectar(e)
	c.Conectar(f)

	e.Conectar(a)

	f.Conectar(e)
	f.Conectar(f)
	f.Conectar(b)

	e.Conectar(a)

	t.Run("MatrizAdjacencia", func(t *testing.T) {
		resultado, _ := conjunto.MatrizAdjacencia()

		//   a  b  c  d  e  f
		// a 0  1  1  1  1  1
		// b 0  0  1  1  1  1
		// c 0  0  0  1  1  1
		// d 0  0  0  0  0  0
		// e 2  0  0  0  0  0
		// f 0  1  0  0  1  1
		esperado := [][]int{
			{0, 1, 1, 1, 1, 1},
			{0, 0, 1, 1, 1, 1},
			{0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 0},
			{2, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1},
		}

		if !reflect.DeepEqual(resultado, esperado) {
			t.Fatalf(
				"matriz de adjacência retornada não é a esperada.\nresultado: %#v\nesperado: %#v",
				resultado, esperado,
			)
		}
	})
	// t.Run("listaAdjacencia", func(t *testing.T) {
	// 	resultado := conjunto.ListaAdjacencia()
	//
	// 	esperado := [][]*Node{
	// 		{b, c, d, e, f},
	// 		{c, d, e, f},
	// 		{d, e, f},
	// 		{},
	// 		{a, a},
	// 		{e, f, b},
	// 	}
	//
	// 	if !reflect.DeepEqual(resultado, esperado) {
	// 		t.Fatalf(
	// 			"lista de adjacência retornada não é a esperada.\nresultado: %#v\nesperado: %#v",
	// 			resultado, esperado,
	// 		)
	// 	}
	// })

	// t.Run("VerticesGrau", func(t *testing.T) {
	// 	resultado := conjunto.VerticesGrau()
	// 	esperado := map[*Node]int{a: 5, b: 4, c: 3, d: 0, e: 2, f: 3}
	//
	// 	if !maps.Equal(resultado, esperado) {
	// 		t.Fatalf(
	// 			"vertices do grau retornados não são os esperados.\nresultado: %#v\nesperado: %#v",
	// 			resultado, esperado,
	// 		)
	// 	}
	// })

	t.Run("GrafoGrau", func(t *testing.T) {
		resultado := conjunto.GrafoGrau()
		esperado := 5

		if resultado != esperado {
			t.Fatalf(
				"grau de grafo retornado não é a esperada.\nresultado: %#v\nesperado: %#v",
				resultado, esperado,
			)
		}
	})
}

func TestVerificarArvore(t *testing.T) {
	conjunto := NovoConjunto()

	var (
		a = conjunto.NovoNode("a")
		b = a.NovoNode("b")
		c = a.NovoNode("c")

		d = b.NovoNode("d")
		e = b.NovoNode("e")

		f = c.NovoNode("f")
		g = c.NovoNode("g")

		_ = d.NovoNode("x1")
		_ = d.NovoNode("x2")
		_ = e.NovoNode("x3")
		_ = e.NovoNode("x4")

		_ = f.NovoNode("x5")
		_ = f.NovoNode("x6")
		_ = g.NovoNode("x7")
		h = g.NovoNode("h")
	)

	arvore, binaria, cheia, completa := conjunto.VerificarArvore(a)
	if !arvore {
		t.Log("verificação retornada não é a esperada: grafo é uma árvore")
		t.Fail()
	}
	if !binaria {
		t.Log("verificação retornada não é a esperada: árvore é binária")
		t.Fail()
	}
	if !cheia {
		t.Log("verificação retornada não é a esperada: árvore é cheia")
		t.Fail()
	}
	if !completa {
		t.Log("verificação retornada não é a esperada: árvore é completa")
		t.Fail()
	}
	if t.Failed() {
		t.FailNow()
	}

	conjunto.Remover(h.rotulo)
	arvore, binaria, cheia, completa = conjunto.VerificarArvore(a)
	if !arvore {
		t.Log("verificação retornada não é a esperada: grafo é uma árvore")
		t.Fail()
	}
	if !binaria {
		t.Log("verificação retornada não é a esperada: árvore é binária")
		t.Fail()
	}
	if cheia {
		t.Log("verificação retornada não é a esperada: árvore não é cheia")
		t.Fail()
	}
	if !completa {
		t.Log("verificação retornada não é a esperada: árvore é completa")
		t.Fail()
	}
	if t.Failed() {
		t.FailNow()
	}

	conjunto.Remover("x5")
	arvore, binaria, cheia, completa = conjunto.VerificarArvore(a)
	if !arvore {
		t.Log("verificação retornada não é a esperada: grafo é uma árvore")
		t.Fail()
	}
	if !binaria {
		t.Log("verificação retornada não é a esperada: árvore é binária")
		t.Fail()
	}
	if cheia {
		t.Log("verificação retornada não é a esperada: árvore não é cheia")
		t.Fail()
	}
	if completa {
		t.Log("verificação retornada não é a esperada: árvore não é completa")
		t.Fail()
	}
	if t.Failed() {
		t.FailNow()
	}

	conjunto = NovoConjunto()

	a = conjunto.NovoNode("a")
	c = a.NovoNode("c")

	f = c.NovoNode("f")
	g = c.NovoNode("g")

	_ = f.NovoNode("x5")
	_ = f.NovoNode("x6")
	_ = g.NovoNode("x7")
	_ = g.NovoNode("h")

	arvore, binaria, cheia, completa = conjunto.VerificarArvore(a)
	if !arvore {
		t.Log("verificação retornada não é a esperada: grafo é uma árvore")
		t.Fail()
	}
	if !binaria {
		t.Log("verificação retornada não é a esperada: árvore é binária")
		t.Fail()
	}
	if cheia {
		t.Log("verificação retornada não é a esperada: árvore ainda é cheia")
		t.Fail()
	}
	if completa {
		t.Log("verificação retornada não é a esperada: árvore não é completa")
		t.Fail()
	}
	if t.Failed() {
		t.FailNow()
	}

	conjunto = NovoConjunto()
	conjunto.NovoNode("a")
	conjunto.NovoNode("b")

	arvore, binaria, cheia, completa = conjunto.VerificarArvore(a)
	if arvore || binaria || cheia || completa {
		t.Log("verificação retornada não é a esperada: grafo não é uma árvore")
		t.Fail()
	}

	conjunto = NovoConjunto()
	a = conjunto.NovoNode("a")
	a.NovoNode("b").NovoNode("c").Conectar(a)

	arvore, binaria, cheia, completa = conjunto.VerificarArvore(a)
	if arvore || binaria || cheia || completa {
		t.Log("verificação retornada não é a esperada: grafo não é uma árvore")
		t.Fail()
	}
}

func TestVerificarCompleto(t *testing.T) {

	conjunto := NovoConjunto()
	vertices := []*Node{
		conjunto.NovoNode("1"), conjunto.NovoNode("2"), conjunto.NovoNode("3"),
		conjunto.NovoNode("4"), conjunto.NovoNode("5"), conjunto.NovoNode("6"),
		conjunto.NovoNode("7"),
	}
	for i := range vertices {
		for j := range vertices {
			if i == j {
				continue
			}
			a, b := vertices[i], vertices[j]
			a.Conectar(b)
			b.Conectar(a)
		}
	}
	resultado := conjunto.VerificarCompleto()
	esperado := true

	if resultado != esperado {
		t.Fatal("verificação retornada não é a esperada: grafo é completo")
	}

	vertices[0].Remover(vertices[1].id)

	resultado = conjunto.VerificarCompleto()
	esperado = false

	if resultado != esperado {
		t.Fatal("verificação retornada não é a esperada: grafo não é completo")
	}

}

func TestVerificarLacos(t *testing.T) {
	conjunto := NovoConjunto()

	a := conjunto.NovoNode("1")
	a.Conectar(a)

	resultado := conjunto.VerificarLacos()
	esperado := true

	if resultado != esperado {
		t.Fatal("verificação retornada não é a esperada: grafo possuí um laço")
	}

	conjunto = NovoConjunto()
	conjunto.NovoNode("1").NovoNode("1")

	resultado = conjunto.VerificarLacos()
	esperado = false

	if resultado != esperado {
		t.Fatal("verificação retornada não é a esperada: segundo grafo não possuí um laço")
	}

}

func TestVerificarSimples(t *testing.T) {
	conjunto := NovoConjunto()

	a := conjunto.NovoNode("1")
	b := a.NovoNode("2")
	a.Conectar(b)

	resultado := conjunto.VerificarSimples()
	esperado := false

	if resultado != esperado {
		t.Fatal("verificação retornada não é a esperada: grafo não é simples")
	}

	conjunto = NovoConjunto()

	a = conjunto.NovoNode("1")
	b = a.NovoNode("2")
	c := b.NovoNode("3")
	c.Conectar(a)

	resultado = conjunto.VerificarSimples()
	esperado = true

	if resultado != esperado {
		t.Fatal("verificação retornada não é a esperada: grafo é simples")
	}

	a.Conectar(a)

	resultado = conjunto.VerificarSimples()
	esperado = false

	if resultado != esperado {
		t.Fatal("verificação retornada não é a esperada: grafo não é simples")
	}

}

func TestImprimir(t *testing.T) {
	// sono....
}
