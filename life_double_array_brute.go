package main

import (
)

type DoubleArrayBrute struct {
	board [][]int
	buffer [][]int
}

func NewDoubleArrayBruteRandom(size int) *DoubleArrayBrute {
	return &DoubleArrayBrute{
		board: randomArrayArrays(size),
		buffer: blankArrayArrays(size),
	}
}

func NewDoubleArrayBruteBoard(size int, board string) *DoubleArrayBrute {
	b := parseArrayArrays(board)
	return &DoubleArrayBrute{
		board: b,
		buffer: blankArrayArrays(len(b)),
	}
}

func (s *DoubleArrayBrute) AsString(alive, dead string) string {
	return arrayArraysBoardToString(s.board, alive, dead)
}

func (s *DoubleArrayBrute) PlayRound() {
	for i, row := range s.board {
		for j, cell := range row {
			count := s.sumNeighbors(i, j)
			s.buffer[i][j] = computeStatus(cell, count)
		}
	}
	for i, row := range s.buffer {
		for j, cell := range row {
			s.board[i][j] = cell
		}
	}
}

func (s *DoubleArrayBrute) sumNeighbors(i, j int) int {
	sum := 0

	leftEdge := j == 0
	rightEdge := j == (len(s.board) - 1)

	if i >= 1 {
		//second row, can go up
		sum += s.board[i-1][j]

		if !leftEdge {
			sum += s.board[i-1][j-1]
		}

		if !rightEdge {
			sum += s.board[i-1][j+1]
		}
	}
	if !leftEdge {
		sum += s.board[i][j-1]
	}
	if !rightEdge {
		sum += s.board[i][j+1]
	}
	if i < (len(s.board) - 1) {
		// not last row
		sum += s.board[i+1][j]

		if !leftEdge {
			sum += s.board[i+1][j-1]
		}

		if !rightEdge {
			sum += s.board[i+1][j+1]
		}
	}
	return sum
}

