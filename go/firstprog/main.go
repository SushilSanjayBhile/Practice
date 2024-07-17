package main

import "fmt"
import "firstmodule/package1"
import "reflect"

type comp struct {
	a int 
}

type comp1 struct {
	// a []int
	b map[string]string
}

func main(){
	fmt.Println("Hello")
	package1.MyPrint("sushil")
	var h int
	h = h +1
	fmt.Println(h)
	fmt.Println(reflect.TypeOf(h))

	var s11 = comp{1}
	var s12 = comp{2}
	fmt.Println(s11, s12)
	fmt.Println(s11==s12)
	s12 = comp{1}
	fmt.Println(s11==s12)

	var s21 = comp1{[]int{1}}
	var s22 = comp1{[]int{2}}
	fmt.Println(s21, s22)
	fmt.Println(s21==s22)

}
