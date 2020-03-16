package main

import "testing"

func TestSumLongSubstring(t *testing.T) {
	testCases := []struct {
		name   string
		s      string
		result int
	}{
		{
			name:   "empty string",
			s:      " ",
			result: 1,
		},
		{
			name:   "normal response",
			s:      "dvdf",
			result: 3,
		},
		{
			name:   "normal response",
			s:      "abba",
			result: 2,
		},
		{
			name:   "normal response",
			s:      "eeydgwdykpv",
			result: 7,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tc.s)
			if got != tc.result {
				t.Fatalf("failed count string got(%d) exp(%d", got, tc.result)
			}
		})
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("pwskwjanx")
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("pwskwjanx")
	}
}
