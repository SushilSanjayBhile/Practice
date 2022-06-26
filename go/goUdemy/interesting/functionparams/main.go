package main

import (
	"fmt"
)

func parameterDataType1(x, y int) {
	fmt.Println("As x & y have same datatype, we need to declare datatype only once.")
}

func parameterDataType2(x, y int, z string) {
	fmt.Println("we can declare params like this as well.")
}

func main() {
	parameterDataType1(1, 2)
	parameterDataType2(1, 2, "String")
}
