package main

import (
	"fmt"

	"leetcode.com/codes"
)

// "fmt"

func main() {
	input := []int{3, 5, -1, 8, 12}
	input = []int{3, 5, -1, 8, 100}
	// input = []int{5, 3, -1, 8, 12}
	result := codes.CanSumToLargest(input)
	fmt.Println(result) // Output: false

	// codes.MergeTwoLists()

	// codes.IsValid()

	// fmt.Println(codes.LongestCommonPrefix([]string{"flower", "flow", "flight"}))
	// fmt.Println(codes.LongestCommonPrefix([]string{"dog", "racecar", "car"}))
	// fmt.Println(codes.LongestCommonPrefix([]string{""}))

	// fmt.Println(codes.TwoSum([]int{2, 7, 11, 15}, 9))

	// fmt.Println(codes.RomanToInt("III"))
	// fmt.Println(codes.RomanToInt("LVIII"))
	// fmt.Println(codes.RomanToInt("MCMXCIV"))

	// ds.QueueOperations()
	// ds.StackOperations()
	// codes.Permutation()

	// fmt.Println(codes.GcdOfStringsRecursion("ABCABC", "ABC"))

	// s := "hello"
	// fmt.Println(codes.ReverseVowels(s))

	// candies := []int{2, 3, 5, 1, 3}
	// extraCandies := 3
	// fmt.Println(codes.KidsWithCandies(candies, extraCandies))

	// fmt.Println(codes.RomanToInt("III"))
	// fmt.Println(codes.RomanToInt("MCMXCIV"))
	// fmt.Println(codes.RomanToInt("LVIII"))

	// fmt.Println("Palindrome function:")
	// x := 121
	// fmt.Println(x, codes.IsPalindrome(x))
	// x = 122
	// fmt.Println(x, codes.IsPalindrome(x))
}
