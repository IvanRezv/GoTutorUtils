package main

import "fmt"

func main () {
	keys := []string{}

	carta := map[string][]int{
		"1": {1,2},
		"2": {2,10},
	}

	for key := rang cacarta {
		keys = append(keys, key)
		fmt.Println(key)
	}

	fmt.Println(keys)
}