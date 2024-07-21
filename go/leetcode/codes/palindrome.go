package codes

func IsPalindrome(x int) bool {
	num := x
	var palindrome int
	for i := 0; num > 0; i++ {
		mod := num % 10
		num = num / 10
		palindrome = palindrome*10 + mod
	}

	return palindrome == x
}
