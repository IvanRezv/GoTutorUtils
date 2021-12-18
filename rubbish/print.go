package main

import (
	"fmt"
)

func main() {
	var name, age = "Art", 32
	c := fmt.Sprintf("My name is %s. And i'm %d years old.", name, age)
	fmt.Println(c)
}
