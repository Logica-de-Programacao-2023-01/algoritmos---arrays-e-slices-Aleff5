package main

import "fmt"

func main() {
	linhas := 3
	colunas := 2

	matriz := make([][]int, linhas)
	for i := range matriz {
		matriz[i] = make([]int, colunas)
	}

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			fmt.Printf("Informe o valor do elemento [%d][%d]: ", i, j)
			fmt.Scanln(&matriz[i][j])
		}
	}

	fmt.Println("A matriz informada é:")
	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			fmt.Printf("%d\t", matriz[i][j])
		}
		fmt.Println()
	}
}
