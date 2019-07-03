package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println("whole array: ", a)

	a[0], a[2] = 200, 300
	fmt.Println("whole array: ", a)
	fmt.Println("a[2]: ", a[2])
	fmt.Println("len: ", len(a))

	b:= [5]int{10, 20, 30, 40, 50}
	fmt.Println(b)

	var threeD [3][4][5]int

	for i:=1; i<4; i++{
		for j:=1; j<5; j++{
			for k:=1; k<6; k++{
				threeD[i-1][j-1][k-1] = i * j * k
			}
		}
	}

	fmt.Println(threeD)
}
