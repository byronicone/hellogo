package main

import "fmt"

type Person struct {
	name string
	age  int8
	id   int
}

const (
	A = iota
	B
)

func greet(person Person) {
	fmt.Println(person.name, person.age)
}

func main() {

	var dude = Person{"Byron", 29, A}
	greet(dude)

}
