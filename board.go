package main
import (
)
func checkSandwhich(GameState Game, player Piece, index int, adder int, found bool) bool{
	board := GameState.board
	adjacentIndex := adder + index
	if !found {
		if board[adjacentIndex] == player.Opposite() {

		}
	} else {

	}
	return true
}