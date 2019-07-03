package main

import "fmt"


func myFunc(a,b int)(int, int){
	return (a * b), (a + b)
}
func main(){
	a,b := myFunc(2,5)
	fmt.Println(a,b)

	_,c := myFunc(2,5) 
	fmt.Println(c)
}