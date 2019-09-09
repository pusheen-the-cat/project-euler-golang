package util

import "strconv"

//PalindromeChecker checks whether IsPalindrome
type PalindromeChecker interface {
	IsPalindrome() bool
}

//String is a wrapper of string
type String struct {
	S string
}

//Int is a wrapper around int
type Int struct {
	I int
}

// IsPalindrome checks if string is a palindrome
func (in String) IsPalindrome() bool {
	s := in.S
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// IsPalindrome checks if int is a palindrome
func (in Int) IsPalindrome() bool {
	s := String{
		S: strconv.Itoa(in.I),
	}
	return s.IsPalindrome()
}
