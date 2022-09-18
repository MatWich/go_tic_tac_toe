package main

import ( 
	"fmt"
	"os"
)
var player_x string = "X"
var player_o string = "O"

func main() {
	fmt.Println("TIC TAC TOE")
	b := Board{}
	b.initBoard()
	b.print()
	var choice string

	for !b.isFull() {
		fmt.Print("Choose field: ")
		fmt.Scanln(&choice)
		getUserInput(&b, choice)
		b.print()
		if b.check_win(player_o) {
			fmt.Printf("%v is a winner", player_o)
			os.Exit(0)
		} else if b.check_win(player_x) {
			fmt.Printf("%v is a winner", player_x)
			os.Exit(0)
		}
	}
	
}

func getUserInput (b *Board, input string) {
	switch input {
	case "h": 
		fmt.Printf("%v | %v | %v\n", 1, 2, 3)
		fmt.Printf("---------\n")
		fmt.Printf("%v | %v | %v\n", 4, 5, 6)
		fmt.Printf("---------\n")
		fmt.Printf("%v | %v | %v\n", 7, 8, 9)
	case "1":
		b.putSign(1)
	case "2":
		b.putSign(2)
	case "3":
		b.putSign(3)
	case "4":
		b.putSign(4)
	case "5":
		b.putSign(5)
	case "6":
		b.putSign(6)
	case "7":
		b.putSign(7)
	case "8":
		b.putSign(8)
	case "9":
		b.putSign(9)
	default:
		fmt.Errorf("invalid option\n")
	}
}

type Board struct {
	fields [9]string
	currentPlayer string
}

func (b *Board) isValidField(index int) bool {
	return b.fields[index] == " "
}

func (b *Board) isFull() bool {
	for index := range b.fields {
		if b.fields[index] == " " {
			return false
		}
	}
	return false
}
func (b *Board) initBoard() {
	b.currentPlayer=player_x
	for index := range b.fields {
		b.fields[index]=" "
	}
}

func (b *Board) print() {
	fmt.Printf("%v | %v | %v\n", b.fields[0], b.fields[1], b.fields[2])
	fmt.Printf("---------\n")
	fmt.Printf("%v | %v | %v\n", b.fields[3], b.fields[4], b.fields[5])
	fmt.Printf("---------\n")
	fmt.Printf("%v | %v | %v\n", b.fields[6], b.fields[7], b.fields[8])
}

func (b *Board) putSign(place int) {
	place -= 1
	if b.isValidField(place) {
		b.fields[place] = b.currentPlayer
	} else {
		return
	}
	if b.currentPlayer == player_x {
		b.currentPlayer = player_o
	} else {
		b.currentPlayer = player_x
	}
}

func (b *Board) check_win(player string) bool {
	// DIAGONAL
	if b.fields[0] == player && b.fields[4] == player && b.fields[8] == player {
		return true
	}
	if b.fields[2] == player && b.fields[4] == player && b.fields[6] == player {
		return true
	}
	// VERTICAL
	if b.fields[0] == player && b.fields[3] == player && b.fields[6] == player {
		return true
	}
	if b.fields[1] == player && b.fields[4] == player && b.fields[7] == player {
		return true
	}
	if b.fields[2] == player && b.fields[5] == player && b.fields[8] == player {
		return true
	}
	// HORIZONTAL
	if b.fields[0] == player && b.fields[1] == player && b.fields[2] == player {
		return true
	}
	if b.fields[3] == player && b.fields[4] == player && b.fields[5] == player {
		return true
	}
	if b.fields[6] == player && b.fields[7] == player && b.fields[8] == player {
		return true
	}
	return false
}
	