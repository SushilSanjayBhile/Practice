package main

import "fmt"

func main(){
	msg:= make(chan int, 4)

	msg <- 100
	msg <- 110
	msg <- 120
	msg <- 130

	p := fmt.Println
	p(<-msg)
	p(<-msg)
	p(<-msg)
	p(<-msg)
}
