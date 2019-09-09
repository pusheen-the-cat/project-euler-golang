package util

import (
	"testing"

	"github.com/pusheen-the-cat/project-euler-golang/projecteuler/golang/util"
)

func TestIsPalindrome_StringTrue(t *testing.T) {
	tests := []string{
		"aba", "abba", "abcba", "",
	}
	for _, test := range tests {
		input := util.String{
			S: test,
		}
		if !input.IsPalindrome() {
			t.Errorf("For \"%s\", IsPalindrome returned false.", test)
		}
	}
}

func TestIsPalindrome_StringFalse(t *testing.T) {
	tests := []string{
		"abc", "ac", "taco cat", "else",
	}
	for _, test := range tests {
		input := util.String{
			S: test,
		}
		if input.IsPalindrome() {
			t.Errorf("For \"%s\", IsPalindrome returned true.", test)
		}
	}
}

func TestIsPalindrome_IntTrue(t *testing.T) {
	tests := []int{
		0, 1, 121, 1221, 12321,
	}
	for _, test := range tests {
		input := util.Int{
			I: test,
		}
		if !input.IsPalindrome() {
			t.Errorf("For \"%d\", IsPalindrome returned false.", test)
		}
	}
}

func TestIsPalindrome_IntFalse(t *testing.T) {
	tests := []int{
		12, 123, 123210,
	}
	for _, test := range tests {
		input := util.Int{
			I: test,
		}
		if input.IsPalindrome() {
			t.Errorf("For \"%d\", IsPalindrome returned true.", test)
		}
	}
}
