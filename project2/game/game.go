package game

import (
	"../utils"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Position struct {
	column int
	row    int
}

func (position *Position) GetPosition() (column int, row int) {
	return position.column, position.row
}

type Game struct {
	occupiedCells     map[int][utils.ROWS]string
	IsFirstPlayerStep bool
	lastInsertedPos   Position
	lastError         string
}

func Construct() Game {
	var occupiedCells = map[int][utils.ROWS]string{}

	for i := 0; i < utils.COLUMNS; i++ {
		row := [utils.ROWS]string{}
		for j := 0; j < utils.ROWS; j++ {
			row[j] = " "
		}

		occupiedCells[i] = row
	}

	return Game{
		occupiedCells:     occupiedCells,
		IsFirstPlayerStep: true,
		lastInsertedPos:   Position{},
		lastError:         "",
	}
}

func (game *Game) Step() error {
	column, err := game.getColumnInput()
	if err != nil {
		return err
	}

	// it's easier for calculations
	column -= 1

	if !(0 <= column && column < utils.COLUMNS) {
		game.lastError = "Columns must be in 0 <= column < COLUMNS."
		return &utils.GameException{Type: utils.GameRuleError}

	}

	currRow := game.occupiedCells[column]
	if !strings.Contains(strings.Join(currRow[:], ""), " ") {
		game.lastError = "Column is fulled."
		return &utils.GameException{Type: utils.GameRuleError}
	}

	occupiedColumn := game.occupiedCells[column]
	for row := len(game.occupiedCells[column]) - 1; row >= 0; row-- {

		if occupiedColumn[row] == " " {
			if game.IsFirstPlayerStep {
				occupiedColumn[row] = "x"
			} else {
				occupiedColumn[row] = "o"
			}

			game.lastInsertedPos = Position{column, row}
			break
		}
	}
	game.occupiedCells[column] = occupiedColumn

	return nil
}

func (game *Game) Render() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("\bConnect 4 v1.0")
	fmt.Println(utils.SectionDivider)

	for row := 0; row < utils.ROWS; row++ {
		for column := 0; column < utils.COLUMNS; column++ {
			fmt.Print("|", game.occupiedCells[column][row])
		}
		fmt.Println("|")
	}

	fmt.Print(utils.SectionDivider)

	currPlayer := 1
	if !game.IsFirstPlayerStep {
		currPlayer = 2
	}

	fmt.Printf("\nStep of %v player.\n", currPlayer)

	if game.lastError != "" {
		fmt.Println(game.lastError)
		game.lastError = ""
	}
}

func (game *Game) IsEndGame() bool {
	var winSequence = utils.WinSeqForFirstPlayer
	if !game.IsFirstPlayerStep {
		winSequence = utils.WinSeqForSecondPlayer
	}

	// check win sequence
	for _, i := range game.occupiedCells {
		joinedCells := strings.Join(i[:], "")
		if strings.Contains(joinedCells, winSequence) {
			return true
		}
	}

	// transpose and check win sequence
	var tr = [utils.ROWS][utils.COLUMNS]string{}

	for i := 0; i < utils.COLUMNS; i++ {
		for j := 0; j < utils.ROWS; j++ {
			tr[j][i] = game.occupiedCells[i][j]
		}
	}

	for _, r := range tr {
		joinedCells := strings.Join(r[:], "")
		if strings.Contains(joinedCells, winSequence) {
			return true
		}
	}

	// check diagonals
	column, row := game.lastInsertedPos.GetPosition()
	currPlayerSym := "x"
	if !game.IsFirstPlayerStep {
		currPlayerSym = "o"
	}

	// current symbol is in suit
	symCounts := 1
	currColumn := column - 1
	currRow := row - 1

	// move left and up
	for currColumn >= 0 && currRow >= 0 {
		// if symbol does not belong player
		if currPlayerSym != game.occupiedCells[currColumn][currRow] {
			break
		}

		symCounts += 1
		currColumn -= 1
		currRow -= 1
	}

	currColumn = column + 1
	currRow = row + 1

	// move right and down
	for currColumn < utils.COLUMNS && currRow < utils.ROWS {
		if currPlayerSym != game.occupiedCells[currColumn][currRow] {
			break
		}

		symCounts += 1
		currColumn += 1
		currRow += 1
	}

	if symCounts >= 4 {
		return true
	}

	// reset it - because it must count symbols by another diagonal
	symCounts = 1
	currColumn = column + 1
	currRow = row - 1

	//move right and up
	for currColumn < utils.COLUMNS && currRow >= 0 {
		if currPlayerSym != game.occupiedCells[currColumn][currRow] {
			break
		}

		symCounts += 1
		currColumn += 1
		currRow -= 1
	}

	currColumn = column - 1
	currRow = row + 1

	// move left and down
	for currColumn >= 0 && currRow < utils.ROWS {
		if currPlayerSym != game.occupiedCells[currColumn][currRow] {
			break
		}

		symCounts += 1
		currColumn -= 1
		currRow += 1
	}

	if symCounts >= 4 {
		return true
	}

	return false
}

func (game *Game) getColumnInput() (int, error) {
	fmt.Printf("Enter column(1-%v),'n' - new game, 'q' - for exit: ", utils.COLUMNS)

	var userInput string
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("User input error:", err)
		os.Exit(1)
	}

	if userInput == "n" {
		return 0, &utils.GameException{Type: utils.NewGameError}
	} else if userInput == "q" {
		os.Exit(0)
	}

	column, err := strconv.Atoi(userInput)
	if err != nil {
		game.lastError = fmt.Sprintf("Must be integer number  between 1 and %v", utils.COLUMNS)
		return 0, &utils.GameException{Type: utils.GameRuleError}
	}

	return column, nil
}

func (game *Game) ChangePlayer() {
	game.IsFirstPlayerStep = !game.IsFirstPlayerStep
}
