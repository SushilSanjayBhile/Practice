package main

import "fmt"

func main(){
	numbers := []int{1,2,3,4,5}
	sum := 0

	for _,number := range numbers{
		sum += number
	}

	fmt.Println("sum", sum)

	for index := range numbers{
		fmt.Println("index", index)
	}

	maps := map[int]string{4:"sushil",1:"sanjay",2:"bharati",3:"suuhas",5:"arati"}

	for i,val := range maps{
		fmt.Println("\n\n",i, val)
		fmt.Printf("%d -> %s", i, val)			//using printf function here
	}

	fmt.Println("\n\n")

	name := "sushil"
	for i, val := range name{
		fmt.Println(i,val)			//prints unicode code points
	}
}