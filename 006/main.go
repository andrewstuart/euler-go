package main

import "fmt"

// The sum of the squares of the first ten natural numbers is,

// 12 + 22 + ... + 102 = 385
// The square of the sum of the first ten natural numbers is,

// (1 + 2 + ... + 10)2 = 552 = 3025
// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025  385 = 2640.

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

func sumSquare(i ...int) int {
	s := 0
	for _, n := range i {
		s += n * n
	}
	return s
}

func squareSum(i ...int) int {
	s := 0
	for _, n := range i {
		s += n
	}
	return s * s
}

func diff(i ...int) int {
	return squareSum(i...) - sumSquare(i...)
}

func main() {
	a := make([]int, 0, 100)

	for i := 1; i <= 100; i++ {
		a = append(a, i)
	}

	fmt.Println(diff(a...))
}
