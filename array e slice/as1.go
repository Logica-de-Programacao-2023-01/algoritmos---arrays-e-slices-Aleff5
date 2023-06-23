package main

import "fmt"

func main() {

	array := [3]int{0, 1, 2}

	soma := 0

	for _, num := range array {
		soma += num
	}

	fmt.Println("A soma dos valores no array é:", soma)
}
