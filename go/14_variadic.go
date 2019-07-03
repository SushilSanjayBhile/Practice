package main

import "fmt"

func myVariadic(numbers ... int){
	fmt.Println(numbers)
}

// Printf / Println are also variadic functions

func main(){
	myVariadic(1,2)
	myVariadic(10,20,30,40,50)
	myVariadic(100,200,300)
}