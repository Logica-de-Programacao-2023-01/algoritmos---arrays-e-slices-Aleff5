package main

import "fmt"

func main() {
	matriz := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	var linha, coluna int
	fmt.Print("Informe o índice de linha: ")
	fmt.Scanln(&linha)
	fmt.Print("Informe o índice de coluna: ")
	fmt.Scanln(&coluna)

	if linha >= 0 && linha < len(matriz) && coluna >= 0 && coluna < len(matriz[linha]) {
		valor := matriz[linha][coluna]
		fmt.Printf("O valor armazenado na posição [%d][%d] é: %d\n", linha, coluna, valor)
	} else {
		fmt.Println("Índices inválidos!")
	}
}
