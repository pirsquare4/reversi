package main
import (
	"errors"
	"strconv"
)
type Game struct {
	board [64]string
}

func (game Game) get(string place) int {
    if len(place) != 2 {
		err = errors.New(place + " is not a valid spot on board")
		return
	}
	letter := place[0]
	number := strconv.Atoi(place[1]) - 1
	var letternum int
	switch letter {
	case "A":
		letternum = 0
	case "B":
		letternum = 1
	case "C":
		letternum = 2
	case "D":
		letternum = 3
	case "E":
		letternum = 4
	case "F":
		letternum = 5
	case "G":
		letternum = 6
	case "H":
		letternum = 7
	}
}