Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

    Input: haystack = "hello", needle = "ll"
    Output: 2

Example 2:
    
    Input: haystack = "aaaaa", needle = "bba"
    Output: -1

Clarification:

    What should we return when needle is an empty string? This is a great question to ask during an interview.
    
    For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

Bench:

    goos: linux
    goarch: amd64
    pkg: leetcode/28.
    BenchmarkStrStr-4       156344504                7.68 ns/op            0 B/op          0 allocs/op
    BenchmarkStrStr2-4       1000000              1047 ns/op             889 B/op         14 allocs/op
    BenchmarkStrStr3-4      100000000               10.2 ns/op             0 B/op          0 allocs/op
    PASS
    ok      leetcode/28.    4.073s

    
Submission:

    Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
    Memory Usage: 2.3 MB, less than 87.50% of Go online submissions for Implement strStr().