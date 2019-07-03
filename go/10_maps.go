package main

import "fmt"

func main(){
	product := make(map[string]int)

	product["mobile"] = 10000
	product["laptop"] = 45000

	fmt.Println("all: ",product)
	fmt.Println("mobile:", product["mobile"])

	product["mobile"] = 12999
	fmt.Println("updated mobile:", product["mobile"])

	v1:= product["laptop"]
	fmt.Println("variable assignment", v1)

	fmt.Println("lenght:", len(product))

	delete(product,"mobile")
	fmt.Println("deleted mobile and updated map: ", product)

	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}