package main

import (
	"fmt"
	"strconv"
	"time"
)

//String reverser
func rev(s string) string {
	a := make([]byte, len(s))

	for i := len(s); i > 0; i-- {
		a[len(s)-i] = s[i-1]
	}

	return string(a)
}

func isPalindrome(n int) bool {
	nStr := strconv.Itoa(n)

	nRev := rev(nStr)

	return nStr == nRev
}

func main() {
	largest := 0

	t := time.Now()

	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			test := i * j

			if isPalindrome(test) && test > largest {
				largest = test
			}
		}
	}

	fmt.Println(time.Since(t))

	fmt.Println(largest)
}
