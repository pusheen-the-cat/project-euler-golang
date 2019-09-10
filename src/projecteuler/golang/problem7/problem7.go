/*
10001st prime

Problem 7
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/
package main

import (
	"fmt"

	"github.com/pusheen-the-cat/project-euler-golang/projecteuler/golang/util"
)

func main() {
	for i := 1; i <= 10000; i++ {
		util.GenNextPrime()
	}
	fmt.Println(util.GenNextPrime())
}
