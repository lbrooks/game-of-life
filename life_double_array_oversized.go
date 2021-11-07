package main

import (
)

type DoubleArrayOversized struct {
	board [][]int
	buffer [][]int
}

func NewDoubleArrayOversizedRandom(size int) *DoubleArrayOversized {
	return &DoubleArrayOversized{
		board: randomArrayArraysOversized(size),
		buffer: blankArrayArrays(size+2),
	}
}

func NewDoubleArrayOversizedBoard(size int, board string) *DoubleArrayOversized {
	b := parseArrayArraysOversized(board)
	return &DoubleArrayOversized{
		board: b,
		buffer: blankArrayArrays(len(b)),
	}
}

func (s *DoubleArrayOversized) AsString(alive, dead string) string {
	return arrayArraysOversizedBoardToString(s.board, alive, dead)
}

func (s *DoubleArrayOversized) PlayRound() {
	for i := 1; i < len(s.board) - 1; i++ {
		for j := 1; j < len(s.board) - 1; j++ {
			s.buffer[i][j] = computeStatus(
				s.board[i][j],
				s.board[i-1][j-1] + s.board[i-1][j] + s.board[i-1][j+1] + s.board[i][j-1] + s.board[i][j+1] + s.board[i+1][j-1] + s.board[i+1][j] + s.board[i+1][j+1],
			)
		}
	}
	for i, row := range s.buffer {
		for j, cell := range row {
			s.board[i][j] = cell
		}
	}
}

