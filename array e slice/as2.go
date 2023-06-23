package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}

	remover := 2

	slice = append(slice[:remover], slice[remover+1:]...)

	fmt.Println(slice)
}
