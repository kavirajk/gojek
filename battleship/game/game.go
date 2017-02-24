package game

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

type Game struct {
	Grid1, Grid2 Grid
	Move1, Move2 Moves
	M, S, T      int

	// private fields
	p1score, p2score         int
	p1ShipCount, p2ShipCount int
}

type Move struct {
	X, Y int
}

type Moves []Move

func (m Moves) populate(line string) {
	csvreader := csv.NewReader(strings.NewReader(line))
	csvreader.Comma = ':'
	records, _ := csvreader.ReadAll()
	for i := range records[0] {
		var x, y int
		fmt.Sscanf(records[0][i], "%d,%d", &x, &y)
		m[i] = Move{x, y}
	}
}

type Grid [][]byte

func (g Grid) String() string {
	// TODO(kaviraj): Fix extra space and newline edge case
	var s string
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[0]); j++ {
			s += fmt.Sprintf("%c ", g[i][j])
		}
		s += fmt.Sprintf("\n")
	}
	return s
}

func (g Grid) populate(line string) {
	csvreader := csv.NewReader(strings.NewReader(line))
	records, _ := csvreader.ReadAll()
	for i := range records[0] {
		var x, y int
		fmt.Sscanf(records[0][i], "%d:%d", &x, &y)
		g[x][y] = 'B'
	}
}

func New(r io.Reader) *Game {
	var g Game
	fmt.Fscanln(r, &g.M)
	fmt.Fscanln(r, &g.S)

	// Get Player1 positions
	g.Grid1 = make([][]byte, g.M)
	for i := 0; i < g.M; i++ {
		g.Grid1[i] = make([]byte, g.M)
		for j := 0; j < g.M; j++ {
			g.Grid1[i][j] = '_'
		}
	}

	var line string
	fmt.Fscanln(r, &line)

	g.Grid1.populate(line)

	// Get Player2 positions
	g.Grid2 = make([][]byte, g.M)
	for i := 0; i < g.M; i++ {
		g.Grid2[i] = make([]byte, g.M)
		for j := 0; j < g.M; j++ {
			g.Grid2[i][j] = '_'
		}
	}

	fmt.Fscanln(r, &line)

	g.Grid2.populate(line)

	fmt.Fscanln(r, &g.T)

	fmt.Fscanln(r, &line)
	g.Move1 = make([]Move, g.T)
	g.Move1.populate(line)

	fmt.Fscanln(r, &line)
	g.Move2 = make([]Move, g.T)
	g.Move2.populate(line)

	g.p1ShipCount = g.S
	g.p2ShipCount = g.S

	return &g
}

func (g *Game) Play() {
	for i := 0; i < len(g.Move1); i++ {
		x, y := g.Move1[i].X, g.Move1[i].Y
		if g.p1ShipCount != 0 && g.Grid2[x][y] == 'B' {
			g.Grid2[x][y] = 'X'
			g.p1score++
			g.p2ShipCount--
		} else {
			g.Grid2[x][y] = 'O'
		}

		x, y = g.Move2[i].X, g.Move2[i].Y
		if g.p2ShipCount != 0 && g.Grid1[x][y] == 'B' {
			g.Grid1[x][y] = 'X'
			g.p2score++
			g.p1ShipCount--
		} else {
			g.Grid1[x][y] = 'O'
		}

	}
}

func (g *Game) P1Score() int {
	return g.p1score
}

func (g *Game) P2Score() int {
	return g.p2score
}
