package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := make([]int, 0)

	for _, num := range array {
		if num%2 == 0 {
			slice = append(slice, num)
		}
	}

	fmt.Println("Novo slice:", slice)
}
