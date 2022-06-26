package main

import (
	"fmt"
)

// in return we don't have to write variable names if return values are defined as follows
func addSub(x int, y int) (add int, sub int) {
	add = x + y
	sub = x - y
	return
}

func main() {
	add, sub := addSub(14, 7)
	fmt.Println(add, sub)
}
