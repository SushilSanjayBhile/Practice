package codes

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	h := len(haystack)
	n := len(needle)
	for i := 0; i < h-n+1; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
func StrStr() {
	fmt.Println("sadbutsad", "sad", strStr("sadbutsad", "sad"))
	fmt.Println("leetcode", "leeto", strStr("leetcode", "leeto"))
	fmt.Println("hello", "ll", strStr("hello", "ll"))
	fmt.Println("a", "a", strStr("a", "a"))
}
