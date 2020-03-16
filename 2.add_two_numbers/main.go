package main

func main() {
	println(addTwoNumbers(makeListNode([]int{9, 8}), makeListNode([]int{1})).Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeListNode(x []int) *ListNode {
	var f *ListNode
	node := &ListNode{}
	for i := range x {
		node.Next = &ListNode{
			Val:  x[i],
			Next: nil,
		}
		node = node.Next

		if f == nil {
			f = node
		}
	}
	return f
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var f *ListNode
	node := &ListNode{}
	var d int
	for l1 != nil || l2 != nil {
		if l1 == nil {
			if d > 0 {
				node.Next = &ListNode{
					Val:  (l2.Val + d) % 10,
					Next: nil,
				}
				d = (l2.Val + d) / 10
				node = node.Next
				l2 = l2.Next
				continue
			}
			node.Next = l2
			break
		}
		if l2 == nil {
			if d > 0 {
				node.Next = &ListNode{
					Val:  (l1.Val + d) % 10,
					Next: nil,
				}
				d = (l1.Val + d) / 10
				node = node.Next
				l1 = l1.Next
				continue
			}
			node.Next = l1
			break
		}
		c := l1.Val + l2.Val + d
		node.Next = &ListNode{
			Val:  c % 10,
			Next: nil,
		}
		node = node.Next
		d = c / 10
		if f == nil {
			f = node
		}
		l1, l2 = l1.Next, l2.Next
	}
	if d > 0 {
		node.Next = &ListNode{
			Val:  d,
			Next: nil,
		}
	}
	return f
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	carry_over := 0
	var start *ListNode
	var lastNode *ListNode
	//Wrong condition: l1==nil, l2==nil, instead of the reverse one
	//Wrong condition: use AND, not OR
	for l1 != nil || l2 != nil {
		new_number := carry_over
		if l1 != nil {
			new_number += l1.Val
		}
		if l2 != nil {
			new_number += l2.Val
		}

		newListNode := &ListNode{Val: new_number % 10, Next: nil}
		carry_over = new_number / 10

		if start == nil {
			start = newListNode
			lastNode = newListNode
		} else {
			lastNode.Next = newListNode
			lastNode = newListNode
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	//FORGOT CARRY OVER
	if carry_over > 0 {
		newListNode := &ListNode{Val: carry_over, Next: nil}
		lastNode.Next = newListNode
		lastNode = newListNode
	}
	//BUG: renamed result up, didn't change it here
	return start
}

func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	res2 := l1
	for {
		if l2 == nil {
			break
		}
		if l1 == nil {
			l1 = l2
			break
		}
		l1.Val += l2.Val
		l2 = l2.Next
		if l1.Next == nil {
			l1.Next = l2
			break
		}
		l1 = l1.Next
	}
	l1 = res2
	for {
		if l1 == nil {
			break
		}
		if l1.Val >= 10 {
			if l1.Next == nil {
				l1.Next = &ListNode{}
			}
			l1.Next.Val += l1.Val / 10
			l1.Val = l1.Val % 10
		}
		l1 = l1.Next
	}
	return res2
}
