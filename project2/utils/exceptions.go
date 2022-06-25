package utils

type GameException struct{ Type int }

func (except *GameException) Error() string {
	return ""
}
