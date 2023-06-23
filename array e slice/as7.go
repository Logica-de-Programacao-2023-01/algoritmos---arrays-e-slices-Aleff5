package main

import "fmt"

func main() {

	slice := make([]int, 0, 5)

	var numero int
	fmt.Print("Informe um número inteiro: ")
	fmt.Scanln(&numero)

	presente := false
	for _, num := range slice {
		if num == numero {
			presente = true
			break
		}
	}

	if !presente {
		slice = append(slice, numero)
	}

	fmt.Println("Slice resultante:", slice)
}
}
