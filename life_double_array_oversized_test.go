package main

import (
	"testing"
)

func createDAO(size int, b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		NewDoubleArrayOversizedRandom(size)
	}
}

func playDAO(size int, b *testing.B) {
	g := NewDoubleArrayOversizedRandom(size)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		g.PlayRound()
	}
}

func BenchmarkCreateDAO10(b *testing.B) { createDAO(10, b) }
func BenchmarkCreateDAO100(b *testing.B) { createDAO(100, b) }
func BenchmarkCreateDAO1000(b *testing.B) { createDAO(1000, b) }

func BenchmarkPlayDAO10(b *testing.B) { playDAO(10, b) }
func BenchmarkPlayDAO100(b *testing.B) { playDAO(100, b) }
func BenchmarkPlayDAO1000(b *testing.B) { playDAO(1000, b) }
