package gamelogic

import (
	"bufio"
	"fmt"
	"os"
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

func (g *GameLogic) ShowResult() {
	fmt.Println(consts.ViselitsaStages[g.game.Errors])
	if g.game.Errors >= len(consts.ViselitsaStages)-1 {
		fmt.Println("Вы проиграли! Загаданное слово было:", g.game.Word)
	} else {
		fmt.Println("Поздравляем! Вы угадали слово:", g.game.Word)
	}
}

func (g *GameLogic) Play() {
	scanner := bufio.NewScanner(os.Stdin)

	wordRunes := []rune(g.game.Word) //забыл что я оперирую символами как байт а не рунами...

	for g.game.Errors < len(consts.ViselitsaStages)-1 && string(wordRunes) != string(g.game.Discovered) {
		fmt.Println(consts.ViselitsaStages[g.game.Errors])
		fmt.Println("Слово:", string(g.game.Discovered))
		fmt.Print("Введите букву: ")

		scanner.Scan()
		input := scanner.Text()
		if len([]rune(input)) != 1 {
			fmt.Println("Введите только одну букву!")
			continue
		}

		letter := []rune(input)[0]

		if utils.Contains(g.game.Guessed, letter) {
			fmt.Println("Вы уже вводили эту букву!")
			continue
		}

		g.game.Guessed = append(g.game.Guessed, letter)
		if !g.RevealLetter(letter) {
			g.game.Errors++
		}
	}

	g.ShowResult()
}
