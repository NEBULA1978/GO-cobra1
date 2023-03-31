package main

import (
	"fmt"
)

func main() {
	SayHello(message())
}

// public function
func SayHello(message string) {
	fmt.Println(message)
}

// private message with explicit return
func message() (message string) {
	message = "Hello World"
	return
}
