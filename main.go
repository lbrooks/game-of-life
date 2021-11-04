package main

import (
	"math/rand"
	"strings"
	"time"

	"github.com/pterm/pterm"
)

var size int
var board, buffer []int

func formatNumber(i int) string {
	if i == 1 {
		return "✅"
	}
	return "❌"
}

func printBoard() string {
	var out strings.Builder
	for i, cell := range board {
		out.WriteString(formatNumber(cell))
		if i > 0 && (i % size) == (size-1) {
			out.WriteString("\n")
		}
	}
	return out.String()
}

func toIndex(row, col int) int {
	return (row * size) + col
}

func fromIndex(idx int) (row, col int) {
	return idx/size, idx%size
}

func valueFromBoard(idx int) int {
	if idx < 0 {
		return 0
	}
	if idx >= (size * size) {
		return 0
	}
	return board[idx]
}

func sumNeighbors(idx int) int {
	sum := 0

	leftEdge := (idx % size) == 0
	rightEdge := (idx % size) == (size - 1)

	if idx >= size {
		//second row, can go up
		sum += valueFromBoard(idx-size)

		if !leftEdge {
			sum += valueFromBoard(idx-size-1)
		}

		if !rightEdge {
			sum += valueFromBoard(idx-size+1)
		}
	}
	if !leftEdge {
		sum += valueFromBoard(idx-1)
	}
	if !rightEdge {
		sum += valueFromBoard(idx+1)
	}
	if (idx/size) < (size - 1) {
		// not last row
		sum += valueFromBoard(idx+size)

		if !leftEdge {
			sum += valueFromBoard(idx+size-1)
		}

		if !rightEdge {
			sum += valueFromBoard(idx+size+1)
		}
	}
	return sum
}

func dumpBuffer() {
	for i, cell := range buffer {
		board[i] = cell
	}
}

func determineLife(current, neighbors int) int {
	if neighbors < 2 {
		return 0
	}
	if neighbors < 4 {
		return 1
	}
	if current == 1 {
		return 0
	}
	return 1
}

func playRound() {
	for i, c := range board {
		count := sumNeighbors(i)
		buffer[i] = determineLife(c, count)
	}
	dumpBuffer();
}

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

func main() {
	size = 25
//	rounds := 1000

	board = make([]int, size*size)
	buffer = make([]int, size*size)

	for i := 0; i < len(board); i++ {
		rv := r.Intn(25)
		if rv == 0 {
			board[i] = 1
		}
	}

	area, _ := pterm.DefaultArea.WithCenter().Start() // Start the Area printer, with the Center option.

	str := pterm.DefaultParagraph.WithMaxWidth(19).Sprint(printBoard())
	area.Update(str)

	//for i := 0; i < rounds; i++ {
	for true {
		time.Sleep(500 * time.Millisecond)
		playRound()
		area.Update(pterm.DefaultParagraph.WithMaxWidth(19).Sprint(printBoard()))
	}
}

