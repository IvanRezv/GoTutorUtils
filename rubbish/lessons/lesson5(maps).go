package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	carta := map[string][]int{
		"1": {1, 2},
		"2": {2, 10},
		"3": {1, 4},
	}
	for i, key := range carta {
		fmt.Println(i, key)
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
