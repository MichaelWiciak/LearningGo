package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func messagePrinting(message string) {
	message, err := greetings.Hello(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	messagePrinting("Gladys")
	messagePrinting("Sam")

	names := []string{"Steve", "Bob", "Alice"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
	
}