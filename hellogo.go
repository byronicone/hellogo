package main

import "fmt"
import "strconv"

type Person struct {
	name string
	age  int
	id   int
}

const (
	A = iota
	B
)

func describe(person Person) int {
	fmt.Println(createMessage(person.name, person.age))
	return person.id
}

func createMessage(name string, age int) string {
	return name + " is " + strconv.Itoa(age) + " years old!"
}

func main() {

	var dude = Person{"Byron", 29, A}
	_ = describe(dude)

}
