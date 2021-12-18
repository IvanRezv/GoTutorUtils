package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	text := []string{"be", "test", "hello"}
	fmt.Println(text[1])
	for _, i := range text {
		if i == "test" {
			continue
		}
		fmt.Println(i)
	}

	// find files

	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
