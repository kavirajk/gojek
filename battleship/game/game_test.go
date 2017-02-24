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
	}{{
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
	}}

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
