package main

import (
	"euler/common"
	"fmt"
)

func main() {
	start := 1

	for i := 1; i <= 20; i++ {
		start = common.LCM(start, i)
	}

	fmt.Println(start)
}
