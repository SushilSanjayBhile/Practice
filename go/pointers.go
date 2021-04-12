package main

import "fmt"

func callByValue(i int) {
	i += 1
}

func callByReference(i *int) {
	*i += 1
}

func main(){
	p := fmt.Println

	i := 1
	p(i)

	callByValue(i)
	p(i)

	callByReference(&i)
	p(i)

	callByReference(&i)
	p(i)
}
