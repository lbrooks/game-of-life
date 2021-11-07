package main

import (
	"testing"
)

func createDAB(size int, b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		NewDoubleArrayBruteRandom(size)
	}
}

func playDAB(size int, b *testing.B) {
	g := NewDoubleArrayBruteRandom(size)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		g.PlayRound()
	}
}

func BenchmarkCreateDAB10(b *testing.B) { createDAB(10, b) }
func BenchmarkCreateDAB100(b *testing.B) { createDAB(100, b) }
func BenchmarkCreateDAB1000(b *testing.B) { createDAB(1000, b) }

func BenchmarkPlayDAB10(b *testing.B) { playDAB(10, b) }
func BenchmarkPlayDAB100(b *testing.B) { playDAB(100, b) }
func BenchmarkPlayDAB1000(b *testing.B) { playDAB(1000, b) }
