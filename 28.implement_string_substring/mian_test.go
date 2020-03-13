package main

import "testing"

func BenchmarkStrStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStr("hello", "ll")
	}
}

func BenchmarkStrStr2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStr2("hello", "ll")
	}
}

func BenchmarkStrStr3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strStr3("hello", "ll")
	}
}
