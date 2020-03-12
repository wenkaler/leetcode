Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

    Input: dividend = 10, divisor = 3
    Output: 3

Example 2:

    Input: dividend = 7, divisor = -3
    Output: -2

Note:

    Both dividend and divisor will be 32-bit signed integers.
    The divisor will never be 0.
    Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.

Bench:

    goos: linux
    goarch: amd64
    pkg: leetcode/29.divide_two_integers
    BenchmarkDivide-4       33333637                36.0 ns/op             0 B/op          0 allocs/op
    BenchmarkStandard-4     1000000000               0.278 ns/op           0 B/op          0 allocs/op
    PASS
    ok      leetcode/29.divide_two_integers 2.529s
    Time: 0h:00m:03s

Submission:

    Runtime: 2500 ms, faster than 24.78% of Go online submissions for Divide Two Integers.
    Memory Usage: 2.5 MB, less than 100.00% of Go online submissions for Divide Two Integers.    