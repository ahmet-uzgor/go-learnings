package main

import (
	"fmt"

	"greetings"
)

func main() {
	message:= greetings.Hello("Ahmet")
	fmt.Println(message)
}