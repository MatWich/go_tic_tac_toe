package main

import "fmt"

func main() {
	fmt.Println("TIC TAC TOE")
	b := Board{}
	for i := range b.fields {
		b.fields[i] = "x"
	}

	b.print()
}


type Board struct {
	fields [9]string
}

func (b *Board) print() {
	fmt.Printf("%v | %v | %v\n", b.fields[0], b.fields[1], b.fields[2])
	fmt.Printf("---------\n")
	fmt.Printf("%v | %v | %v\n", b.fields[3], b.fields[4], b.fields[5])
	fmt.Printf("---------\n")
	fmt.Printf("%v | %v | %v\n", b.fields[6], b.fields[7], b.fields[8])
}