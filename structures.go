package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Developer struct {
	Human
	workAge      int
	techonlogies []Technology
}

type Technology struct {
	name           string
	specialization string
}

func main() {
	hleb := Developer{Human: Human{name: "hleb", age: 17}, workAge: 3, techonlogies: []Technology{Technology{name: "hleb"}}}
	fmt.Println(hleb)
}
