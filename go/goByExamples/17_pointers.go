package main

import "fmt"

func withoutPointer(val string){
	val = "changed"
}

func withPointer(val *string){
	*val = "changed"
}

func main(){
	mainString := "ORIGINAL"
	fmt.Println(mainString)

	withoutPointer(mainString)
	fmt.Println(mainString)

	withPointer(&mainString)
	fmt.Println(mainString)

	fmt.Println(&mainString)
}