package main

import (
	"bufio"
	"fmt"
	"os"
	"viselitsa/internal/entity/game"
	"viselitsa/internal/usecase/gamelogic"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("1) Начать новую игру")
		fmt.Println("2) Выйти")
		fmt.Print("Введите номер действия: ")
		scanner.Scan()
		choice := scanner.Text()

		if choice == "1" {
			game := gamelogic.NewGameLogic(game.NewGame())
			game.Play()
		} else if choice == "2" {
			fmt.Println("Выход...")
			break
		} else {
			fmt.Println("Неверный ввод. Попробуйте снова.")
		}
	}
}
