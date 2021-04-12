package main

import "fmt"

func recurse(n int) int{
	if n == 0{
		return 0
	}
	return n + recurse(n-1)
}

func main(){
	fmt.Println(recurse(3))
	fmt.Println(recurse(10))
	fmt.Println(recurse(8))
}
