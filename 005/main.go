package main

import "fmt"

func gcd(m, n int) int {
	var tmp int

	for m != 0 {
		tmp = m
		m = n % m
		n = tmp
	}

	return n
}

func lcm(m, n int) int {
	return m / gcd(m, n) * n
}

func main() {
	start := 1

	for i := 1; i <= 20; i++ {
		start = lcm(start, i)
	}

	fmt.Println(start)
}
