You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

    Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
    Output: 7 -> 0 -> 8
    Explanation: 342 + 465 = 807.

Bench: 
    
    goos: linux
    goarch: amd64
    pkg: github.com/wenkaler/leetcode/2.
    BenchmarkAddTwoNumbers-4         3537435               331 ns/op             192 B/op         12 allocs/op
    BenchmarkAddTwoNumbers2-4        4012284               294 ns/op             176 B/op         11 allocs/op
    BenchmarkAddTwoNumbers3-4       38243866                32.0 ns/op             0 B/op          0 allocs/op
    PASS
    ok      github.com/wenkaler/leetcode/2. 4.265s
    Time: 0h:00m:04s                