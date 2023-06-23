package main

import "fmt"

func main() {

	array := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	var valor int
	fmt.Print("Informe um valor: ")
	fmt.Scanln(&valor)

	encontrado := false
	for _, num := range array {
		if num == valor {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Printf("O valor %d está presente no array.\n", valor)
	} else {
		fmt.Printf("O valor %d não está presente no array.\n", valor)
	}
}
