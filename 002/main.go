package main

import "fmt"

var m = make(map[int]int)

func fib(n int) int {
	if n <= 1 {
		return n
	} else if val, ok := m[n]; !ok {
		return fib(n-1) + fib(n-2)
	} else {
		return val
	}
}

func main() {
	largest := 0
	i := 0
	sum := 0

	for largest < 4000000 {
		if largest%2 == 0 {
			sum += largest
		}

		i++

		largest = fib(i)
	}

	fmt.Println(sum)
}
