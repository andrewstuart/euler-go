package main

import (
	"fmt"
	"math"
)

const BIG float64 = 600851475143

func isPrime(n float64) (t bool) {
	if math.Remainder(n, 2) == 0 {
		return false
	}

	i := 3.0

	for ; !t && i < math.Sqrt(n); i += 2 {
		//t will become true and the loop will break and we'll return false (!t)
		t = (math.Remainder(n, i) == 0)
	}

	//If we never set t = true then we'll return true (!t)
	return !t
}

func main() {
	var i float64 = math.Floor(math.Sqrt(BIG))

	if math.Remainder(i, 2) == 0 {
		i--
	}

	for ; i > 0; i -= 2 {
		if math.Remainder(BIG, i) == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}
