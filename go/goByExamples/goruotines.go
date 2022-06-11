package main

import "fmt"
import "time"

func p(no string){
	for i:=0; i < 30; i++{
		fmt.Println("No is:", no, i)
	}
}

func main(){
	go p("6")
	go p("5")
	go func(){
		for i := 10; i < 15; i++{
			fmt.Println(i)
		}
	}()
	go p("10")
	time.Sleep(time.Second * 1)
}
