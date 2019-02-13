package main

func WhiteStrategy(game Game, depth int, alpha int, beta int) (string, int) {
	if depth <= 0 || game.gameOver() {
		return "", heuristic(game)
	}
	moves := getMoves(game, WHITE)
	if len(moves) == 0 {
		return "", -1000
	}

	bestSoFar := MinInt
	var bestMove string
	for _, move := range moves {
		myCopy := copyBoard(game)
		var gameCopy Game
		gameCopy.board = myCopy

		rawMove, _ := getIndex(move)
		gameCopy.flipAll(WHITE, rawMove)
		_, points := BlackStrategy(gameCopy, depth - 1, alpha, beta)

		if points >= bestSoFar {
			bestMove = move
			bestSoFar = points
		}
		alpha = max(alpha, bestSoFar)
		if beta <= alpha {
			break
		}
	}
	return bestMove, heuristic(game)
}

func BlackStrategy(game Game, depth int, alpha int, beta int) (string, int) {
	if depth <= 0 || game.gameOver() {
		return "", heuristic(game)
	}
	moves := getMoves(game, BLACK)
	if len(moves) == 0 {
		return "", 1000
	}

	bestSoFar := MaxInt
	var bestMove string
	for _, move := range moves {
		myCopy := copyBoard(game)
		var gameCopy Game
		gameCopy.board = myCopy

		rawMove, _ := getIndex(move)
		gameCopy.flipAll(BLACK, rawMove)
		_, points := WhiteStrategy(gameCopy, depth - 1, alpha, beta)

		if points <= bestSoFar {
			bestMove = move
			bestSoFar = points
		}
		beta = min(beta, bestSoFar)
		if beta <= alpha {
			break
		}
	}
	return bestMove, heuristic(game)
}

func copyBoard(game Game) [64]Piece {
	board := game.board
	var myCopy [64]Piece
	for i := 0; i < len(board); i++ {
		myCopy[i] = board[i]
	}
	return myCopy
}

func heuristic(game Game) int {
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

	edgeCount := 0
	edgePieces := []string{"A3","A4","A5","A6","C1","D1","E1","F1","C8","D8","E8","F8","H3","H4","H5","H6"}
	for _, edgePiece := range edgePieces {
		if piece, _ := game.get(edgePiece); piece == WHITE {
			edgeCount++
		} else if piece, _ := game.get(edgePiece); piece == BLACK {
			edgeCount--
		}
	}
	nextToCount := 0
	nextToCornerPieces := []string{"A2","B1","B2","A7","B8","B7","G1","H2","G2","G8","H7","G7"}
	for _, nextToCornerPiece := range nextToCornerPieces {
		if piece, _ := game.get(nextToCornerPiece); piece == WHITE {
			nextToCount++
		} else if piece, _ := game.get(nextToCornerPiece); piece == BLACK {
			nextToCount--
		}
	}

	safeCount := 0
	for i := 0; i < len(board); i++ {
		if board[i] == WHITE && safe(game, TranslateToMove(i)) {
			safeCount++
		} else if board[i] == BLACK && safe(game, TranslateToMove(i)) {
			pieceScore--
		}
	} 
	return pieceScore + cornerCount * 10 + pieceScore * 6 + edgeCount * 4 - (nextToCount * 8) 
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