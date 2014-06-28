package main

import (
	"euler/common"
	"fmt"
)

func main() {
	sum := 0
	for i := 2; i < 2000000; i++ {
		if common.IsPrime(float64(i)) {
			sum += i
		}
	}

	fmt.Println(sum)
}
