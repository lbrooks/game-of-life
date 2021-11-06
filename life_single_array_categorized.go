package main

import (
	"strings"
)

type SingleArrayCategorized struct {
	size int
	board []int
	buffer []int
}

func NewSingleArrayCategorizedRandom(size int) *SingleArrayCategorized {
	return &SingleArrayCategorized{
		size: size,
		board: randomArray(size),
		buffer: make([]int, size*size),
	}
}

func NewSingleArrayCategorizedBoard(size int, board string) *SingleArrayCategorized {
	boardArray := parseArray(board)
	return &SingleArrayCategorized{
		size: strings.Index(board, ";") + 1,
		board: boardArray,
		buffer: make([]int, len(boardArray)),
	}
}

func (s *SingleArrayCategorized) AsString(alive, dead string) string {
	return arrayBoardToString(s.board, s.size, alive, dead)
}

func (s *SingleArrayCategorized) checkCorners() {
	ul := 0
	s.buffer[ul] = computeStatus(
		s.board[ul],
		s.board[ul + 1] + s.board[ul + s.size] + s.board[ul + s.size + 1],
	)

	ur := s.size - 1
	s.buffer[ur] = computeStatus(
		s.board[ur],
		s.board[ur - 1] + s.board[ur - 1 + s.size] + s.board[ur + s.size],
	)

	dl := len(s.board) - s.size
	s.buffer[dl] = computeStatus(
		s.board[dl],
		s.board[dl-s.size] + s.board[dl-s.size+1] + s.board[dl+1],
	)

	dr := len(s.buffer) - 1
	s.buffer[dr] = computeStatus(
		s.board[dr],
		s.board[dr-s.size-1] + s.board[dr-s.size] + s.board[dr-1],
	)
}

func (s *SingleArrayCategorized) checkEdges() {
	//top
	for i := 1; i < s.size - 1; i++ {
		s.buffer[i] = computeStatus(
			s.board[i],
			s.board[i-1] + s.board[i+1] + s.board[i-1+s.size] + s.board[i+s.size] + s.board[i+1+s.size],
		)
	}

	//left
	for i := s.size; i < len(s.buffer) - s.size; i += s.size {
		s.buffer[i] = computeStatus(
			s.board[i],
			s.board[i-s.size] + s.board[i-s.size+1] + s.board[i+1] + s.board[i+s.size] + s.board[i+s.size+1],
		)
	}

	//right
	for i := s.size + s.size - 1; i < len(s.buffer) - 1; i += s.size {
		s.buffer[i] = computeStatus(
			s.board[i],
			s.board[i-s.size-1] + s.board[i-s.size] + s.board[i-1] + s.board[i-1+s.size] + s.board[i+s.size],
		)
	}

	//bottom
	for i := len(s.buffer) - s.size + 1; i < len(s.buffer) - 1; i++ {
		s.buffer[i] = computeStatus(
			s.board[i],
			s.board[i-s.size-1] + s.board[i-s.size] + s.board[i-s.size+1] + s.board[i-1] + s.board[i+1],
		)
	}
}

func (s *SingleArrayCategorized) checkMiddle() {
	for i := s.size + 1; i < len(s.buffer) - 1 - s.size; {
		s.buffer[i] = computeStatus(
			s.board[i],
			s.board[i-s.size-1] + s.board[i-s.size] + s.board[i-s.size+1] + s.board[i-1] + s.board[i+1] + s.board[i+s.size-1] + s.board[i+s.size] + s.board[i+s.size+1],
		)
		if (i % s.size) == (s.size - 2) {
			i += 3
		} else {
			i++
		}
	}
}

func (s *SingleArrayCategorized) PlayRound() {
	s.checkCorners()
	s.checkEdges()
	s.checkMiddle()

	for i, c := range s.buffer {
		s.board[i] = c
	}
}

