package codes

import "fmt"

func search(nums []int, target int) int {
	startIndex := 0
	currentIndex := len(nums) / 2
	endIndex := len(nums)

	for {
		if currentIndex == startIndex && nums[currentIndex] != target {
			return -1
		}
		if nums[currentIndex] == target {
			return currentIndex
		} else if nums[currentIndex] > target {
			endIndex = currentIndex
			currentIndex = int((startIndex + currentIndex) / 2)
		} else if nums[currentIndex] < target {
			startIndex = currentIndex
			currentIndex = int((startIndex + endIndex) / 2)
		}
	}
}

type MyStruct struct {
	Property string
}

func Search() {
	nums := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 20
	fmt.Println(search(nums, target))

	nums = []int{-1, 0, 3, 5, 9, 12}
	target = 9
	fmt.Println(search(nums, target))

	s1 := MyStruct{
		Property: "hey",
	}

	val := s1.Property
	fmt.Println("s1.Property has been set", val)

}
