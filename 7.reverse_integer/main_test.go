package main

import "testing"

func TestReverse(t *testing.T) {
	testCases := []struct {
		name  string
		value int
		exp   int
	}{
		{
			name:  "normal response",
			value: 123,
			exp:   321,
		}, {
			name:  "negative normal response",
			value: -123,
			exp:   -321,
		}, {
			name:  "zero",
			value: 0,
			exp:   0,
		}, {
			name:  "overflow",
			value: 1534236469,
			exp:   0,
		}, {
			name:  "negative overflow",
			value: -1534236469,
			exp:   0,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := reverse(tc.value)
			if tc.exp != got {
				t.Fatalf("failed reverse got(%d) exp(%d)", got, tc.exp)
			}
		})
	}
}
