package main

import (
	"testing"
)

func createSAC(size int, b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		NewSingleArrayCategorizedRandom(size)
	}
}

func playSAC(size int, b *testing.B) {
	g := NewSingleArrayCategorizedRandom(size)
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		g.PlayRound()
	}
}

func BenchmarkCreateSAC10(b *testing.B) { createSAC(10, b) }
func BenchmarkCreateSAC100(b *testing.B) { createSAC(100, b) }
func BenchmarkCreateSAC1000(b *testing.B) { createSAC(1000, b) }

func BenchmarkPlaySAC10(b *testing.B) { playSAC(10, b) }
func BenchmarkPlaySAC100(b *testing.B) { playSAC(100, b) }
func BenchmarkPlaySAC1000(b *testing.B) { playSAC(1000, b) }
