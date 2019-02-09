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
		if board[adjacentIndex] == player.Opposite() && isAdjacent(adjacentIndex, adjacentIndex + adder) {
				valid := checkSandwhich(GameState, player, adjacentIndex, adder, found)
				return valid
		} else if board[adjacentIndex] == player {
			return true
		} else {
			return false
		}
	}
}

func destructiveSandwhich(GameState *Game, player Piece, index int, adder int, found bool) bool{
	board := GameState.board
	adjacentIndex := adder + index
	if !found {
		if board[adjacentIndex] == player.Opposite() && isAdjacent(adjacentIndex, adjacentIndex + adder) {
			valid := destructiveSandwhich(GameState, player, adjacentIndex, adder, true)
			if valid {
				GameState.setRaw(adjacentIndex, player)
			}
			return valid
		} else {
			return false
		}
	} else {
		if board[adjacentIndex] == player.Opposite() && isAdjacent(adjacentIndex, adjacentIndex + adder) {
				valid := destructiveSandwhich(GameState, player, adjacentIndex, adder, found)
				if valid {
				GameState.setRaw(adjacentIndex, player)
				}
				return valid
		} else if board[adjacentIndex] == player {
			return true
		} else {
			return false
		}
	}
}

func (GameState *Game) flipAll (player Piece, index int) {
	adjacentPlaces := [...]int{-9, -8, -7, -1, 1, 7, 8, 9}
	for _, adjacentPlace := range adjacentPlaces {
		destructiveSandwhich(GameState, player, index, adjacentPlace, false)
	}
	GameState.setRaw(index, player)
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
	row := strconv.Itoa(place / BOARDSIZE + 1)
	remainder := (place % BOARDSIZE)
	var column string
	switch remainder {
	case 0:
		column = "A"
	case 1:
		column = "B"
	case 2:
		column = "C"
	case 3:
		column = "D"
	case 4:
		column = "E"
	case 5:
		column = "F"
	case 6:
		column = "G"
	case 7:
		column = "H"
	}
	columnAndRow := column + row
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

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}