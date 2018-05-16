package main

import (
	"fmt"
	"strconv"
)

type Printer func(Person) string

func describe(person Person, c chan string, do Printer) {
	str := do(person)
	c <- str
}

func simpleMessageFunc(person Person) (p Printer) {
	return func(person Person) (method string) {
		printNameAndAge(person.name, person.age)
		return "simple"
	}
}

func verboseMessageFunc(person Person) (p Printer) {
	return func(person Person) (method string) {
		printNameAndAge(person.name, person.age)
		printCommentary(person.age + 1)
		printNicknames(person.nicknames)
		return "verbose"
	}
}

func printNameAndAge(name string, age int) {
	strAge := strconv.Itoa(age)
	fmt.Println(name + " is " + strAge + " years old!")
}

func printCommentary(age int) {
	var comment string
	switch {
	case age >= 1 && age < 10:
		comment = "Growing up so fast!"
	case age >= 10 && age < 20:
		comment = "When I was your age..."
	case age >= 20 && age < 30:
		comment = "Stop and smell the roses."
	default:
		comment = "How ancient!"
	}
	nextAge := strconv.Itoa(age)
	fmt.Println("You will be " + nextAge + " next year.")
	fmt.Println(comment)
}

func printNicknames(nicknames []string) {
	numNicks := len(nicknames)
	if numNicks > 0 {
		fmt.Printf("Also known as")
		for i, nick := range nicknames {
			fmt.Printf(" " + nick)
			if i != numNicks-1 {
				fmt.Printf(",")
			} else {
				fmt.Println()
			}
		}
	}
}
