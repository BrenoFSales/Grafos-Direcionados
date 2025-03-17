// TODO: fazer os tratamentos de erros;
// TODO: Por tudo em português (Daniel H. to Daniel A.)?;
package io

import (
	"fmt"
)

// Método para a instância de uma Matriz A(m x n);
func GenMatrix() [][]int {
	// m são as linhas, n as colunas;
	var m, n int
	fmt.Print("Número de Linhas da matriz: ")
	fmt.Scan(&m)
	fmt.Print("Número de Colunas da matriz: ")
	fmt.Scan(&n)

	// Instancia de uma matrix A(m x n) vazia e preenchimento da mesma com inputs do usuário;
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			var data int
			fmt.Printf("Elemento %d da linha %d: ", j+1, i+1)
			fmt.Scan(&data)
			matrix[i][j] = data
		}
	}

	return matrix
}

// Método para mostrar uma matriz A(m x n) qualquer passada como argumento;
func GetMatriz(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Print("| ")
		for j := 0; j < len(matrix); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Print("|\n")
	}
	// Retorno já formatado;
}
