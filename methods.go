package main

import "fmt"

type square struct {
	side int
}

func (s square) area() int {
	return s.side * s.side
}

type person struct {
	name string
	age  int
}

func (person person) greet(quote string) {
	fmt.Printf("Hello from %s, he is %v, %s's quote: %s\n", person.name, person.age, person.name, quote)
}

type Cow struct {
	name   string
	height int
}

func (cow *Cow) rename(newName string) {
	cow.name = newName
}

func main() {
	obj := square{side: 2}
	fmt.Println(obj.side)
	fmt.Println(obj.area())
	fmt.Println("----------")
	bob := person{name: "Bob", age: 17}
	bob.greet("Never give up!!")
	fmt.Println("----------")
	cow := Cow{name: "Burenka", height: 185}
	fmt.Printf("old name: %s\n", cow.name)
	cow.rename("zalupka")
	fmt.Printf("new name: %s\n", cow.name)

}
