package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Evans")
	fmt.Println(message)
}