package codes

import (
	"fmt"
	"strings"
)

func isVowel(ch string) bool {
	var vowels = map[string]string{"a": "", "e": "", "i": "", "o": "", "u": ""}
	if _, ok := vowels[strings.ToLower(ch)]; ok {
		return true
	}
	return false
}

func reverseVowels(s string) string {
	fmt.Println(s)
	ss := []rune(s)

	var idx, jdx = 0, len(s) - 1
	var flag1, flag2 = false, false
	for idx < jdx {
		if ok := isVowel(string(s[idx])); ok {
			flag1 = true
		} else {
			idx++
		}
		if ok := isVowel(string(s[jdx])); ok {
			flag2 = true
		} else {
			jdx--
		}

		if flag1 && flag2 {
			ss[idx], ss[jdx] = ss[jdx], ss[idx]
			idx, jdx = idx+1, jdx-1
			flag1, flag2 = false, false
		}

	}
	return string(ss)
}

func ReverseVowels() {
	fmt.Println(reverseVowels("IceCreAm"))
	fmt.Println(reverseVowels("leetcode"))

}
