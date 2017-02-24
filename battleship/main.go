package main

import (
	"fmt"
	"os"

	"github.com/kavirajk/gojek/battleship/game"
)

func main() {
	g := game.New(os.Stdin)
	g.Play()

	fmt.Println("Player1")
	fmt.Println(g.Grid1)
	fmt.Println("Player2")
	fmt.Println(g.Grid2)
	p1Score := g.P1Score()
	p2Score := g.P2Score()
	fmt.Println("P1:", p1Score)
	fmt.Println("P2:", p2Score)

	if p1Score > p2Score {
		fmt.Println("P1 Won")
	} else if p2Score > p2Score {
		fmt.Println("P2 Won")
	} else {
		fmt.Println("It is a draw")
	}
}
