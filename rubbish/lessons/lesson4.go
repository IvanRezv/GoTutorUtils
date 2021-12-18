package main

import "fmt"

func main() {
	// array := []int{1, 2, 3}

	array := make([]int, 10)

	fmt.Println(len(array))

	array[2] = 12
	array = append(array, 223)

	fmt.Println(len(array))
	for _, i := range array {
		fmt.Println(i)
	}
	fmt.Println(array)
}
