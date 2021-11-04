package main

import (
	"flag"
	"time"

	"github.com/pterm/pterm"
)

var flagSize = flag.Int("size", 10, "size of the board")
var flagBoard = flag.String("board", "", "starting board")
var flagAlive = flag.String("alive", "✅", "character representing 'alive'")
var flagDead = flag.String("dead", "❌", "character representing 'dead'")

var alive, dead string

func main() {
	flag.Parse()

	alive = *flagAlive
	dead = *flagDead

	game := NewLife(*flagSize)
	game.InitializeGame(*flagBoard)

	area, _ := pterm.DefaultArea.WithCenter().Start()
	para := pterm.DefaultParagraph.WithMaxWidth(*flagSize)

	area.Update(para.Sprint(game.AsString(alive, dead)))

	for true {
		time.Sleep(500 * time.Millisecond)
		game.PlayRound()
		area.Update(para.Sprint(game.AsString(alive, dead)))
	}
}

