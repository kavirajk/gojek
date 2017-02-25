package game

import (
	"io"
	"strings"
	"testing"
)

// Test whether input is properly populated to Game struct
func TestNewGame(t *testing.T) {
	var tests = []struct {
		input                           io.Reader
		expectedM, expectedS, expectedT int
	}{
		{
			input: strings.NewReader(`5
5
1:1,2:0,2:3,3:4,4:3
0:1,2:3,3:0,3:4,4:1
5
0,1:4,3:2,3:3,1:4,1
0,1:0,0:1,1:2,3:4,3
`),
			expectedM: 5,
			expectedS: 5,
			expectedT: 5,
		},
		{
			input: strings.NewReader(`5
3
1:1,2:0,2:3
0:1,2:3,3:0
4
0,1:4,3:2,3:3,1
0,1:0,0:1,1:2,3
`),
			expectedM: 5,
			expectedS: 3,
			expectedT: 4,
		},
		{
			// Edge case
			input: strings.NewReader(`1
1
0:0
0:0
1
0,0
0,0
`),
			expectedM: 1,
			expectedS: 1,
			expectedT: 1,
		},
	}

	for _, test := range tests {
		output := New(test.input)
		if output.M != test.expectedM {
			t.Errorf("failed. M scaned incorrect. exptected %d, got %d", test.expectedM, output.M)
		}
		if output.S != test.expectedS {
			t.Errorf("failed. S scaned incorrect. exptected %d, got %d", test.expectedS, output.S)
		}
		if output.T != test.expectedT {
			t.Errorf("failed. T scaned incorrect. exptected %d, got %d", test.expectedT, output.T)
		}

	}
}

// Test the play() logic, by testing p1 and p2 score for different inputs
func TestGameResult(t *testing.T) {
	var tests = []struct {
		input            io.Reader
		p1Score, p2Score int
		result           string
	}{
		{
			input: strings.NewReader(`5
5
1:1,2:0,2:3,3:4,4:3
0:1,2:3,3:0,3:4,4:1
5
0,1:4,3:2,3:3,1:4,1
0,1:0,0:1,2:2,3:4,3
`),
			p1Score: 3,
			p2Score: 2,
			result:  "Player 1 wins",
		},
		{
			input: strings.NewReader(`5
5
1:1,2:0,2:3,3:4,4:3
0:1,2:3,3:0,3:4,4:1
5
0,1:4,3:2,3:3,1:4,1
0,1:0,0:1,1:2,3:4,3
`),
			p1Score: 3,
			p2Score: 3,
			result:  "It is a draw",
		},
		{
			// Edge case
			input: strings.NewReader(`1
1
0:0
0:0
1
0,0
0,0
`),
			p1Score: 1,
			p2Score: 0, // P1 already made a move and all the p2 ships are dead
			result:  "Player 1 wins",
		},
	}

	for _, test := range tests {
		output := New(test.input)
		output.Play() // Simulate Play
		if output.P1Score() != test.p1Score {
			t.Errorf("failed. P1Score incorrect. exptected %d, got %d", test.p1Score, output.P1Score())
		}
		if output.P2Score() != test.p2Score {
			t.Errorf("failed. P2Score incorrect. exptected %d, got %d", test.p2Score, output.P2Score())
		}
		if output.Result() != test.result {
			t.Errorf("failed. Result incorrect. exptected %s, got %s", test.result, output.Result())
		}

	}

}
