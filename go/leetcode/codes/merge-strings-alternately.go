package codes

import "fmt"

func mergeAlternately(word1 string, word2 string) string {
	var newStr string = ""
	var totalLength = 0

	if len(word1) > len(word2) {
		totalLength = len(word2)
	} else {
		totalLength = len(word1)
	}

	for idx := 0; idx < totalLength; idx++ {
		newStr += string(word1[idx])
		newStr += string(word2[idx])
	}

	if len(word1) > len(word2) {
		for idx := len(word2); idx < len(word1); idx++ {
			newStr += string(word1[idx])
		}
	} else if len(word2) > len(word1) {
		for idx := len(word1); idx < len(word2); idx++ {
			newStr += string(word2[idx])
		}
	}

	return newStr
}

func MergeAlternately() {
	fmt.Println(mergeAlternately("abc", "pqr"))
	fmt.Println(mergeAlternately("ab", "pqrs"))
	fmt.Println(mergeAlternately("abcd", "pq"))
	fmt.Println(mergeAlternately("cdf", "a"))
}
