package main

type GameOfLife interface {
	PlayRound()
	AsString(alive, dead string) string
}
