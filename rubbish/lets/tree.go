package main

import "fmt"

func main() {
	a := 0
	for a < 1000 {
		i := isTest(2)
		switch {
		case i == 1:
			fmt.Println("a = 2")
		case i == 2:
			fmt.Println("a = 3")
		default:
			fmt.Println(fmt.Sprintln("unknown a. a=%d", a))
		}
		// -----------=======================================--------------------------
		if i == 1 {
			fmt.Println("a = 2")
		} else if i == 2 {
			fmt.Println("a = 3")
		} else if i == 3 {
			fmt.Println("a = 4")
		} else {
			fmt.Println(fmt.Sprintf("unknown a. a=%d", a))
		}
		a++
	}
}

func isTest(a int) int {
	if a == 2 {
		return 1
	} else if a == 3 {
		return 2
	}
	return 3
}
