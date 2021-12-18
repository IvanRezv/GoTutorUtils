package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p Point) method() {
	fmt.Println("finish")
}

func main() {
	pointsMap := map[string]Point{
		"b": {
			Y: 0,
			X: 0,
		},
	}
	otherPointsMap := make(map[int]Point)
	// var oneMorePointsMap map[string]Point

	pointsMap["a"] = Point{
		X: 1,
		Y: 2,
	}
	fmt.Println(pointsMap)
	fmt.Println(pointsMap["a"])
	otherPointsMap[1] = Point{X: 1, Y: 2}
	fmt.Println(otherPointsMap)

	var oneMorePointsMap map[string]Point
	if oneMorePointsMap == nil {
		oneMorePointsMap = map[string]Point{
			"a": {Y: 1, X: 2},
		}
	}
	fmt.Println(oneMorePointsMap["a"].X)
	fmt.Println(oneMorePointsMap["a"].Y)
	oneMorePointsMap["a"].method()

	key := "a"
	value, ok := oneMorePointsMap[key]
	if ok {
		fmt.Printf("key=%s exist in map\n", key)
		fmt.Println(value)
	} else {
		fmt.Println("No shuch value")
	}
}
