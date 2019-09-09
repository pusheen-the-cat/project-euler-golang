/*
Smallest multiple

Problem 5
2520 is the smallest number that can be divided by each of the numbers
from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all
of the numbers from 1 to 20?
*/
package main

import (
	"fmt"
	"math"

	"github.com/pusheen-the-cat/project-euler-golang/projecteuler/golang/util"
)

func main() {
	target := 20
	factors := make(map[int]int, 0)
	primes := util.GenPrimesUpTo(target)

	for i := 1; i <= target; i++ {
		currFactors := findFactors(primes, i)
		for k, v := range currFactors {
			if factors[k] < v {
				factors[k] = v
			}
		}
	}

	result := 1
	for k, v := range factors {
		result = result * int(math.Pow(float64(k), float64(v)))
	}
	fmt.Println(factors)
	fmt.Println(result)
}

func findFactors(primes []int, i int) map[int]int {
	factors := make(map[int]int, 0)
	for _, prime := range primes {
		for i%prime == 0 {
			factors[prime] = factors[prime] + 1
			i = i / prime
		}
	}

	return factors
}
