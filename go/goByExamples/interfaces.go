package main

import "fmt"

const pi float64 = 3.146
var p = fmt.Println

type geometry interface{
	area() float64
	perim() float64
}

type rect struct {
	w,h int
}

type circle struct {
	r int
}

func (r rect) area() float64{
	p("area of rectangle: width * height")
	return float64(r.w * r.h)
}

func (r rect) perim() float64 {
	p("perimeter of rectangle: 2 * (width + height)")
	return float64(2 * (r.w + r.h))
}

func (c circle) area() float64 {
	p("area of circle: 2 * pi * radius * radius")
	return float64(2) * pi * float64(c.r) * float64(c.r)
}

func (c circle) perim() float64 {
	p("perimeter of circle: 2 * pi * radius")
	return float64(2) * pi * float64(c.r)
}

func geo(g geometry){
	p(g)
	p(g.area())
	p(g.perim())
}

func main(){
	r := rect{2, 1}
	geo(r)

	c:= circle{3}
	geo(c)
}
