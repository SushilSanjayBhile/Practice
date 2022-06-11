package main

import "fmt"

func myFunc() func() int{
	i := 1
	return func() int{
		i++
		return i
	}
}
func main(){
	fmt.Println(myFunc(),"\n\n")

	var funcVar = myFunc()

	fmt.Println(funcVar())
	fmt.Println(funcVar())
	fmt.Println(funcVar())


	fmt.Println("\n\nagain calling myFunc and assigning it to new variable ensures that state is unique. so this is stateful function call.")
	newVar := myFunc()

	fmt.Println(newVar())
	fmt.Println(newVar())
	fmt.Println(newVar())
}