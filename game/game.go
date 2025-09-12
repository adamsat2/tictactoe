package game

import (
	"fmt"

	"github.com/adamsat2/tgutils"
	"github.com/fatih/color"
)

var UtilTool = tgutils.Utils{
	SpaceBeforeText: true,
	ColoredText:     true,
}

const EMPTYCELL = " "
const SYMBOLONE = "X"
const SYMBOLTWO = "O"

func setGameSlice(gameSlice []string) {
	for i := range gameSlice {
		gameSlice[i] = EMPTYCELL
	}
}

func updateSymbolInBoard(gameSlice []string, symbol string) {
	displayGameSlice(gameSlice)
	question := fmt.Sprintf("%s's turn, select where to place your symbol (1-9)", symbol)
	goodAnswer := false

	for !goodAnswer {
		pos := UtilTool.GetNumber(question) - 1
		if pos < 0 || pos > 8 {
			color.Red(" The position you've selected is out of bounds!")
		} else if gameSlice[pos] != EMPTYCELL {
			color.Red(" The position you've selected already has a symbol in it!")
		} else {
			gameSlice[pos] = symbol
			goodAnswer = true
			color.Cyan(" %s placed at position %d", symbol, pos+1)
		}
	}

}

func checkWinner(gameSlice []string) bool {
	var symbol string
	for i := 0; i <= 6; i += 3 {
		if (gameSlice[i] != EMPTYCELL) && (gameSlice[i] == gameSlice[i+1]) && (gameSlice[i] == gameSlice[i+2]) {
			symbol = gameSlice[i]
			color.Green(" %s won by row combination!", symbol)
			return true
		}
	}
	for i := 0; i < 3; i++ {
		if (gameSlice[i] != EMPTYCELL) && (gameSlice[i] == gameSlice[i+3]) && (gameSlice[i] == gameSlice[i+6]) {
			symbol = gameSlice[i]
			color.Green(" %s won by column combination!", symbol)
			return true
		}
	}
	if (gameSlice[0] != EMPTYCELL) && (gameSlice[0] == gameSlice[4]) && (gameSlice[0] == gameSlice[8]) {
		symbol = gameSlice[0]
		color.Green(" %s won by diagonal combination!", symbol)
		return true
	}
	if (gameSlice[2] != EMPTYCELL) && (gameSlice[2] == gameSlice[4]) && (gameSlice[2] == gameSlice[6]) {
		symbol = gameSlice[2]
		color.Green(" %s won by diagonal combination!", symbol)
		return true
	}
	return false
}

func checkDraw(gameSlice []string) bool {
	for i := range gameSlice {
		if gameSlice[i] == EMPTYCELL {
			return false
		}
	}
	color.Yellow(" The board is full and there is no winner. It's a draw!")
	return true
}

func Play() {
	gameSlice := make([]string, 9)
	setGameSlice(gameSlice)
	turn := 1

	UtilTool.ClearScreen() // for new games

	for !checkDraw(gameSlice) && !checkWinner(gameSlice) {
		color.Magenta(" Welcome to tictactoe!")
		if turn%2 == 0 {
			updateSymbolInBoard(gameSlice, SYMBOLTWO)
		} else {
			updateSymbolInBoard(gameSlice, SYMBOLONE)
		}
		UtilTool.ClearScreen()
		turn++
		if checkDraw(gameSlice) || checkWinner(gameSlice) {
			displayGameSlice(gameSlice)
			break
		}
	}
}

func displayGameSlice(gameSlice []string) {
	for i := 0; i < len(gameSlice); i++ {
		print(" ")
		fmt.Print("[", gameSlice[i], "]")
		if (i+1)%3 == 0 {
			fmt.Println()
		}
	}
}
