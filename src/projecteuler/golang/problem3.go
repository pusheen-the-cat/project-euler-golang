/**
 * Largest prime factor
 *
 * Problem 3
 * The prime factors of 13195 are 5, 7, 13 and 29.
 *
 * What is the largest prime factor of the number 600851475143 ?
 */
package main

import (
	"fmt"
	"projecteuler/golang/util"
)

func main() {
	target := 600851475143
	var currPrime int
	for target > 1 {
		currPrime = util.GenNextPrime()
		if target%currPrime == 0 {
			target = target / currPrime
		}
	}
	fmt.Println(currPrime)
}
