package main

import (
	"fmt"
	"math"
)

func main() {
	x := divide(-2147483648, -1)
	fmt.Println(x)
}

func divide(dividend int, divisor int) int {
	var counter int
	d := divisor
	di := dividend
	if di < 0 {
		di = -di
	}
	if d < 0 {
		d = -d
	}
	for di-d >= 0 {
		di -= d
		counter++
	}
	if (divisor < 0 || dividend < 0) && (divisor > 0 || dividend > 0) {
		return -counter
	}
	if counter == int(math.Pow(2, 31)) {
		counter--
	}
	return counter
}
