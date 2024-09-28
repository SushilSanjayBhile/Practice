package codes

import "fmt"

func searchInsert(nums []int, target int) int {
	idx := 0

	for _, v := range nums {
		if v >= target {
			return idx
		}
		idx++
	}

	return idx
}

func SearchInsert() {
	fmt.Println("nums = [3], target = 5", searchInsert([]int{3}, 5))
	fmt.Println("nums = [1,3,5,6], target = 5", searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println("nums = [1,3,5,6], target = 2", searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println("nums = [1,3,5,6], target = 7", searchInsert([]int{1, 3, 5, 6}, 7))
}
