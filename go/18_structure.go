package main

import "fmt"

type student struct{
	rollNo int
	name string
}

func main(){
	fmt.Println(student{1,"sushil"})

	obj := student{2,"arati"}
	fmt.Println(obj)

	fmt.Println(obj.name)
	fmt.Println(obj.rollNo)
	
	fmt.Println(&obj)
	fmt.Println(&obj.rollNo)
	fmt.Println(&obj.name)

	obj1 := student{rollNo: 3, name: "bharati"}
	fmt.Println(obj1)

	obj1.name = "sanjay"
	fmt.Println(obj1)
}