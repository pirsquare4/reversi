package main

import (
	"fmt"
	"bufio"
	"os"
	"errors"
	"strconv"
	"strings"
)  
//Where the games begin!
func main() {
	loop := true
	for loop {
		reader := bufio.New(os.Stdin)
		fmt.Println("Player 1, Choose your Color: White or Black?")
		text, _ := reader.ReadString('\n')
		if (text == "black\n" || text == "Black\n"){
			Player1 = 1
			Player2 = 0
			loop = false
		} else if (text == "white\n" || text == "White\n") {
			Player1 = 0
			Player2 = 1
			loop = false
		} else {
			fmt.Println("Not a valid input, please try again")
		}
	}
	fmt.Println("Player 1 is", Player1)
	fmt.Println("Player 2 is", Player2)
	//PlayGame()
	game := CreateNewBoard()
	PrintBoard(game.board)
	E4, _ := game.get("E4")
	D4, _ := game.get("D4")
	A4, _ := game.get("A4")
	fmt.Println("E4 is", E4)
	fmt.Println("D4 is", D4)
	fmt.Println("A4 is", A4)
	game.set("e4", BLACK)
	game.set("d4", EMPTY)
	game.set("a1", WHITE)
	game.set("H8", BLACK)
	PrintBoard(game.board)
	

	fmt.Println("Thanks for playing!")

}
//Values for Player 1 and Player 2. 
//White is 0, Black is 1.
var Player1 int
var Player2 int

//Board struct stuff
type Piece int

const (
	EMPTY Piece = iota
	WHITE
	BLACK
)
BLACK
WHITE
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
	result = int((letternum - 1) * 8 + (number - 1))
	return 
}

func PrintBoard(board [64]Piece) {
	for i := 56; i >= 0; i = i - 8 {
		switch i {
		case 56:
    		fmt.Print("H ")
		case 48:
    		fmt.Print("G ")
		case 40:
    		fmt.Print("F ")
    	case 32:
    		fmt.Print("E ")
    	case 24:
    		fmt.Print("D ")
    	case 16:
    		fmt.Print("C ")
    	case 8:
    		fmt.Print("B ")
    	case 0:
    		fmt.Print("A ")
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
	fmt.Println("  1 2 3 4 5 6 7 8")
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
		if piece == EMPTY {
			adjacentPlaces = [-9, -8, -7, -1, 1, 7, 8, 9]
			for _, adder := range board {
				//CHECK K
				if adder == -9 || adder == -1 || adder == 7 {

				}
			}
		}
	}
	return moves
}