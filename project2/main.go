package main

import (
	"./game"
	"./utils"
	"fmt"
	"os"
)

func main() {
	g := game.Construct()

	for {
		g.Render()

		err := g.Step()
		if err != nil {

			if gameException, ok := err.(*utils.GameException); ok {
				if gameException.Type == utils.NewGameError {
					g = game.Construct()
				}

				continue
			} else {
				fmt.Println("Unknown error: ", err)
				os.Exit(1)
			}
		}

		g.Render()

		if g.IsEndGame() {
			v := 1
			if !g.IsFirstPlayerStep {
				v = 2
			}

			fmt.Printf("Win %v player!\n New game(y/n)?", v)
			var isNewGame string
			if _, err := fmt.Scan(&isNewGame); err != nil {
				fmt.Println("User input error: ", err)
				os.Exit(1)
			}

			if isNewGame == "y" {
				g = game.Construct()
				continue
			}

			break
		}

		g.ChangePlayer()
	}
}
