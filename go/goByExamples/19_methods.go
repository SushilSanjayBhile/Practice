package main

import "fmt"

type student struct{
	rollno int
	name string
}

func (s *student) getRollNo() int{
	return s.rollno
}

func (s *student) getName() string{
	return s.name
}

func main(){
	obj1 := student{name: "sushil", rollno: 16107}

	fmt.Println("whole obj\t", obj1)
	fmt.Println("address of object\t", &obj1)
	fmt.Println("address of name inside of object\t",&obj1.name)
	fmt.Println("address of rollNo inside of object\t",&obj1.rollno)
	// fmt.Println("address of getName() inside of object\t",&obj1.getName)
	// fmt.Println("address of getRollNo() inside of object\t",&obj1.rollNo)
	fmt.Println("obj1.getRollNo()\t",obj1.getRollNo(),"\t")
	fmt.Println("obj.getName()\t",obj1.getName())

	obj2 := student{name: "arati", rollno: 16121}

	fmt.Println("\n",obj2)
	fmt.Print(obj2.getRollNo(),"\t")
	fmt.Println(obj2.getName())
}