package main

import (
)

type SingleArrayBrute struct {
	size int
	board []int
	buffer []int
}

func NewSingleArrayBruteRandom(size int) *SingleArrayBrute {
	return &SingleArrayBrute{
		size: size,
		board: randomArray(size),
		buffer: make([]int, size*size),
	}
}

func NewSingleArrayBruteBoard(size int, board string) *SingleArrayBrute {
	return &SingleArrayBrute{
		size: size,
		board: parseArray(board),
		buffer: make([]int, size*size),
	}
}

func (s *SingleArrayBrute) AsString(alive, dead string) string {
	return arrayBoardToString(s.board, s.size, alive, dead)
}

func (s *SingleArrayBrute) PlayRound() {
	for i, c := range s.board {
		count := s.sumNeighbors(i)
		s.buffer[i] = computeStatus(c, count)
	}
	for i, c := range s.buffer {
		s.board[i] = c
	}
}

func (s *SingleArrayBrute) sumNeighbors(idx int) int {
	sum := 0

	leftEdge := (idx % s.size) == 0
	rightEdge := (idx % s.size) == (s.size - 1)

	if idx >= s.size {
		//second row, can go up
		sum += s.valueFromBoard(idx-s.size)

		if !leftEdge {
			sum += s.valueFromBoard(idx-s.size-1)
		}

		if !rightEdge {
			sum += s.valueFromBoard(idx-s.size+1)
		}
	}
	if !leftEdge {
		sum += s.valueFromBoard(idx-1)
	}
	if !rightEdge {
		sum += s.valueFromBoard(idx+1)
	}
	if (idx/s.size) < (s.size - 1) {
		// not last row
		sum += s.valueFromBoard(idx+s.size)

		if !leftEdge {
			sum += s.valueFromBoard(idx+s.size-1)
		}

		if !rightEdge {
			sum += s.valueFromBoard(idx+s.size+1)
		}
	}
	return sum
}

func (s *SingleArrayBrute) valueFromBoard(idx int) int {
	if idx < 0 {
		return 0
	}
	if idx >= (s.size * s.size) {
		return 0
	}
	return s.board[idx]
}

