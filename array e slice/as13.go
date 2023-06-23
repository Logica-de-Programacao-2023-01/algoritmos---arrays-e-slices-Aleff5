package main

import "fmt"

func main() {
	array := [7]int{2, 4, 6, 8, 10, 12, 14}
	var numero int
	
	fmt.Print("Informe um número: ")
	fmt.Scanln(&numero)

	array[0] += numero
	array[len(array)-1] += numero

	fmt.Println("Array resultante:", array)
}
