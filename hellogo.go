package main

import "fmt"

const (
	A = iota
	B
)

func main() {

	var dude = Person{"Byron", []string{"Bigwig", "The Cleaner"}, 29, A}
	fmt.Println("Simple func:")
	describe(dude, simpleMessageFunc(dude))
	fmt.Println("\nVerbose func:")
	describe(dude, verboseMessageFunc(dude))
}
