package main

import (
	"github.com/wenkaler/leetcode/utils"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	makeListNode([]int{0, 1})
	testCases := []struct {
		name     string
		listOne  *ListNode
		listTwo  *ListNode
		response *ListNode
	}{
		{
			name:     "check 3x3",
			listOne:  makeListNode([]int{2, 4, 3}),
			listTwo:  makeListNode([]int{5, 6, 4}),
			response: makeListNode([]int{7, 0, 8}),
		},
		{
			name:     "check 0",
			listOne:  makeListNode([]int{0}),
			listTwo:  makeListNode([]int{0}),
			response: makeListNode([]int{0}),
		},
		{
			name:     "check 1x1",
			listOne:  makeListNode([]int{5}),
			listTwo:  makeListNode([]int{5}),
			response: makeListNode([]int{0, 1}),
		},
		{
			name:     "check 0x2",
			listOne:  makeListNode([]int{}),
			listTwo:  makeListNode([]int{7, 3}),
			response: makeListNode([]int{7, 3}),
		},
		{
			name:     "check 2x1",
			listOne:  makeListNode([]int{9, 8}),
			listTwo:  makeListNode([]int{1}),
			response: makeListNode([]int{0, 9}),
		},
		{
			name:     "check 1x2",
			listOne:  makeListNode([]int{1}),
			listTwo:  makeListNode([]int{9, 9}),
			response: makeListNode([]int{0, 0, 1}),
		},
		{
			name:     "check 7x8",
			listOne:  makeListNode([]int{6, 9, 7, 6, 3, 1, 8, 3, 1}),
			listTwo:  makeListNode([]int{2, 6, 3, 6, 1, 4, 2, 1, 9, 1}),
			response: makeListNode([]int{8, 5, 1, 3, 5, 5, 0, 5, 0, 2}),
		},
		{
			name:     "check 30x3",
			listOne:  makeListNode([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
			listTwo:  makeListNode([]int{5, 6, 4}),
			response: makeListNode([]int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := addTwoNumbers(tc.listOne, tc.listTwo)
			for tc.response != nil || got != nil {
				if got == nil || tc.response == nil {
					t.Fatalf("failed value got(%v) exp(%v)", got, tc.response)
				}
				if tc.response.Val != got.Val {
					t.Fatalf("failed value got(%v) exp(%v)", got, tc.response)
				}
				tc.response, got = tc.response.Next, got.Next
			}
		})
	}
}

var l1, l2 = makeListNode(utils.GenSliceInt(10, 100)), makeListNode(utils.GenSliceInt(6, 100))

func BenchmarkAddTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}

func BenchmarkAddTwoNumbers2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addTwoNumbers2(l1, l2)
	}
}

func BenchmarkAddTwoNumbers3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addTwoNumbers3(l1, l2)
	}
}
