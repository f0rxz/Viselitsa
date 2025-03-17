package gamelogic

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"viselitsa/internal/consts"
	"viselitsa/internal/entity/game"
	"viselitsa/pkg/utils"
)

type GameLogic struct {
	game *game.Game
}

func NewGameLogic(g *game.Game) *GameLogic {
	return &GameLogic{game: g}
}

func (g *GameLogic) RevealLetter(letter rune) bool {
	found := false
	wordRunes := []rune(g.game.Word)

	for i, r := range wordRunes {
		if r == letter {
			g.game.Discovered[i] = letter
			found = true
		}
	}
	return found
}

func (g *GameLogic) ShowErrors() {
	fmt.Println(consts.ViselitsaStages[g.game.Errors])
}

func (g *GameLogic) ShowResult() {
	g.ShowErrors()
	if g.game.Errors >= len(consts.ViselitsaStages)-1 {
		fmt.Println("Вы проиграли! Загаданное слово было:", g.game.Word)
	} else {
		fmt.Println("Поздравляем! Вы угадали слово:", g.game.Word)
	}
}

func (g *GameLogic) Play() {
	scanner := bufio.NewScanner(os.Stdin)

	for g.canContinue() {
		g.displayGameState()
		letter := g.getInput(scanner)
		if letter == 0 {
			continue
		}

		if g.isLetterGuessed(letter) {
			fmt.Println("Вы уже вводили эту букву!")
			continue
		}

		g.processGuess(letter)
	}

	g.ShowResult()
}

func (g *GameLogic) canContinue() bool {
	return g.game.Errors < len(consts.ViselitsaStages)-1 && !utils.RuneSlicesEqual([]rune(g.game.Word), g.game.Discovered)
}

func (g *GameLogic) displayGameState() {
	g.ShowErrors()
	fmt.Println("Слово:", string(g.game.Discovered))
}

func (g *GameLogic) getInput(scanner *bufio.Scanner) rune {
	fmt.Print("Введите букву: ")
	scanner.Scan()
	input := scanner.Text()

	if len([]rune(input)) != 1 {
		fmt.Println("Введите только одну букву!")
		return 0
	}

	return []rune(input)[0]
}

func (g *GameLogic) isLetterGuessed(letter rune) bool {
	return slices.Contains(g.game.Guessed, letter)
}

func (g *GameLogic) processGuess(letter rune) {
	g.game.Guessed = append(g.game.Guessed, letter)
	if !g.RevealLetter(letter) {
		g.game.Errors++
	}
}
