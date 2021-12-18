package main

import "fmt"

func main() {

}

func isnill() {
	var snil []int
	fmt.Println(snil, len(snil), cap(snil))
	if snil == nil {
		fmt.Print("is nill")
	}
}

func animals() {
	animalsArr := [3]string{
		"dog",
		"cat",
		"giraffe",
	}

	a := animalsArr[0:2]
	fmt.Println(a)

	b := animalsArr[1:3]
	fmt.Println(b)

	b[0] = "123"
	fmt.Println(a)
	fmt.Println(animalsArr)

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	t := s[5:]
	fmt.Println(t)
}
