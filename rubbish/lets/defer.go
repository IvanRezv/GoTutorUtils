package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("adas")
	defer fmt.Println("jdjsa")
	fmt.Println("hello")
}
