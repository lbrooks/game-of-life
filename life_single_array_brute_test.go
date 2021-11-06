package main

import (
	"testing"
)

func createSAB(size int, b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		NewSingleArrayBruteRandom(size)
	}
}

func playSAB(size int, b *testing.B) {
	g := NewSingleArrayBruteRandom(size)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		g.PlayRound()
	}
}

func BenchmarkCreateSAB10(b *testing.B) { createSAB(10, b) }
func BenchmarkCreateSAB100(b *testing.B) { createSAB(100, b) }
func BenchmarkCreateSAB1000(b *testing.B) { createSAB(1000, b) }

func BenchmarkPlaySAB10(b *testing.B) { playSAB(10, b) }
func BenchmarkPlaySAB100(b *testing.B) { playSAB(100, b) }
func BenchmarkPlaySAB1000(b *testing.B) { playSAB(1000, b) }
