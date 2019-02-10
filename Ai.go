package main

func WhiteStrategy(game Game, depth int, alpha int, beta int) (string, int) {
	if depth <= 0 || game.gameOver() {
		return "", heuristic(game)
	}
	moves := getMoves(game, WHITE)
	if len(moves) == 0 {
		return "", MinInt
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
		return "", MaxInt
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
	a := 0
	for i := 0; i < len(board); i++ {
		if board[i] == WHITE {
			a++
		} else if board[i] == BLACK {
			a--
		}
	}
	if game.gameOver() {
		if a > 0 {
			a = MaxInt
		} else if a < 0 {
			a = MinInt
		}
	}
	return a
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