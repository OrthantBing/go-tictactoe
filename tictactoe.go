package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	X = "x"
	O = "o"
)

type Board [][]string

func (b Board) IsBoardFull() bool {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == "_" {
				return false
			}
		}
	}

	return true
}

func (b Board) Initialize() {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			b[i][j] = "_"
		}
	}
}

func (b Board) IsSlotEmpty(i, j int) bool {
	return b[i][j] == "_"
}

func (b Board) IsValidSlot(i, j int) bool {
	if i >= len(b) {
		return false
	}
	if j >= len(b) {
		return false
	}

	return true
}

func (b Board) String() string {
	s := strings.Builder{}
	for _, val := range b {
		s.WriteString(strings.Join(val, " "))
		s.WriteString("\n")
	}

	return s.String()
}

func makeBoard(size int) Board {
	b := make([][]string, size)
	for i := 0; i < len(b); i++ {
		b[i] = make([]string, size)
	}
	return b
}

type Player struct {
	Name              string
	AssignedCharacter string
}

func newPlayer(name string, ac string) (*Player, error) {

	switch ac {
	case "x":
		break
	case "o":
		break
	default:
		return nil, fmt.Errorf("Invalid Character")
	}

	return &Player{
		Name:              name,
		AssignedCharacter: ac,
	}, nil
}

type Game struct {
	CurrentPlayer *Player
	NextPlayer    *Player
	Board         Board
}

func newGame(b Board, cp *Player, np *Player) *Game {
	return &Game{
		CurrentPlayer: cp,
		NextPlayer:    np,
		Board:         b,
	}
}

func (g *Game) SwitchPlayer() error {
	g.CurrentPlayer, g.NextPlayer = g.NextPlayer, g.CurrentPlayer
	return nil
}

func (g *Game) IsGameWon(p *Player) bool {
	characterToCheck := p.AssignedCharacter

	for i := 0; i < len(g.Board); i++ {
		canWin := true
		for j := 0; j < len(g.Board[i]); j++ {
			if g.Board[i][j] != characterToCheck {
				canWin = false
				break
			}
		}
		if canWin {
			return true
		}
	}

	for j := 0; j < len(g.Board); j++ {
		canWin := true
		for i := 0; i < len(g.Board[j]); i++ {
			if g.Board[i][j] != characterToCheck {
				canWin = false
				break
			}
		}
		if canWin {
			return true
		}
	}

	// Diagonal
	canWin := true
	for i, j := 0, 0; i < len(g.Board) && j < len(g.Board); i, j = i+1, j+1 {
		if g.Board[i][j] != characterToCheck {
			canWin = false
		}
	}
	if canWin {
		return true
	}

	//Alternate diagnoal
	canWin = true

	for i, j := 0, len(g.Board)-1; i < len(g.Board) && j >= 0; i, j = i+1, j-1 {
		if g.Board[i][j] != characterToCheck {
			canWin = false
		}
	}

	if canWin {
		return true
	}

	return false
}

func (g *Game) SetValue(i, j int, val string) error {
	if !g.Board.IsValidSlot(i, j) {
		return fmt.Errorf("Invalid Slot")
	}

	if !g.Board.IsSlotEmpty(i, j) {
		return fmt.Errorf("Slot not empty")
	}

	g.Board[i][j] = val
	return nil
}

func (g *Game) Play() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Println(g.Board)
	fmt.Printf("Player: %s \n", g.CurrentPlayer.Name)
	fmt.Println("row column")

	for s.Scan() {

		arr := strings.Split(s.Text(), " ")
		i, err := strconv.Atoi(arr[0])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		j, err := strconv.Atoi(arr[1])
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if err := g.SetValue(i, j, g.CurrentPlayer.AssignedCharacter); err != nil {
			fmt.Println(err.Error())
			continue
		}

		if g.IsGameWon(g.CurrentPlayer) {

			fmt.Printf("%s won the game \n", g.CurrentPlayer.Name)
			fmt.Println(g.Board)
			break
		}

		if g.Board.IsBoardFull() {
			fmt.Println("Game Draw")
			break
		}

		g.SwitchPlayer()
		fmt.Println(g.Board)
		fmt.Printf("Player: %s \n", g.CurrentPlayer.Name)
		fmt.Println("row column value")
	}
}

func main() {
	size := 3
	b := makeBoard(size)
	b.Initialize()
	p1, _ := newPlayer("Anton", "x")
	p2, _ := newPlayer("Winston", "o")
	g := newGame(b, p1, p2)
	g.Play()
}
