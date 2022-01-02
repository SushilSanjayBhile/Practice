package main

import "fmt"

func main(){
	msg := make(chan int)
	go func() { msg <- 100 }()
	go func() { msg <- 101 }()
	go func() { msg <- 103 }()
	go func() { msg <- 102 }()

	m := <-msg
	fmt.Println(m)
	fmt.Println(<-msg)
}
