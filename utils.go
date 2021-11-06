package main

import (
	"math/rand"
	"strings"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomAlive() int {
	if r.Intn(2) == 0 {
		return 1;
	}
	return 0;
}

func randomArray(size int) []int {
	arr := make([]int, size*size)
	for i := range arr {
		arr[i] = randomAlive()
	}
	return arr
}

func randomArrayArrays(size int) [][]int {
	arr := make([][]int, size)
	for i := range arr {
		arr[i] = make([]int, size)
		for j := range arr[i] {
			arr[i][j] = randomAlive()
		}
	}
	return arr
}

func parseArray(board string) []int {
	singleBoard := strings.ReplaceAll(board, ";", "")
	arr := make([]int, len(singleBoard))
	for i, c := range board {
		if c == '1' {
			arr[i] = 1
		}
	}
	return arr
}

func parseArrayArrays(board string) [][]int {
	rows := strings.Split(board, ";")
	arr := make([][]int, len(rows))
	for i, row := range rows {
		arr[i] = make([]int, len(row))
		for j, c := range row {
			if c == '1' {
				arr[i][j] = 1
			}
		}
	}
	return arr
}

func arrayBoardToString(board []int, size int, alive, dead string) string {
	var out strings.Builder
	for i, cell := range board {
		if cell == 1 {
			out.WriteString(alive)
		} else {
			out.WriteString(dead)
		}
		if (i % size) == (size - 1) {
			out.WriteString("\n")
		}
	}
	return out.String()
}

func arrayArraysBoardToString(board [][]int, alive, dead string) string {
	var out strings.Builder
	for _, cols := range board {
		for _, cell := range cols {
			if cell == 1 {
				out.WriteString(alive)
			} else {
				out.WriteString(dead)
			}
		}
		out.WriteString("\n")
	}
	return out.String()
}

func computeStatus(cell, neighbors int) int {
	if cell == 0 {
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

