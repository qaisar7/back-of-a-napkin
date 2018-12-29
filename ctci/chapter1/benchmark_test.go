package main

import (
	"testing"
)

var (
	s1 string
)

func BenchmarkPalindromPermutate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palinPermut(s1)
	}
}

func BenchmarkPalindromPermutateOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palinPermutOptimized(s1)
	}
}

func init() {
	s1 = "Tact Csoa"
}
