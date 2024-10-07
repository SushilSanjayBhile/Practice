package codes

import (
	"sort"
)

func CanSumToLargest() {
	input := []int{3, 5, -1, 8, 12}
	canSumToLargest(input)

	input = []int{3, 5, -1, 8, 100}
	canSumToLargest(input)

	input = []int{5, 3, -1, 8, 12}
	canSumToLargest(input)
}

func canSumToLargest(nums []int) bool {
	sort.Ints(nums)
	target := nums[len(nums)-1]
	nums = nums[:len(nums)-1]
	return combinations(nums, target)
}

func findSum(arr []int, num int) bool {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum == num
}

// Function to get all combinations of a given array
func combinations(arr []int, num int) bool {
	if len(arr) == 0 {
		return false
	}
	result := findSum(arr, num)
	if result {
		return result
	}
	var finalresult = false
	for idx := 0; idx < len(arr); idx++ {
		var newArr []int
		newArr = append(newArr, arr[:idx]...)
		newArr = append(newArr, arr[idx+1:]...)
		finalresult = combinations(newArr, num)
		if finalresult {
			break
		}
	}
	return finalresult
}
