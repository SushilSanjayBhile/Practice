package main

import "fmt"

func main(){
	fmt.Println("Program to demonstrate logical operations in Go:")

	var (
		num1 = true
		num2 = false
	)

	fmt.Println(num1, "||", num2, ":", num1 || num2)
	fmt.Println(num1, "&&", num2, ":", num1 && num2)
	fmt.Println("!(", num1, "&&", num2, ") :", !(num1 && num2))
}