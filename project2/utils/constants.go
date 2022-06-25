package utils

import "strings"

const (
	NewGameError  = 1
	GameRuleError = 2

	ROWS    int = 6
	COLUMNS int = 7
)

var SectionDivider = strings.Repeat("*", 20)
var WinSeqForFirstPlayer = strings.Repeat("x", 4)
var WinSeqForSecondPlayer = strings.Repeat("o", 4)
