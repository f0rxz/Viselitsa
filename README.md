# Viselitsa Game
This version i tried make it OOP style.
This is a command-line application written in Go that implements the classic "Viselitsa" game. The player is tasked with guessing a word by entering letters one by one. Each incorrect guess increases the error counter, and the corresponding viselitsa stage is displayed in ASCII graphics.

## Features
- Choose to start a new game or exit
- Random word selection from a predefined list
- Error tracking and ASCII viselitsa visualization
- Checking guessed letters
- Displaying the result (win or lose)

## Project Structure
```
├── cmd
│   └── main.go               # Entry point of the application
├── go.mod                    # Go modules
├── internal
│   ├── consts
│   │   └── consts.go         # Project constants
│   ├── entity
│   │   └── game
│   │       └── game.go       # Definition of the Game structure
│   └── usecase
│       └── gamelogic
│           └── gamelogic.go  # Game logic (input handling, word checking)
└── pkg
    └── utils
        └── utils.go          # Utility functions
```

## Installation and Running
1. Clone the repository:
   ```sh
   git clone https://github.com/f0rxz/Viselitsa.git
   cd viselitsa-game
   ```
2. Build and run the game:
   ```sh
   go run cmd/main.go
   ```

## How to Play
1. At startup, choose an action:
   - `1` – Start a new game
   - `2` – Exit
2. Enter a letter to try and guess the word.
3. If the letter is in the word, it will be revealed.
4. If the letter is incorrect, the error counter will increase, and the next stage of the viselitsa will be displayed.
5. The game continues until the word is guessed or the maximum number of errors is reached.

## TODO
- Add support for user-defined words
- Improve the interface (colorful terminal output)
- Add an option to guess the entire word at once

## License
GNU GENERAL PUBLIC LICENSE
Version 3, 29 June 2007

