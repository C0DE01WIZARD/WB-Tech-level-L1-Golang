package main

import (
	"fmt"
	
)

func main() {
	var first int32
	// create two big nums
	first = 19980435006
	second := big.NewFloat(20400761021.3214)

	// process addition
	add := new(big.Float)
	fmt.Printf("%f + %f = ", first, second)
	fmt.Println(add.Add(first, second))

	// process subtraction
	sub := new(big.Float)
	fmt.Printf("%f - %f = ", first, second)
	fmt.Println(sub.Sub(first, second))

	// process division
	div := new(big.Float)
	fmt.Printf("%f / %f = ", first, second)
	fmt.Println(div.Quo(second, first))
}