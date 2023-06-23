package main

import "fmt"

func main() {
	array := [5]int{10, 5, 6, 9, 15}

	slice := make([]int, 0)
	for _, num := range array {
		if num%3 == 0 {
			slice = append(slice, num)
		}
	}
	fmt.Println("Novo slice:", slice)
}
