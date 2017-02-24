# BattleShip

A simple BattleShip simulation game played between two Players.

## Problem Statement

This try to solve problem discribed in [here](https://github.com/gojek-engineering/gophercon-2017)

## Design

1. Represent Input based on given file. Input format varies for grid and moves
2. Make all the moves (first Player1 then Player2 sequentially)
3. Calculate the score and update the grid
4. End the game once done with all the moves
5. Print the score and result of the game

## Usage

```go

package main

import (
	"fmt"
	"os"

	"github.com/kavirajk/gojek/battleship/game"
)

func main() {
	g := game.New(os.Stdin)
	g.Play()

	fmt.Println("P1 score:", g.P1Score())
	fmt.Println("P2 score:", g.P2Score())
}

```

## Run with inputs

```bash
$ go run main.go < input.txt

```


