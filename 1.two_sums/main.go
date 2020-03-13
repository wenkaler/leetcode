package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{-2, 11, 15}, 9))
}

func twoSum(list []int, t int) []int {
	var a []int
	if t == 0 {
		a = make([]int, len(list))
	} else {
		a = make([]int, t+1)
	}
	for i := 0; i < len(list); i++ {
		exp := list[i] - t
		if exp < 0 && list[i] > t || list[i] > t {
			continue
		} else if exp < 0 {
			exp = -exp
		}
		if list[a[list[i]]]+list[i] == t && a[exp] != i {
			return []int{a[exp], i}
		}
		a[exp] = i
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		_, prs := m[n]
		if prs {
			return []int{m[n], i}
		} else {
			m[target-n] = i
		}
	}
	return nil
}

func twoSum3(list []int, t int) []int {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == t {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
