package main

import (
)

type DoubleArrayCategorized struct {
	board [][]int
	buffer [][]int
}

func NewDoubleArrayCategorizedRandom(size int) *DoubleArrayCategorized {
	return &DoubleArrayCategorized{
		board: randomArrayArrays(size),
		buffer: blankArrayArrays(size),
	}
}

func NewDoubleArrayCategorizedBoard(board string) *DoubleArrayCategorized {
	b := parseArrayArrays(board)
	return &DoubleArrayCategorized{
		board: b,
		buffer: blankArrayArrays(len(b)),
	}
}

func (s *DoubleArrayCategorized) AsString(alive, dead string) string {
	return arrayArraysBoardToString(s.board, alive, dead)
}

func (s *DoubleArrayCategorized) checkCorners() {
	e := len(s.board) - 1

	s.buffer[0][0] = computeStatus(
		s.board[0][0],
		s.board[0][1] + s.board[1][0] + s.board[1][1],
	)

	s.buffer[0][e] = computeStatus(
		s.board[0][e],
		s.board[0][e-1] + s.board[1][e-1] + s.board[1][e],
	)

	s.buffer[e][0] = computeStatus(
		s.board[e][0],
		s.board[e-1][0] + s.board[e-1][1] + s.board[e][1],
	)

	s.buffer[e][e] = computeStatus(
		s.board[e][e],
		s.board[e-1][e-1] + s.board[e-1][e] + s.board[e][e-1],
	)
}

func (s *DoubleArrayCategorized) checkEdges() {
	e := len(s.board) - 1
	for i := 1; i < len(s.board) - 1; i++ {
		// top
		s.buffer[0][i] = computeStatus(
			s.board[0][i],
			s.board[0][i-1] + s.board[0][i+1] + s.board[1][i-1] + s.board[1][i] + s.board[1][i+1],
		)

		// bottom
		s.buffer[e][i] = computeStatus(
			s.board[e][i],
			s.board[e][i-1] + s.board[e][i+1] + s.board[e-1][i-1] + s.board[e-1][i] + s.board[e-1][i+1],
		)

		// left
		s.buffer[i][0] = computeStatus(
			s.board[i][0],
			s.board[i-1][0] + s.board[i-1][1] + s.board[i][1] + s.board[i+1][0] + s.board[i+1][1],
		)

		// right
		s.buffer[i][e] = computeStatus(
			s.board[i][e],
			s.board[i-1][e] + s.board[i-1][e-1] + s.board[i][e-1] + s.board[i+1][e] + s.board[i+1][e-1],
		)
	}
}

func (s *DoubleArrayCategorized) checkMiddle() {
	for i := 1; i < len(s.buffer) - 1; i++ {
		for j := 1; j < len(s.buffer) - 1; j++ {
			s.buffer[i][j] = computeStatus(
				s.board[i][j],
				s.board[i-1][j-1] + s.board[i-1][j] + s.board[i-1][j+1] + s.board[i][j-1] + s.board[i][j+1] + s.board[i+1][j-1] + s.board[i+1][j] + s.board[i+1][j+1],
			)
		}
	}
}

func (s *DoubleArrayCategorized) PlayRound() {
	s.checkCorners()
	s.checkEdges()
	s.checkMiddle()

	for i, row := range s.buffer {
		for j, c := range row {
			s.board[i][j] = c
		}
	}
}

