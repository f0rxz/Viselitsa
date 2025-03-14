package game

import (
	"math/rand"
	"time"
	"viselitsa/internal/consts"
)

type Game struct {
	Word       string
	Guessed    []rune
	Errors     int
	Discovered []rune
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	word := consts.Words[rand.Intn(len(consts.Words))]
	discovered := make([]rune, len(word)/2)
	for i := range discovered {
		discovered[i] = '_'
	}

	return &Game{
		Word:       word,
		Guessed:    []rune{},
		Errors:     0,
		Discovered: discovered,
	}
}
