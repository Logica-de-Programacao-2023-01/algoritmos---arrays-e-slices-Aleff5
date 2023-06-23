package main

import "fmt"

func main() {
	array := [10]float64{3.2, 6.7, 2.1, 8.9, 4.5, 7.3, 1.8, 9.6, 5.4, 6.1}
	slice := make([]float64, 0)

	for _, num := range array {
		if num > 5 {
			slice = append(slice, num)
		}
	}
	fmt.Println("Novo slice:", slice)
}
