package main

import "fmt"

type Point struct {
	X int
	Y int
	S string
}

func main() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "hello",
	}
	p1.method()
}

func (p Point) method() {
	fmt.Println(p.X)
	fmt.Println(p.Y)
	fmt.Println(p.S)
}

func structs() {
	p1 := Point{
		X: 1,
		Y: 2,
		S: "asd",
	}
	fmt.Println(p1)
	fmt.Println(p1.X)
	p1.X = 123
	fmt.Println(p1)

	p2 := Point{
		X: 123,
	}
	fmt.Println(p2)

	g := &p1
	fmt.Println(g)
}

func pointers() {
	a := "hello world"
	b := 42
	fmt.Println(a)
	fmt.Println(b)

	p := &a
	fmt.Println(p)
	fmt.Println(*p)
	*p = "oh my"
	fmt.Println(a)

	b = 42
	g := &b
	*g = *g / 2
	fmt.Println(b)
}
