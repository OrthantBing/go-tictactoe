package main

import "testing"

func TestIsGameWon(t *testing.T) {
	// initialize the board

	tests := []struct {
		Name string
		game func() *Game
		mock func()
		want func(isGameWon bool)
	}{
		{
			Name: "Empty",
			game: func() *Game {
				p1, _ := newPlayer("Player1", "x")
				p2, _ := newPlayer("Player 2", "o")
				b := makeBoard(3)
				g := newGame(b, p1, p2)
				return g
			},
			mock: func() {},
			want: func(isGameWon bool) {
				if isGameWon != false {
					t.Error("Game should not have been won")
				}
			},
		},
		{
			Name: "First Row",
			game: func() *Game {
				p1, _ := newPlayer("Player1", "x")
				p2, _ := newPlayer("Player 2", "o")
				b := makeBoard(3)
				b[0][0] = "x"
				b[0][1] = "x"
				b[0][2] = "x"
				g := newGame(b, p1, p2)
				return g
			},
			mock: func() {},
			want: func(isGameWon bool) {
				if isGameWon != true {
					t.Error("Game should have been won")
				}
			},
		},
		{
			Name: "First Column",
			game: func() *Game {
				p1, _ := newPlayer("Player1", "x")
				p2, _ := newPlayer("Player 2", "o")
				b := makeBoard(3)
				b[0][0] = "x"
				b[1][0] = "x"
				b[2][0] = "x"
				g := newGame(b, p1, p2)
				return g
			},
			mock: func() {},
			want: func(isGameWon bool) {
				if !isGameWon {
					t.Error("Game should have been won")
				}
			},
		},
		{
			Name: "Diagonal",
			game: func() *Game {
				p1, _ := newPlayer("Player1", "x")
				p2, _ := newPlayer("Player 2", "o")
				b := makeBoard(3)
				b[0][0] = "x"
				b[1][1] = "x"
				b[2][2] = "x"
				g := newGame(b, p1, p2)
				return g
			},
			mock: func() {},
			want: func(isGameWon bool) {
				if !isGameWon {
					t.Error("Game should have been won")
				}
			},
		},
		{
			Name: "Cross Diagonal",
			game: func() *Game {
				p1, _ := newPlayer("Player1", "x")
				p2, _ := newPlayer("Player 2", "o")
				b := makeBoard(3)
				b[0][2] = "x"
				b[1][1] = "x"
				b[2][0] = "x"
				g := newGame(b, p1, p2)
				return g
			},
			mock: func() {},
			want: func(isGameWon bool) {
				if !isGameWon {
					t.Error("Game should have been won")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			tt.mock()
			g := tt.game()
			tt.want(g.IsGameWon(g.CurrentPlayer))
		})
	}
}
