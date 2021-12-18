package main

import "fmt"

func main() {
	s, s2, s3 := test()
	s4 := test1()
	fmt.Println(s, s2, s3, s4)
}

func test() (a, b, c string) {
	a = "hello"
	c = "abobus"
	b = "world"
	return
}

func test1() string {
	return "empty"
}
