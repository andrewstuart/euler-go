// A Pythagorean triplet is a set of three natural numbers, a < b < c, for
// which,
//                              a^2 + b^2 = c^2

// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

package main

import "fmt"

//a=m^2-n^2, b=2mn, c=m^2+n^2

func main() {
	sum := 0
	var a, b, c int
	for m := 1; sum != 1000; m++ {
		for n := 1; n < m && sum != 1000; n++ {

			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n

			sum = a + b + c
		}
	}
	fmt.Println(a * b * c)
}
