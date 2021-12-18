package main

import (
	"encoding/json"
	"fmt"
)

type info struct {
	Markets []struct {
		Name  string
		Price int
	}
}

type status struct {
	Name string
}

func main() {
	text := `{"markets":[{"name":"1", "price":100},{"name":"2","price":101},{"name":"3", "price":102}]}`
	var infos info
	json.Unmarshal([]byte(text), &infos)
	fmt.Println(infos)
	fmt.Println(infos.Markets)
	fmt.Println(infos.Markets[0].Name)

	var st = status{Name: "privet"}
	fmt.Println(st.Name)

	for i := range infos.Markets {
		fmt.Println(i, infos.Markets[i].Name, infos.Markets[i].Price)
	}
}
