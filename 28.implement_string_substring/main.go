package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(strStr("hello", "ll"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for ; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

func strStr2(haystack string, needle string) int {
	if needle != "" {
		var re = regexp.MustCompile(needle)
		idx := re.FindStringIndex(haystack)
		if len(idx) > 0 {
			return idx[0]
		}
		return -1
	}
	return 0
}

func strStr3(haystack string, needle string) int {
	if needle != "" {
		b := []byte(haystack)
	next:
		for i := range b {
			idx := i
			for _, v := range []byte(needle) {
				if i >= len(b) || b[i] != v {
					continue next
				}
				i++
			}
			return idx
		}
		return -1
	}
	return 0
}
