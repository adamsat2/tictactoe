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

	game.UtilTool.Dialogue("Goodbye", fmt.Println)
}
