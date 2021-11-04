package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGameOfLife(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GameOfLife Suite")
}
