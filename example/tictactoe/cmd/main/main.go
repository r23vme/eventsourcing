package main

import (
	"fmt"
	"math/rand"

	"github.com/r23vme/eventsourcing/aggregate"
	"github.com/r23vme/eventsourcing/eventstore/memory"
	"github.com/r23vme/eventsourcing/example/tictactoe"
)

func main() {
	es := memory.Create()
	aggregate.Register(&tictactoe.Game{})
	for i := 0; i < 10; i++ {
		game := PlayGame()
		fmt.Printf("game %d\n", i)
		game.Render()
		aggregate.Save(es, game)
	}
}

func PlayGame() *tictactoe.Game {
	game := tictactoe.NewGame()
	for !game.Done() {
		x := rand.Intn(3)
		y := rand.Intn(3)
		game.PlayMove(x, y)
	}
	return game
}
