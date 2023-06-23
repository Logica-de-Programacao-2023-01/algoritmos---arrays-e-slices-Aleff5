package main

import "fmt"

func main() {
	array := [6]float64{7.2, 6.4, 9.6, 7.8, 1.0, 2.1}

	var numero float64
	fmt.Print("Informe um número: ")
	fmt.Scanln(&numero)

	for i := 0; i < len(array); i++ {
		array[i] += numero
	}

	fmt.Println("Array resultante:", array)
}
