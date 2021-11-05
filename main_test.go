package main

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game Of Life", func() {
	DescribeTable(
		"determine life",
		func(alive, neighbors, expected int) {
			Ω(determineLife(alive, neighbors)).To(Equal(expected))
		},
		Entry("Alive - 0", 1, 0, 0),
		Entry("Alive - 1", 1, 1, 0),
		Entry("Alive - 2", 1, 2, 1),
		Entry("Alive - 3", 1, 3, 1),
		Entry("Alive - 4", 1, 4, 0),
		Entry("Alive - 5", 1, 5, 0),
		Entry("Alive - 6", 1, 6, 0),
		Entry("Alive - 7", 1, 7, 0),
		Entry("Alive - 8", 1, 8, 0),
		Entry("Dead - 0", 0, 0, 0),
		Entry("Dead - 1", 0, 1, 0),
		Entry("Dead - 2", 0, 2, 0),
		Entry("Dead - 3", 0, 3, 1),
		Entry("Dead - 4", 0, 4, 0),
		Entry("Dead - 5", 0, 5, 0),
		Entry("Dead - 6", 0, 6, 0),
		Entry("Dead - 7", 0, 7, 0),
		Entry("Dead - 8", 0, 8, 0),
	)
})

func createGame(size int, b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		NewLife(size)
	}
}

func BenchmarkCreate10(b *testing.B) { createGame(10, b) }
func BenchmarkCreate100(b *testing.B) { createGame(100, b) }
func BenchmarkCreate1000(b *testing.B) { createGame(1000, b) }

func randomGame(size int, b *testing.B) {
	g := NewLife(size)
	g.InitializeGame("")
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		g.PlayRound()
	}
}

func BenchmarkPlayOneLoop10(b *testing.B) { randomGame(10, b) }
func BenchmarkPlayOneLoop100(b *testing.B) { randomGame(100, b) }
func BenchmarkPlayOneLoop1000(b *testing.B) { randomGame(1000, b) }


func random2Game(size int, b *testing.B) {
	g := NewLife(size)
	g.InitializeGame("")
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		g.PlayRoundCategorized()
	}
}

func BenchmarkPlayCategorized10(b *testing.B) { random2Game(10, b) }
func BenchmarkPlayCategorized100(b *testing.B) { random2Game(100, b) }
func BenchmarkPlayCategorized1000(b *testing.B) { random2Game(1000, b) }

