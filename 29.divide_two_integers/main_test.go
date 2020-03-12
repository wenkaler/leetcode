package main

import "testing"

func TestDivide(t *testing.T) {
	testCases := []struct {
		name     string
		dividend int
		divisor  int
		want     int
	}{
		{
			name:     "divided 10 by 2",
			dividend: 10,
			divisor:  2,
			want:     5,
		},
		{
			name:     "divided 10 by 3",
			dividend: 10,
			divisor:  3,
			want:     3,
		},
		{
			name:     "divided 10 by -3",
			dividend: 10,
			divisor:  -3,
			want:     -3,
		},
		{
			name:     "divided -7 by 2",
			dividend: -7,
			divisor:  2,
			want:     -3,
		},
		{
			name:     "divided -1 by 1",
			dividend: -1,
			divisor:  1,
			want:     -1,
		},
		{
			name:     "divided -2147483648 by -1",
			dividend: -2147483648,
			divisor:  -1,
			want:     2147483647,
		},
		{
			name:     "divided -2147483648 by 1",
			dividend: -2147483648,
			divisor:  1,
			want:     -2147483648,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := divide(tc.dividend, tc.divisor)
			if got != tc.want {
				t.Fatalf("failed divide got(%d) want(%d)", got, tc.want)
			}
		})
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		divide(10, 2)
	}
}

func BenchmarkStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = 10 / 2
	}
}
