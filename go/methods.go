package main
import "fmt"

type number struct {
	a, b int
}

func (r number) add() int {
	return r.a + r.b
}

func (r number) sub() int {
	return r.a - r.b
}

func main(){
	num := number{11, 12}
	p:= fmt.Println

	p(num)
	p("addition:", num.add())
	p("subtraction:", num.sub())
}
