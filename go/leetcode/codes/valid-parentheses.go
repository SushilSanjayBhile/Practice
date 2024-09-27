package codes

import "fmt"

func isValid(s string) bool {
	fmt.Println()
	fmt.Println(s)
	if len(s) == 0 || len(s)%2 == 1 {
			return false
		}
		
		pairs := map[rune]rune{
			'(': ')',
			'{': '}',
			'[': ']',
		}
		stack := []rune{}
	
		for _, r := range s {
			if _, ok := pairs[r]; ok {
				stack = append(stack, r)
			} else if len(stack) == 0 || pairs[stack[len(stack)-1]] != r {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	
		return len(stack) == 0
	}

	func IsValid(){
		fmt.Println(isValid("]"))
		fmt.Println(isValid("()"))
		fmt.Println(isValid("()[]{}"))
		fmt.Println(isValid("(]"))
		fmt.Println(isValid("([])"))
	
	
	}