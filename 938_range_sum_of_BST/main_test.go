package main

import (
	"leetcode/utils"
	"math/rand"
	"testing"
)

func TestCreateTreeNode(t *testing.T) {
	testCases := []struct {
		name     string
		list     []int
		wantResp *TreeNode
	}{
		{
			name: "create tree list 8 elem",
			list: []int{10, 5, 3, 15, 7, 18},
			wantResp: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  15,
					Left: nil,
					Right: &TreeNode{
						Val:   18,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
		{
			name: "create tree list 2 elem",
			list: []int{10, 5},
			wantResp: &TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val:   5,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
		},
		{
			name: "create tree list 1 elem",
			list: []int{10},
			wantResp: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
		},
		{
			name:     "create tree list 0 elem",
			list:     []int{},
			wantResp: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			x := CreateTreeNode(tc.list)
			for _, v := range tc.list {
				a, b := tc.wantResp, x
				for a.Val != v {
					if v < a.Val {
						a = a.Left
						b = b.Left
					} else {
						a = a.Right
						b = b.Right
					}
				}
				if tc.wantResp.Val != x.Val {
					t.Fatalf("create tree node fail got(%v) want(%v)", x, tc.wantResp)
				}
			}
		})
	}
}

func TestTreeNode_RangeSum(t *testing.T) {
	testCases := []struct {
		name     string
		left     int
		right    int
		list     []int
		wantResp int
	}{
		{
			name:     "example 1 response",
			left:     7,
			right:    15,
			list:     []int{10, 5, 3, 15, 7, 18},
			wantResp: 32,
		}, {
			name:     "example 2 response",
			left:     6,
			right:    10,
			list:     []int{10, 5, 15, 3, 7, 13, 18, 1, 6},
			wantResp: 23,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tree := CreateTreeNode(tc.list)
			got := tree.RangeSum(tc.left, tc.right)
			if got != tc.wantResp {
				t.Fatalf("failed sum got(%d) want(%d)", got, tc.wantResp)
			}
		})
	}
}

func BenchmarkTreeNode_RangeSum(b *testing.B) {
	list := utils.GenSliceInt(30, 100)
	x := CreateTreeNode(list)
	l, r := rand.Intn(50), rand.Intn(50)+50
	for i := 0; i < b.N; i++ {
		x.RangeSum(l, r)
	}
}

func BenchmarkCreateTreeNode(b *testing.B) {
	list := utils.GenSliceInt(30, 100)
	for i := 0; i < b.N; i++ {
		CreateTreeNode(list)
	}
}
