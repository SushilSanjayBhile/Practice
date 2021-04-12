package main

import "fmt"

func nextint() func() int{
	i := 0
	return  func() int {
		i += 1
		return i
	}
}
func main(){
	ni := nextint()
	fmt.Println(ni())
	fmt.Println(ni())
	fmt.Println(ni())
}
