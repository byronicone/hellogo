package main

import "fmt"

func main() {
	message, closing := "hello, mars", "i bid you adieu"
	fmt.Println(message, closing)

	var pointer *string = &message
	fmt.Println(pointer)
}
