package main

import "fmt"

func main() {
	sum := 0
	for sum < 1000 {
		sum += 10
		fmt.Println(sum)
	}
}
