package main

import "fmt"

func main() {
	destination := make([]int, 10, 10)
	source := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copy(destination, source)
	fmt.Println(destination)
}

func showAllElemtnts(values ...int) {
	for _, value := range values {
		fmt.Println(value)
	}
}
