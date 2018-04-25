package main

import "fmt"

type Person struct {
	name string
	age  int8
}

func main() {
	message, closing := "hello, mars", "i bid you adieu"
	fmt.Println(message, closing)

	var pointer *string = &message
	fmt.Println(pointer, *pointer)

	var dude = Person{"Byron", 29}
	fmt.Println(dude, dude.name, dude.age)

}
