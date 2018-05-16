package main

import (
	"fmt"
	"time"
)

const (
	A = iota
	B
)

func main() {

	c := make(chan string)

	var dude = Person{"Byron", []string{"Bigwig", "The Cleaner"}, 29, A}

	go describe(dude, c, simpleMessageFunc(dude))
	go describe(dude, c, verboseMessageFunc(dude))
	for count := 0; count < 2; {
		select {
		case s, _ := <-c:
			fmt.Printf("Method: %s\n", s)
			count++
		default:
			fmt.Println("\nWaiting for thread to complete.")
			time.Sleep(5 * time.Nanosecond)
		}
	}
}
