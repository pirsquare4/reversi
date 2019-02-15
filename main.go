package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
	"strconv"
	"strings"
	"math/rand"
	"time"
	"log"
)

var BOARDSIZE = 8
var DEPTH = 6

const MaxInt = 1000000
const MinInt = -1000000

//Where the games begin!
func main() {
	c := make(chan string, 1)
	loop := true
	ActivateAI := false
	ActivateAI2 := false
	a := true
	for loop {
		reader := bufio.NewReader(os.Stdin)
		if a {
			fmt.Println("Hello, and welcome to reversi, a strategy game where you must outplay and outskill your opponent!")
			fmt.Println("First of all, how many human players will be participating? 1 or 2?")
		}
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		if text == "0" {
			ActivateAI = true
			ActivateAI2 = true
			loop = false
		} else if (text == "1") {
			ActivateAI = true
			loop = false
		} else if text == "2" {
			ActivateAI = false
			loop = false
		} else {
			fmt.Println("Invalid number of players, please try again!")
			a = false
		}
	}
	rand.Seed(time.Now().UnixNano())
	loop = true
	fmt.Println("Enter seed")
	seedReader := bufio.NewReader(os.Stdin)
	seedtext, _ := seedReader.ReadString('\n')
	seedtext = strings.Replace(seedtext, "\n", "", -1)
	seedtext = strings.Replace(seedtext, "\r", "", -1)
	seed, _ := strconv.ParseInt(string(seedtext), 0, 64)
	rand.Seed(seed)

	fmt.Println("Please enter, in seconds, how much time each player has to make a move (e.g. 30, or 60)")
	timerReader := bufio.NewReader(os.Stdin)
	timertext, _ := timerReader.ReadString('\n')
	timertext = strings.Replace(timertext, "\n", "", -1)
	timertext = strings.Replace(timertext, "\r", "", -1)
	timer, _ := strconv.ParseInt(string(timertext), 0, 64)

	for loop {
		reader := bufio.NewReader(os.Stdin)
		if ActivateAI && ActivateAI2 {
			fmt.Println("Choose a Color for Computer 1: White or Black?")		
		} else {
			fmt.Println("Player 1, Choose your Color: White or Black?")
		}
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		if (text == "black" || text == "Black") {
			Player1 = BLACK
			Player2 = WHITE
			loop = false
			print(text)
		} else if (text == "white" || text == "White") {
			Player1 = WHITE
			Player2 = BLACK
			loop = false
		} else {
			fmt.Println("Not a valid input, please try again")
		}
	}
	fmt.Println("Player 1 is", Player1)
	fmt.Println("Player 2 is", Player2)
	game := CreateNewBoard()
	currentplayer := BLACK
	loop = true 
	for !game.gameOver() {

		PrintBoard(game.board)
		whitePoints, blackPoints := game.score()
		fmt.Println("")
		fmt.Println("White:", whitePoints, "     ", "Black:", blackPoints)
		fmt.Println("")
		moves := getMoves(game, currentplayer)
		if Player1 == currentplayer {
			fmt.Println("Player 1, it's your turn!")
		} else if Player2 == currentplayer && !ActivateAI {
			fmt.Println("Player 2, it's your turn!")
		} else if Player2 == currentplayer && ActivateAI {
			fmt.Println("Your opponent's moves are", moves)
			fmt.Print("Your opponent places a ", currentplayer.String(), " piece at: ")
		}
		if len(moves) > 0 && (!ActivateAI || (ActivateAI && currentplayer == Player1)) {
			fmt.Println("Your Available moves are:")
			for i, move := range moves {
			fmt.Print(move)
			if i == len(moves) - 1{
				fmt.Println(".")
			} else {
				fmt.Print(" ")
			}
		}
		}
		if (len(moves) == 0) {
			fmt.Println("No moves available.", currentplayer.String(), "turn is skipped")
			currentplayer = currentplayer.Opposite()
			continue
		}
		loop = true 
		for loop {
			var playerChoice string
			if ActivateAI && currentplayer == Player2 || ActivateAI2 && currentplayer == Player1 {
				if currentplayer == BLACK {
					go concurrentBlack(c, game)
					select{
					case a := <- c:
						playerChoice = a
					case <-time.After(time.Duration(timer * 1000) * time.Millisecond):
						moves = getMoves(game, currentplayer)
						randnum := rand.Intn(len(moves))
						playerChoice = moves[randnum]//DUMB AI
						fmt.Println(" ")
						fmt.Println("AI ran out of time! Choosing a random move..")
						fmt.Println(" ")
						fmt.Print("Random Black move is: ")
					}
				}
				if currentplayer == WHITE {
					go concurrentWhite(c, game)
					select{
					case a := <- c:
						playerChoice = a
					case <-time.After(time.Duration(timer * 1000) * time.Millisecond):
						moves = getMoves(game, currentplayer)
						randnum := rand.Intn(len(moves))
						playerChoice = moves[randnum]//DUMB AI
						fmt.Println(" ")
						fmt.Println("AI ran out of time! Choosing a random move..")
						fmt.Println(" ")
						fmt.Print("Random White move is: ")
					}

				}
				fmt.Println(playerChoice)
			} else {
    			go scan(c)

    			select {
    			case playerChoice = <-c:
    				playerChoice = strings.Replace(playerChoice, "\n", "", -1)
					playerChoice = strings.Replace(playerChoice, "\r", "", -1)
					playerChoice = strings.ToUpper(playerChoice)
				case <-time.After(time.Duration(timer * 1000) * time.Millisecond):
        			fmt.Println("Didn't enter a move quick enough")
        			randnum := rand.Intn(len(moves))
					playerChoice = moves[randnum ]
					fmt.Println("Didn't enter a move quick enough, the move ", playerChoice, " was chosen for you")

    			}
    			}
			if Contains(moves, playerChoice) {
				rawMove, _ := getIndex(playerChoice)
				game.flipAll(currentplayer, rawMove)
				loop = false
				currentplayer = currentplayer.Opposite()
			} else {
				fmt.Println(playerChoice, " is not a valid move, please try again")
			}
		}


	}
	PrintBoard(game.board)
	whiteScore, blackScore := game.score()
	if whiteScore > blackScore {
		fmt.Println("White Wins!")
	} else if blackScore > whiteScore {
		fmt.Println("Black Wins!")
	} else {
		fmt.Println("Its a tie!")
	}

	fmt.Println("Thanks for playing!")
	fmt.Println("White's Score was", whiteScore)
	fmt.Println("Black's Score was", blackScore)
	fmt.Println("Seed was", seedtext)
}
//Values for Player 1 and Player 2.
var Player1 Piece
var Player2 Piece

