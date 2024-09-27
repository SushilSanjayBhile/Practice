package codes

import (
	"fmt"
	"sort"
)

func maxProductDifference1(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return (nums[len(nums)-1] * nums[len(nums)-2]) - (nums[0] * nums[1])
}

func maxProductDifference2(nums []int) int {
	sort.Ints(nums)

	left := nums[0] * nums[1]
	rigth := nums[len(nums)-1] * nums[len(nums)-2]

	res := rigth - left

	return res
}

func MaxProductDifference() {
	nums := []int{5, 6, 2, 7, 4}
	fmt.Println(maxProductDifference1(nums))
	fmt.Println(maxProductDifference2(nums))

	nums = []int{4, 2, 5, 9, 7, 4, 8}
	fmt.Println(maxProductDifference1(nums))
	fmt.Println(maxProductDifference2(nums))
}
