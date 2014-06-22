package main

import (
	"euler/common"
	"fmt"
	"math"
)

const BIG float64 = 600851475143

func main() {
	var i float64 = math.Floor(math.Sqrt(BIG))

	if math.Remainder(i, 2) == 0 {
		i--
	}

	for ; i > 0; i -= 2 {
		if math.Remainder(BIG, i) == 0 && common.IsPrime(i) {
			fmt.Println(i)
			break
		}
	}
}
