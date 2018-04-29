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
	description, _ := createMessage(person.name, person.age)
	fmt.Println(description)
	return person.id
}

func createMessage(name string, age int) (message string, alternate string) {
	strAge := strconv.Itoa(age)
	nextAge := strconv.Itoa(age + 1)
	message = name + " is " + strAge + " years old!"
	alternate = name + " is going to be " + nextAge + " next year.  How ancient..."
	return
}

func main() {

	var dude = Person{"Byron", 29, A}
	_ = describe(dude)

}
