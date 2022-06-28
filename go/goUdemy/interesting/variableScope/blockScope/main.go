package main

import (
	"fmt"
)

// Block is defined by curly braces

func main() {
	var1 := 1
	fmt.Printf("\nVariable var1: %d have block scope restricted to main function", var1)

	{
		var1 := 2
		fmt.Printf("\nVariable var1: %d have block scope restricted to unnamed function/block. Because of this we can have var1 twice as their scope is different.", var1)
	}

	anotherBlock()
}

func anotherBlock() {
	var1 := 3
	fmt.Printf("\nVariable var1: %d have block scope restricted to anotherBlock function. Because of this we can have var1 again as their scope is different.", var1)
}
