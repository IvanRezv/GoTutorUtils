package main

import "fmt"

func main() {

	const trash string = "trash"
	textHello := "Hello"
	textName := "Name"
	i := 10

	var h float32
	h = 10.3

	next := float32(i) + h

	fmt.Println(fmt.Sprintf("%T %T %T %T", textHello, next, textName, trash))

	var input float64
	fmt.Scanf("%f", &input)
	output := input + 1
	fmt.Println(output)
}
