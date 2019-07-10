package main

import "fmt"

func main(){
	fmt.Println("Program to demonstrate arithmetic operaions in Go:\n")

	var (
		num1 = 10
		num2 = 2
	)

	fmt.Println("Numbers are:", num1, num2)

	fmt.Println("Addition is:", num1 + num2)
	fmt.Println("Subtraction is:", num1 - num2)
	fmt.Println("Multiplication is:",num1 * num2)
	fmt.Println("Division is:",num1 / num2)
	fmt.Println("Modulo is:",num1 % num2)
}