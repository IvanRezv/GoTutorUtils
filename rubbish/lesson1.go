package main

import "fmt"

var (
	i    int
	text = "test"
	zar  bool
)

func main() {
	bob := Cats{"Bob", 7, 0.87}
	fmt.Println("Bob function is", bob.test())
	fmt.Println("Abobus", text)
}

type Cats struct {
	name     string
	age      int
	happines float64
}

func (cat *Cats) test() float64 {
	return float64(cat.age) * cat.happines
}
