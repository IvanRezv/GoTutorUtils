package main

import "fmt"

func main() {
	// slice
	// append = add

	slice := []string{
		"van",
		"serg",
	}
	fmt.Println(slice)

	letters := []string{"a", "b", "c"}
	letters[0] = "1"
	letters = append(letters, "d")
	letters = append(letters, "d", "o", "g")
	fmt.Println(letters)

	createSlice := make([]string, 3)
	createSlice[0] = "1"
	createSlice[1] = "2"
	createSlice[2] = "3"
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))
	createSlice = append(createSlice, "4")
	fmt.Println(createSlice)
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))
	createSlice = append(createSlice, "4", "5", "7")
	fmt.Println(createSlice)
	fmt.Println(len(createSlice))
	fmt.Println(cap(createSlice))
}

func arrays() {
	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	// a[2] = "ok"
	fmt.Println(a)
	fmt.Println(a[1])

	numbers := [...]int{1, 2, 3}
	numbers[2] = 4
}