//Board struct stuff
type Piece int

const (
	EMPTY Piece = iota
	WHITE
	BLACK
)
type Game struct {
	board [64]Piece
}

func (piece Piece) String() string {
	names := [...]string{
		"Empty",
        "White",
    	"Black"}
    if piece > 2 {
    	return "Unknown"
    } else {
    	return names[piece]
    }
}

func (piece Piece) Opposite() Piece {
	if piece == WHITE {
		return BLACK
	} else if piece == BLACK {
		return WHITE
	} else {
		return EMPTY
	}
}
//Returns a Piece RESULT that was found at PLACE. For example if
//a1 is empty, then board.get("a1") will return the piece empty.

func (game Game) get(place string) (result Piece, err error) {
    index, err := getIndex(place)
    if err != nil {
    	return
    }
	result = game.board[index]
	return
}
//Sets a string PLACE to a piece PIECE on the board, for example board.set("a1", BLACK)
//sets a1 to black. This function is destructive.
func (game *Game) set(place string, piece Piece) (err error) {
	index, err := getIndex(place)
	if err != nil {
		return
	}
	game.board[index] = piece
	return
}

func (game *Game) setRaw(place int, piece Piece) (err error) {
	if place < 0 || place > 63 {
		return
	}
	game.board[place] = piece
	return
}
//Translates a string, ie "A1" to the according number
// "A1" = 0, "A2" = 1... etc
func getIndex(place string) (result int, err error) {
	if len(place) != 2 {
		err = errors.New(place + " is not a valid spot on board")
		return
	}
	place = strings.ToUpper(place)
	letter := string(place[0])
	number, _ := strconv.ParseInt(string(place[1]), 0, 64)
	if number > 8 {
		err = errors.New(place + " is not a valid spot on board")
		return
	}
	var letternum int64
	switch letter {
	case "A":
		letternum = 1
	case "B":
		letternum = 2
	case "C":
		letternum = 3
	case "D":
		letternum = 4
	case "E":
		letternum = 5
	case "F":
		letternum = 6
	case "G":
		letternum = 7
	case "H":
		letternum = 8
	default:
		err = errors.New(place + " is not a valid spot on board")
		return
	}
	result = int((number - 1) * 8 + (letternum - 1))
	return
}

