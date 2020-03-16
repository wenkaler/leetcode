package main

import (
	"github.com/wenkaler/leetcode/utils"
	"testing"
)

var list = utils.GenSliceInt(100000, 100)

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum(list, 21333)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum2(list, 21333)
	}
}

func BenchmarkTwoSum3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum3(list, 21333)
	}
}

func TestTwoSum(t *testing.T) {
	var testCases = []struct {
		name string
		list []int
		t    int
		resp []int
	}{
		{
			name: "first",
			list: []int{2, 7, 11, 15},
			t:    9,
			resp: []int{0, 1},
		}, {
			name: "second",
			list: []int{0, 7, 11, 0},
			t:    0,
			resp: []int{0, 3},
		}, {
			name: "third",
			list: []int{-2, -1, 7, 11, 0},
			t:    18,
			resp: []int{2, 3},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := twoSum(tc.list, tc.t)
			if len(tc.resp) != len(r) {
				t.Fatalf("failed got(%v) want(%v)", r, tc.resp)
			}
			for i := range tc.resp {
				if tc.resp[i] != r[i] {
					t.Fatalf("failed got(%v) want(%v)", r, tc.resp)
				}
			}
		})
	}
}
