package main

import (
	"testing"
)

func createDAC(size int, b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		NewDoubleArrayCategorizedRandom(size)
	}
}

func playDAC(size int, b *testing.B) {
	g := NewDoubleArrayCategorizedRandom(size)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		g.PlayRound()
	}
}

func BenchmarkCreateDAC10(b *testing.B) { createDAC(10, b) }
func BenchmarkCreateDAC100(b *testing.B) { createDAC(100, b) }
func BenchmarkCreateDAC1000(b *testing.B) { createDAC(1000, b) }

func BenchmarkPlayDAC10(b *testing.B) { playDAC(10, b) }
func BenchmarkPlayDAC100(b *testing.B) { playDAC(100, b) }
func BenchmarkPlayDAC1000(b *testing.B) { playDAC(1000, b) }
