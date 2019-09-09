package util

import (
	"math"
)

var (
	primes = make([]int, 0)
)

//GenPrimesUpTo generates []int of Primes up to max
func GenPrimesUpTo(max int) []int {
	for i := GenNextPrime(); i <= max; i = GenNextPrime() {
		GenNextPrime()
	}
	result := make([]int, len(primes))
	copy(result, primes)
	return result
}

//GenNextPrime generates the next prime number
func GenNextPrime() int {
	if len(primes) == 0 {
		primes = append(primes, 2)
	} else {
		num := primes[len(primes)-1] + 1
		for !isPrime(num) {
			num++
		}
		primes = append(primes, num)
	}
	return primes[len(primes)-1]
}

func isPrime(l int) bool {
	maxToCheck := int(math.Sqrt(float64(l)))

	for _, e := range primes {
		if e > maxToCheck {
			return true
		}

		if l%e == 0 {
			return false
		}
	}

	return true
}
