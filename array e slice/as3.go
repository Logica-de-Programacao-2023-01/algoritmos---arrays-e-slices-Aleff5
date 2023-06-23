package main

import "fmt"

func main() {

	array := [4]float64{13.5, 8.5, 5.5, 7.5}

	produto := 1.0
	for _, value := range array {
		produto *= value
	}

	fmt.Println(produto)
}
