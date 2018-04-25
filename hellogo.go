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

func main() {
	message, closing := "hello, mars", "i bid you adieu"
	fmt.Println(message, closing)

	var pointer *string = &message
	fmt.Println(pointer, *pointer)

	var dude = Person{"Byron", 29, A}
	var lass = Person{"Jane", 47, B}
	fmt.Println(dude, dude.name, dude.age, dude.id)
	fmt.Println(lass, lass.name, lass.age, lass.id)

}
