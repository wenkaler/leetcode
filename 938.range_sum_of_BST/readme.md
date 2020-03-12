Given the root node of a binary search tree, return the sum of values of all nodes with value between L and R (inclusive).

The binary search tree is guaranteed to have unique values.

Task: https://leetcode.com/problems/range-sum-of-bst/

Example 1:

    Input: root = [10,5,15,3,7,null,18], L = 7, R = 15
    Output: 32

Example 2:

    Input: root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
    Output: 23

Bench:

    goos: linux
    goarch: amd64
    pkg: leetcode/938_range_sum_of_BST
    BenchmarkTreeNode_RangeSum-4    20527737                51.2 ns/op             0 B/op          0 allocs/op
    BenchmarkCreateTreeNode-4         895252              1120 ns/op             960 B/op         30 allocs/op
    PASS
    ok      leetcode/938_range_sum_of_BST   2.154s

Tests:

    ok      leetcode/938_range_sum_of_BST   1.009s  coverage: 85.7% of statements

Submission:

    Runtime: 88 ms, faster than 72.41% of Go online submissions for Range Sum of BST.
    Memory Usage: 7.8 MB, less than 100.00% of Go online submissions for Range Sum of BST.