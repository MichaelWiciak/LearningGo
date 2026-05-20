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
	messagePrinting("")
	
}