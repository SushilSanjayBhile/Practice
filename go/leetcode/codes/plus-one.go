package codes

import "fmt"

func plusOne(digits []int) []int {
	var finalDigit int64 = 0
	for _, v := range digits {
		finalDigit = finalDigit*10 + int64(v)
	}
	finalDigit += 1
	fmt.Println(finalDigit)

	slc := []int{}
	for {
		if finalDigit <= 0 {
			break
		}
		slc = append([]int{int(finalDigit % 10)}, slc...)
		finalDigit = finalDigit / 10
	}
	return slc
}

func PlusOne() {
	fmt.Println("1, 2, 3:", plusOne([]int{1, 2, 3}))
	fmt.Println("4,3,2,1:", plusOne([]int{4, 3, 2, 1}))
	fmt.Println("9:", plusOne([]int{9}))
	fmt.Println("8,9:", plusOne([]int{8, 9}))
	fmt.Println("7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,4,0,0,6:", plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
}
