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
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Player 1, Choose your Color: White or Black?")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = text[0:5]
		if (text == "black" || text == "Black"){
			Player1 = 1
			Player2 = 0
			loop = false
			print(text)
		} else if (text == "white" || text == "White") {
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
	fmt.Println("E4 is", E4)
	fmt.Println("D4 is", D4)
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
    if len(place) != 2 {
		err = errors.New(place + " is not a valid spot on board")
		return
	}
	place = strings.ToUpper(place)
	letter := string(place[0])
	number, _ := strconv.ParseInt(string(place[1]), 0, 64)
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
	}
	result = game.board[(letternum - 1) * 8 + (number - 1)]
	return
}

//Sets a string PLACE to a piece PIECE on the board, for example board.set("a1", BLACK)
//sets a1 to black. This function is destructive.
func (game *Game) set(place string, piece Piece) (err error) {
	if len(place) != 2 {
		err = errors.New(place + " is not a valid spot on board")
		return
	}
	place = strings.ToUpper(place)
	letter := string(place[0])
	number, _ := strconv.ParseInt(string(place[1]), 0, 64)
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
	}
	game.board[(letternum - 1) * 8 + (number - 1)] = piece
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