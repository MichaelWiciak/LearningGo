package main

import (
	"fmt"

	"rsc.io/quote"
)

func HelloWorld() string {
	return "Hello World!"
}

func Quote() string {
	return quote.Go()
}

func main() {
	fmt.Println(HelloWorld())
	fmt.Println(Quote())
}
