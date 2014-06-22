package main

import (
	"euler/common"
	"fmt"
)

// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

// What is the 10 001st prime number?

func main() {
	primes := []int{2}

	for i := 3; len(primes) <= 10001; i += 2 {
		if common.IsPrime(float64(i)) {
			primes = append(primes, i)
		}
	}

	fmt.Println(primes[10000])
}
