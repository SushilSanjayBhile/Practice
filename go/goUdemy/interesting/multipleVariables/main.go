package main

import (
	"fmt"
)

func main() {
	var a, b = 1, "This is string"
	fmt.Printf("\nType: %T, Data: %d", a, a)
	fmt.Printf("\nType: %T, Data: %s", b, b)

	fmt.Println()

	var c, d, e = 1, 3.14, true
	fmt.Printf("\nType: %T, Data: %d", c, c)
	fmt.Printf("\nType: %T, Data: %f", d, d)
	fmt.Printf("\nType: %T, Data: %t", e, e)

	fmt.Println()

	var f, g, h int = 1, 2, 3
	fmt.Printf("\nType: %T, Data: %d", f, f)
	fmt.Printf("\nType: %T, Data: %d", g, g)
	fmt.Printf("\nType: %T, Data: %d", h, h)
}
