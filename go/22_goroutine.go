package main

import "fmt"

func temp(name string){
	for i:=0; i<3; i++{
		fmt.Println(name,":",i)
	}
}

func main(){
	temp("sushil")

	go temp("sanjay")

	go func(str string){
		fmt.Println(str)
	}("bhile")

	// needs some pause till execution of asynchronous go-routine completes
	for i:=0;i<100000000;i++{}
	// fmt.Scanln()
}
