package main

import (
	"fmt"
	"tictactoe/game"
)

func main() {
	playAgain := true

	for playAgain {
		game.Play()
		playAgain = game.UtilTool.GetYesOrNo("Would you like to play again (y/n)?")
	}

	fmt.Println(" Goodbye")
}
