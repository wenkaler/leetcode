package utils

import (
	"math/rand"
	"time"
)

func GenSliceInt(length, n int) []int {
	a := make([]int, length)
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		a[i] = rand.Intn(n)
	}
	return a
}
