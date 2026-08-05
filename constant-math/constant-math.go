package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	const x = 1.0 / 3.0
	fmt.Println(x)
	fmt.Println(float32(x))
	fmt.Println(float64(x))

	fmt.Println(math.Sin(n))
}
