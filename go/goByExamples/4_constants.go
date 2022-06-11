package main

import "math"

const name string = "bhile"

func main() {
	println(name)

	//name = "sushil"			: changing value of const name, cannot be done

	const pi = 3.146
	println(pi)

	const salary = 10000
	println(salary)

	const inhand = salary - 5000
	println(inhand)

	println(math.Sin(inhand))
}
