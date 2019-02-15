package main

// func WhiteStrategy(game Game, depth int, alpha int, beta int) (string, int) {
// 	moves := getMoves(game, WHITE)
// 	amtMoves := len(moves)
// 	if depth <= 0 || game.gameOver() {
// 		return "", heuristic(game, amtMoves)
// 	}
// 	if len(moves) == 0 {
// 		return BlackStrategy(game, depth - 1, alpha, beta)

// 	}

// 	bestSoFar := MinInt
// 	var bestMove string
// 	for _, move := range moves {
// 		myCopy := copyBoard(game)
// 		var gameCopy Game
// 		gameCopy.board = myCopy

// 		rawMove, _ := getIndex(move)
// 		gameCopy.flipAll(WHITE, rawMove)
// 		_, points := BlackStrategy(gameCopy, depth - 1, alpha, beta)

// 		if points >= bestSoFar {
// 			bestMove = move
// 			bestSoFar = points
// 		}
// 		alpha = max(alpha, bestSoFar)
// 		if beta <= alpha {
// 			break
// 		}
// 	}
// 	return bestMove, bestSoFar
// }

// func BlackStrategy(game Game, depth int, alpha int, beta int) (string, int) {
// 	moves := getMoves(game, BLACK)
// 	amtMoves := len(moves)
// 	if depth <= 0 || game.gameOver() {
// 		return "", heuristic(game, amtMoves)
// 	}
// 	if len(moves) == 0 {
// 		return WhiteStrategy(game, depth - 1, alpha, beta)
// 	}

// 	bestSoFar := MaxInt
// 	var bestMove string
// 	for _, move := range moves {
// 		myCopy := copyBoard(game)
// 		var gameCopy Game
// 		gameCopy.board = myCopy

// 		rawMove, _ := getIndex(move)
// 		gameCopy.flipAll(BLACK, rawMove)
// 		_, points := WhiteStrategy(gameCopy, depth - 1, alpha, beta)

// 		if points <= bestSoFar {
// 			bestMove = move
// 			bestSoFar = points
// 		}
// 		beta = min(beta, bestSoFar)
// 		if beta <= alpha {
// 			break
// 		}
// 	}
// 	return bestMove, bestSoFar
// }

func minimax(game Game, depth int, maximizing bool, alpha int, beta int) (int, string) {
	if game.gameOver() {
		return heuristic(game, 0), " "
	}
	if depth == 0 {
		mobi:= 0
		if maximizing {
			mobi = len(getMoves(game,WHITE))
		} else {
			mobi = len(getMoves(game, BLACK))
		}
		return heuristic(game, mobi), " "
	}
	if maximizing {
		bestSoFar := MinInt
		bestMove := " "
		moves := getMoves(game, WHITE)
		if len(moves) == 0 {
			return minimax(game, depth - 1, false, alpha, beta)
		}
		for _, move := range moves {
			gameCopy := copyGame(game)
			index, err := getIndex(move)
			if err != nil {
				print("ERRRRRRRRRRRRRRRR!!!")
			}
			gameCopy.flipAll(WHITE, index)
			val, _ := minimax(gameCopy, depth - 1, false, alpha, beta)
			if bestSoFar <= val {
				bestSoFar = val
				bestMove = move
			}
			alpha = max(bestSoFar, alpha)
			if beta <= alpha {
				break
			}
		}
		return bestSoFar, bestMove
	} else {
		bestSoFar := MaxInt
		bestMove := " "
		moves := getMoves(game, BLACK)
		for _, move := range moves {
			if len(moves) == 0 {
			return minimax(game, depth - 1, true, alpha, beta)
			}
			gameCopy := copyGame(game)
			index, err := getIndex(move)
			if err != nil {
				print("ERRRRRRRRRRRRRRRRRR!!! 2")
			}
			gameCopy.flipAll(BLACK, index)
			val, _ := minimax(gameCopy, depth - 1, true, alpha, beta)
			if bestSoFar >= val {
				bestMove = move
				bestSoFar = val
			}
			beta = min(beta, bestSoFar)
			if beta <= alpha {
				break
			}
		}
		return bestSoFar, bestMove
	}
	return -1, " "
}

func copyGame(game Game) Game {
	board := game.board
	var myCopy [64]Piece
	for i := 0; i < len(board); i++ {
		myCopy[i] = board[i]
	}
	var newGame Game
	newGame.board = myCopy
	return newGame
}

func heuristic(game Game, mobility int) int {
	board := game.board
	pieceScore := 0
	for i := 0; i < len(board); i++ {
		if board[i] == WHITE {
			pieceScore++
		} else if board[i] == BLACK {
			pieceScore--
		}
	}
	if game.gameOver() {
		if pieceScore > 0 {
			return MaxInt
		} else if pieceScore < 0 {
			return MinInt
		}
	}
	cornerCount := 0
	cornerPieces := []string{"A1","A8","H1","H8"}
	for _, cornerPiece := range cornerPieces {
		if piece, _ := game.get(cornerPiece); piece == BLACK {
			cornerCount--
		} else if piece, _ := game.get(cornerPiece); piece == WHITE {
			cornerCount++
		}
	}

	return pieceScore * 1 + mobility * 100 + cornerCount * 1000
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func safe(game Game, place string) bool {
	board := game.board
	pieceColor, _  := game.get(place)
	x := false
	y := false
	diagonalLeft := false
	diagonalRight := false 
	//x
	startingSpot, _ := getIndex(place)
	foundOpposite := false
	for i:= startingSpot; isAdjacent(i, i + 8); i += 8 {
		if !(board[i + 8] == pieceColor) {
			foundOpposite = true
			break
		}
	}
	if !foundOpposite {
		y = true
	}

	foundOpposite = false
	for i:= startingSpot; isAdjacent(i, i - 8); i -= 8 {
		if !(board[i - 8] == pieceColor) {
			foundOpposite = true
			break
		}
	}
	if !foundOpposite {
		y = true
	}

	foundOpposite = false
	for i:= startingSpot; isAdjacent(i, i - 1); i -= 1 {
		if !(board[i - 1] == pieceColor) {
			foundOpposite = true
			break
		}
	}
	if !foundOpposite {
		x = true
	}

	foundOpposite = false
	for i:= startingSpot; isAdjacent(i, i + 1); i += 1 {
		if !(board[i + 1] == pieceColor) {
			foundOpposite = true
			break
		}
	}
	if !foundOpposite {
		x = true
	}

	foundOpposite = false
	for i:= startingSpot; isAdjacent(i, i + 9); i += 9 {
		if !(board[i + 9] == pieceColor) {
			foundOpposite = true
			break
		}
	}
	if !foundOpposite {
		diagonalRight = true
	}

	foundOpposite = false
	for i:= startingSpot; isAdjacent(i, i - 9); i -= 9 {
		if !(board[i - 9] == pieceColor) {
			foundOpposite = true
			break
		}
	}
	if !foundOpposite {
		diagonalRight = true
	}

	foundOpposite = false
	for i:= startingSpot; isAdjacent(i, i - 7); i -= 7 {
		if !(board[i - 7] == pieceColor) {
			foundOpposite = true
			break
		}
	}
	if !foundOpposite {
		diagonalLeft = true
	}

	foundOpposite = false
	for i:= startingSpot; isAdjacent(i, i + 7); i += 7 {
		if !(board[i + 7] == pieceColor) {
			foundOpposite = true
			break
		}
	}
	if !foundOpposite {
		diagonalLeft = true
	}

	return x && y && diagonalLeft && diagonalRight

}