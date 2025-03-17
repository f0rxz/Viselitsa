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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	word := consts.Words[r.Intn(len(consts.Words))]

	cyrillicHalfLen := len(word) / 2
	discovered := make([]rune, cyrillicHalfLen)
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
