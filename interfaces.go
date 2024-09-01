package main

import "fmt"

type Runner interface {
	Run() string
}

type Person struct {
	name string
}

type Duck struct {
	name string
}

func (p *Person) Run() string {
	return fmt.Sprintf("%s person runs", p.name)
}

func (d *Duck) Run() string {
	return fmt.Sprintf("%s duck run", d.name)
}

func (d *Duck) Fly() string {
	return fmt.Sprintf("%s duck fly", d.name)
}

func (p *Person) WriteCode() string {
	return fmt.Sprintf("%s person write code", p.name)
}

func makeRun(r Runner) {
	fmt.Println(r.Run())
}

func TypeAccessIfElse(runner Runner) {
	fmt.Printf("Type: %T, Value: %v\n", runner, runner)
	if person, ok := runner.(*Person); ok {
		fmt.Println(person.WriteCode())
	}
	if duck, ok := runner.(*Duck); ok {
		fmt.Println(duck.Fly())
	}
}

func TypeAccessSwitch(runner Runner) {
	fmt.Printf("Type: %T, Value: %v\n", runner, runner)
	switch v := runner.(type) {
	case *Person:
		fmt.Println(v.WriteCode())
	case *Duck:
		fmt.Println(v.Fly())
	}
}

func main() {
	// три способа работы с интерфейсными штуками

	//1 -- ваще хуета
	var runner Runner
	person := Person{name: "Hleb"}
	runner = &person
	fmt.Println(runner.Run())
	TypeAccessIfElse(runner)
	TypeAccessSwitch(runner)

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	//2 -- более менее
	runner = &Duck{name: "Bobby"}
	fmt.Println(runner.Run())
	TypeAccessIfElse(runner)
	TypeAccessSwitch(runner)

	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	//3 - самый крутой, тут не нужно создавать runner, просто передать по ссылке
	newDuck := Duck{name: "Jonny"}
	makeRun(&newDuck)
	TypeAccessSwitch(&newDuck)
	TypeAccessIfElse(&newDuck)
}
