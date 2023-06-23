package main

import "fmt"

func main() {
	var n int
	fmt.Print("Informe o tamanho do array: ")
	fmt.Scanln(&n)

	array := make([]int, n)
	fmt.Printf("Informe os %d elementos do array:\n", n)
	for i := 0; i < n; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scanln(&array[i])
	}
	ordenado := true
	for i := 0; i < n-1; i++ {
		if array[i] > array[i+1] {
			ordenado = false
			break
		}
	}
	if ordenado {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}
