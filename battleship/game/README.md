# BattleShip

A simple BattleShip simulation game played between two Players.

## Problem Statement

This try to solve problem discribed in [here](https://github.com/gojek-engineering/gophercon-2017)

## Design

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


