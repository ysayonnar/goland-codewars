package main

import "fmt"

func main() {
	//var defaultMap map[int]string
	//mapByMake := make(map[string]string, 3)
	mapByliteral := map[string]int{"Vasya": 20, "hleb": 17}
	mapByliteral["zalupa"] = 20
	mapByliteral["hleb"] = 10000
	fmt.Println(mapByliteral)
	fmt.Printf("Map length: %v\n", len(mapByliteral))
}
