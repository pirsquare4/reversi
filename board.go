package main
import (
	"strconv"
)

func checkSandwhich(GameState Game, player Piece, index int, adder int, found bool) bool{
	board := GameState.board
	adjacentIndex := adder + index
	if !found {
		if board[adjacentIndex] == player.Opposite() && isAdjacent(adjacentIndex, adjacentIndex + adder) {
			valid := checkSandwhich(GameState, player, adjacentIndex, adder, true)
			return valid
		} else {
			return false
		}
	} else {
		if board[adjacentIndex] == player.Opposite() {
			if isAdjacent(adjacentIndex, adjacentIndex + adder) {
				valid := checkSandwhich(GameState, player, adjacentIndex, adder, true)
				return valid
			}
			return false
		} else {
			return true
		}
	}
}

func isAdjacent(tile1 int, tile2 int) bool {
	if tile1 > 63 || tile2 > 63 {
		return false
	} else if  tile1 < 0 || tile2 < 0 {
		return false 
	} else if (tile1 % BOARDSIZE == tile2 % BOARDSIZE - 1) ||
	 tile1 % BOARDSIZE == tile2 % BOARDSIZE || 
	 (tile1 % BOARDSIZE == tile2 % BOARDSIZE + 1) {
	 	if isAbove(tile1, tile2) || isSame(tile1, tile2) || isBelow(tile1, tile2) {
	 		return true
	 	}
	 }

	return false
}

func TranslateToMove(place int) string {
	if place < 0 || place > 63 {
		return "NaN"
	}
	quotient := place / BOARDSIZE
	column := strconv.Itoa((place % BOARDSIZE) + 1)
	var row string
	switch quotient {
	case 0:
		row = "A"
	case 1:
		row = "B"
	case 2:
		row = "C"
	case 3:
		row = "D"
	case 4:
		row = "E"
	case 5:
		row = "F"
	case 6:
		row = "G"
	case 7:
		row = "H"
	}
	columnAndRow := row + column
	return columnAndRow

}

func isAbove(tile1 int, tile2 int) bool {
	return (tile1/8) - (tile2/8) == -1
}
func isSame(tile1 int, tile2 int) bool {
	return (tile1/8) - (tile2/8) == 0
}

func isBelow(tile1 int, tile2 int) bool {
	return (tile1/8) - (tile2/8) == 1
}