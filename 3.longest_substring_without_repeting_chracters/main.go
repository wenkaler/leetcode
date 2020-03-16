package main

import "strings"

func main() {
	println(lengthOfLongestSubstring("dvdf"))
	println(lengthOfLongestSubstring2("dvdf"))
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	count := 1
	maxcount := 1
	i := 0
	for j := 1; i < len(s) && j < len(s); {
		if strings.Contains(s[i:j], s[j:j+1]) {
			if count > maxcount {
				maxcount = count
			}
			count = 1
			i += 1
			j = i + 1
		} else {
			count += 1
			j += 1
		}
	}
	if count > maxcount {
		maxcount = count
	}
	return maxcount
}

func lengthOfLongestSubstring(s string) int {
	var c, max, d int
	m := make(map[uint8]int, len(s))
	for i := range s {
		if idx, ok := m[s[i]]; ok {
			if max < c {
				max = c
			}
			if idx >= d {
				c = i - (idx + 1)
				d = idx
			}
		}
		m[s[i]] = i
		c++
	}
	if max < c {
		return c
	}
	return max
}
