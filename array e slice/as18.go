package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Informe um número inteiro positivo: ")
	fmt.Scanln(&n)

	if n <= 0 {
		fmt.Println("O número deve ser inteiro positivo!")
		return
	}
	primos := make([]int, 0)
	num := 2

	for len(primos) < n {
		ehPrimo := true
		raiz := int(math.Sqrt(float64(num)))
		for i := 2; i <= raiz; i++ {
			if num%i == 0 {
				ehPrimo = false
				break
			}
		}
		if ehPrimo {
			primos = append(primos, num)
		}
		num++
	}
	fmt.Println("Números primos:", primos)
}
