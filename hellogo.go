package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	name      string
	nicknames []string
	age       int
	id        int
}

const (
	A = iota
	B
)

type Printer func(Person) int

func main() {
	var dude = Person{"Byron", []string{"Bigwig", "The Cleaner"}, 29, A}
	fmt.Println("Simple func:")
	describe(dude, simpleMessageFunc(dude))
	fmt.Println("\nVerbose func:")
	describe(dude, verboseMessageFunc(dude))
}

func describe(person Person, do Printer) {
	_ = do(person)
}

func printNameAndAge(name string, age int) {
	strAge := strconv.Itoa(age)
	fmt.Println(name + " is " + strAge + " years old!")
}

func simpleMessageFunc(person Person) (p Printer) {
	return func(person Person) (id int) {
		printNameAndAge(person.name, person.age)
		return person.id
	}
}

func verboseMessageFunc(person Person) (p Printer) {
	return func(person Person) (id int) {
		name := person.name
		age := person.age
		printNameAndAge(person.name, person.age)
		nextAge := strconv.Itoa(age + 1)
		fmt.Println(name + " is going to be " + nextAge + " next year.  How ancient...")
		nicknames := person.nicknames
		numNicks := len(nicknames)
		if numNicks > 0 {
			fmt.Printf("Also known as")
			for i := 0; i < numNicks; i++ {
				fmt.Printf(" " + nicknames[i])
				if i != numNicks-1 {
					fmt.Printf(",")
				} else {
					fmt.Println()
				}
			}
		}
		return person.id
	}
}
