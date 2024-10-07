package codes

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	splt := strings.Split(s, " ")

	for idx := len(splt) - 1; idx >= 0; idx-- {
		if splt[idx] != " " && splt[idx] != "" {
			return len(splt[idx])
		}
	}
	return 0
}

func LengthOfLastWord() {
	fmt.Println("H", lengthOfLastWord("H"))
	fmt.Println("", lengthOfLastWord(""))
	fmt.Println("Hello World", lengthOfLastWord("Hello World"))
	fmt.Println("   fly me   to   the moon  ", lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println("luffy is still joyboy", lengthOfLastWord("luffy is still joyboy"))
}
