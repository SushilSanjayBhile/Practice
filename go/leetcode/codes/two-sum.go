package codes

func TwoSum(nums []int, target int) []int {
	for idx := 0; idx < len(nums); idx++ {
		for jdx := idx + 1; jdx < len(nums); jdx++ {
			if nums[idx]+nums[jdx] == target {
				return []int{idx, jdx}
			}
		}
	}
	return []int{}
}