func PrintBoard(board [64]Piece) {
	for i := 56; i >= 0; i = i - 8 {
		switch i {
		case 56:
    		fmt.Print("8 ")
		case 48:
    		fmt.Print("7 ")
		case 40:
    		fmt.Print("6 ")
    	case 32:
    		fmt.Print("5 ")
    	case 24:
    		fmt.Print("4 ")
    	case 16:
    		fmt.Print("3 ")
    	case 8:
    		fmt.Print("2 ")
    	case 0:
    		fmt.Print("1 ")
    }
		for j:= 0; j < 8; j++ {
			color := board[i + j]
			switch color {
			case WHITE:
				fmt.Print("W")
			case BLACK:
				fmt.Print("B")
			case EMPTY:
				fmt.Print("-")
			}
			fmt.Print(" ")
		}
		fmt.Println("")
	}
	fmt.Println("  A B C D E F G H")
	return
}

//Creates a Board at the
func CreateNewBoard() (game Game) {
	var new_board [64]Piece
	for i := 0; i < 64; i ++ {
		new_board[i] = EMPTY
	}
	new_board[27] = BLACK
	new_board[28] = WHITE
	new_board[35] = WHITE
	new_board[36] = BLACK
	game.board = new_board
	return game
}

//Returns all the available moves for PLAYER in the current GAME as an array.
//The length of the array is equal to the amount of moves available
//Example: If black can place a piece in E3 and D2, then this function will return
//["E3", "D2"]
func getMoves(game Game, player Piece) []string {
	board := game.board
	var moves []string
	for index, piece := range board {
		foundmove := false
		if piece == EMPTY {
			adjacentPlaces := [...]int{-9, -8, -7, -1, 1, 7, 8, 9}
			for _, adder := range adjacentPlaces {
				//CHECK K
				adjacentIndex := index + adder
				if adjacentIndex > 63 || adjacentIndex < 0 {
					continue
				}
				if adder == -9 || adder == -1 || adder == 7 {
					if index % BOARDSIZE < adjacentIndex % BOARDSIZE { //then valid tile
						continue
					}
				}
				if adder == -7 || adder == 1 || adder == 9 {
					if index % BOARDSIZE > adjacentIndex % BOARDSIZE { //then valid tile
						continue
					}
				}
				isValidMove := checkSandwhich(game, player, index, adder, false)
				if isValidMove {
					foundmove = true
				}

			}
		}
		if foundmove {
			moves = append(moves, TranslateToMove(index))
		}
	}
	return moves
}

func (game Game) gameOver() bool {
	return len(getMoves(game, WHITE)) == 0 && len(getMoves(game, BLACK)) == 0
}

func (game Game) score() (int, int) {
	board := game.board
	black := 0
	white := 0
	for _, tile := range(board) {
		if tile == WHITE {
			white++
		} else if tile == BLACK {
			black++
		}
	}
	return white, black
}

func scan(input chan string) {
	for {
	    in := bufio.NewReader(os.Stdin)
	    result, err := in.ReadString('\n')
	    if err != nil {
	        log.Fatal(err)
	    }

	    input <- result
	}
}

func concurrentBlack(input chan string, game Game) {
	_ , a := minimax(game, DEPTH, false, MinInt, MaxInt)
	input <- a
}

func concurrentWhite(input chan string, game Game) {
	_ , a := minimax(game, DEPTH, true, MinInt, MaxInt)
	input <- a

}