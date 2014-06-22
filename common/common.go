package common

import "math"

func GCD(m, n int) int {
	var tmp int

	for m != 0 {
		tmp = m
		m = n % m
		n = tmp
	}

	return n
}

func LCM(m, n int) int {
	return m / GCD(m, n) * n
}

func IsPrime(n float64) (t bool) {
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
