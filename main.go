package main

import ( 
	"fmt"
)
var player_x string = "X"
var player_o string = "O"

func main() {
	fmt.Println("TIC TAC TOE")
	b := Board{}
	b.initBoard()
	b.print()
	fmt.Println(player_o, player_x)
	
}


type Board struct {
	fields [9]string
}

func (b *Board) initBoard() {
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
	