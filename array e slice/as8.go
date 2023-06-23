package main

import "fmt"

func main() {

	slice := []string{"banana", "maçã", "uva", "laranja", "banana", "abacaxi", "banana", "morango"}

	var valor string
	fmt.Print("Informe um valor: ")
	fmt.Scanln(&valor)

	novoSlice := make([]string, 0, len(slice))
	for _, elemento := range slice {
		if elemento != valor {
			novoSlice = append(novoSlice, elemento)
		}
	}

	fmt.Println("Slice resultante:", novoSlice)
}
