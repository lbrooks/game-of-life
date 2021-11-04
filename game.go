package main

import (
	"math/rand"
	"strings"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())
var r = rand.New(src)

func randomIsAlive() int {
	v := r.Intn(2)
	if v == 0 {
		return 1
	}
	return 0
}

type Life struct {
	size int
	board []int
	buffer []int
}

func NewLife(size int) *Life {
	return &Life{
		size: size,
		board: make([]int, size*size),
		buffer: make([]int, size*size),
	}
}

func (l *Life) InitializeGame(board string) {
	useString := len(board) == len(l.board)

	for i := range l.board {
		if useString {
			c := board[i]
			if c == '1' {
				l.board[i] = 1
			} else {
				l.board[i] = 0
			}
		} else {
			v := r.Intn(25)
			if v == 0 {
				l.board[i] = 1
			} else {
				l.board[i] = 0
			}
		}

		l.buffer[i] = 0
	}
}

func (l *Life) AsString(alive, dead string) string {
	var out strings.Builder
	for i, cell := range l.board {
		if cell == 1 {
			out.WriteString(alive)
		} else {
			out.WriteString(dead)
		}
		if (i % l.size) == (l.size - 1) {
			out.WriteString("\n")
		}
	}
	return out.String()
}

func (l *Life) PlayRound() {
	for i, c := range l.board {
		count := l.sumNeighbors(i)
		l.buffer[i] = determineLife(c, count)
	}
	for i, c := range l.buffer {
		l.board[i] = c
	}
}

func (l *Life) sumNeighbors(idx int) int {
	sum := 0

	leftEdge := (idx % l.size) == 0
	rightEdge := (idx % l.size) == (l.size - 1)

	if idx >= l.size {
		//second row, can go up
		sum += l.valueFromBoard(idx-l.size)

		if !leftEdge {
			sum += l.valueFromBoard(idx-l.size-1)
		}

		if !rightEdge {
			sum += l.valueFromBoard(idx-l.size+1)
		}
	}
	if !leftEdge {
		sum += l.valueFromBoard(idx-1)
	}
	if !rightEdge {
		sum += l.valueFromBoard(idx+1)
	}
	if (idx/l.size) < (l.size - 1) {
		// not last row
		sum += l.valueFromBoard(idx+l.size)

		if !leftEdge {
			sum += l.valueFromBoard(idx+l.size-1)
		}

		if !rightEdge {
			sum += l.valueFromBoard(idx+l.size+1)
		}
	}
	return sum
}

func (l *Life) valueFromBoard(idx int) int {
	if idx < 0 {
		return 0
	}
	if idx >= (l.size * l.size) {
		return 0
	}
	return l.board[idx]
}

func determineLife(current, neighbors int) int {
	if current == 0 {
		if neighbors == 3 {
			return 1
		}
		return 0
	}

	if neighbors < 2 {
		return 0
	}
	if neighbors < 4 {
		return 1
	}
	return 0
}

