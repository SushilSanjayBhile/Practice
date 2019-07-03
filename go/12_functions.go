package main

import "fmt"


func sum2(a int, b int) int {
	return a+b
}

func sum3(a, b, c int) {
	fmt.Println(a + b + c)
}

func main(){
	fmt.Println(sum2(1,2))
	sum3(1,2,3)
}
