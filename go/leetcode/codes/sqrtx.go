package codes

import "fmt"

func mySqrt(x int) int {
	start, end := 0, x+1

	for start < end {
		mid := start + (end-start)/2

		if mid*mid > x {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start - 1
}

func MySqrt() {
	fmt.Println("8:", mySqrt(8))
	fmt.Println("4:", mySqrt(4))
}
