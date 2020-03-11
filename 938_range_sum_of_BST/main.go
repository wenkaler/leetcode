package main

import "fmt"

func main() {
	x := CreateTreeNode([]int{10,5,6,18,3,7,15})
	y := CreateTreeNode([]int{10,5,15,3,7,13,18,1,6})
	fmt.Println(x.RangeSum(7, 15))
	fmt.Println(y.RangeSum(6, 10))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func CreateTreeNode(list []int) *TreeNode{
	f := new(TreeNode)
	for i, x := range list{
		if i == 0 {
			f.Val = x
			continue
		}
		f.insertNode(x)
	}
	return f
}

func (t *TreeNode)insertNode(x int){
	if t.Val > x {
		if t.Left != nil{
			t.Left.insertNode(x)
			return
		}
		t.Left = &TreeNode{
			Val:   x,
		}
	}else{
		if t.Right != nil{
			t.Right.insertNode(x)
			return
		}
		t.Right = &TreeNode{
			Val:   x,
		}
	}
}

func (t *TreeNode) RangeSum(l,r int) int{
	resp := 0
	if t.Val >= l && t.Val <= r {
		resp += t.Val
	}
	if t.Left != nil && t.Val > l{
			resp += t.Left.RangeSum(l, r)
	}
	if t.Right != nil && t.Val < r{
			resp += t.Right.RangeSum(l, r)
	}
	return resp
}

