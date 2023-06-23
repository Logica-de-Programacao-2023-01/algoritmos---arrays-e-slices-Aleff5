package main

import "fmt"

func main() {
	slice := []int{10, 5, 8, 2, 15, 3, 7, 12, 9, 6}
	minimo := slice[0]
	maximo := slice[0]
	for _, valor := range slice {
		if valor < minimo {
			minimo = valor
		}
		if valor > maximo {
			maximo = valor
		}
	}
	fmt.Println("Valor mínimo:", minimo)
	fmt.Println("Valor máximo:", maximo)
}
