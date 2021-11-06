package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game Of Life", func() {
	DescribeTable(
		"determine life",
		func(alive, neighbors, expected int) {
			Ω(computeStatus(alive, neighbors)).To(Equal(expected))
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

