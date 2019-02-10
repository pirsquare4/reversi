package main

func WhiteStrategy(game Game, depth int) string {
	if depth <= 0 {
		return game.score
	}
	moves := getMoves(game, WHITE)
	//put if gameover and if len(moves) < 0 case
	bestSoFar := -1000
	var bestMove string
	var bestBoard [64]Piece
	for _, move := range moves {
		myCopy := copyBoard(game)
		rawMove := getIndex(move)
		myCopy.flipAll(currentplayer, rawMove)
		heuristic = heuristic(myCopy)
		if heuristic > bestSoFar {
			bestMove = move
			bestSoFar = heuristic
		}
	}
	return bestMove
}

func BlackStrategy(game Game) string {
	return "hi"
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
	a = 0
	for i := 0; i < len(board); i++ {
		if board[i] == WHITE {
			a++
		} else if board[i] == BLACK {
			a--
		}
	}
	return a
}