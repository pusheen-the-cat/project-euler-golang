/*
Largest palindrome product

Problem 4
A palindromic number reads the same both ways.
The largest palindrome made from the product of two 2-digit numbers
is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
package main

import (
	"fmt"

	"github.com/pusheen-the-cat/project-euler-golang/projecteuler/golang/util"
)

func main() {
	max := 0
	input := 999

	for i := input; i > 0; i-- {
		for j := input; j > 0; j-- {
			if i*j < max {
				continue
			}
			if i*j > max {
				temp := util.Int{
					I: i * j,
				}
				if temp.IsPalindrome() {
					max = i * j
				}
			}
		}
	}

	fmt.Println(max)
}
