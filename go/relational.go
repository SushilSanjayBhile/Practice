package main

import "fmt"

func main(){
	fmt.Println("program to demonstrate relational operations in Go:")

	var (
		num1 = 19
		num2 = 25
	)

	fmt.Println("Numbers are:", num1, num2)

	fmt.Println(num1, ">", num2, ":", num1 > num2)
	fmt.Println(num1, "<", num2, ":", num1 < num2)
	fmt.Println(num1, ">=", num1, ":", num1 >= num1)
	fmt.Println(num1, "<=", num2, ":", num1 <= num2)
	fmt.Println(num1, "==", num2, ":", num1 == num2)
	fmt.Println(num1, "!=", num2, ":", num1 != num2)


}