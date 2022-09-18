package main

import "testing"

func TestWinningConditionHorizontal(t *testing.T) {
	b := Board{}
	b.fields[0]="X"
	b.fields[1]="X"
	b.fields[2]="X"

	if !b.check_win("X") {
		t.Errorf("Horizontal win check failed")
	}
}

func TestWinningConditionVertical(t *testing.T) {
	b := Board{}
	b.fields[0]="X"
	b.fields[3]="X"
	b.fields[6]="X"

	if !b.check_win("X") {
		t.Errorf("Vertical win check failed")
	}
}

func TestWinningConditionDiagonal(t *testing.T) {
	b := Board{}
	b.fields[0]="X"
	b.fields[3]="X"
	b.fields[6]="X"

	if !b.check_win("X") {
		t.Errorf("Diagonal win check failed")
	}
}

func TestBoardInitialization(t *testing.T) {
	b := Board{}
	b.initBoard()

	for _, field := range b.fields {
		if field != " " {
			t.Errorf("Field is has %v symbol instead of %v", field, " ")
		}
	}
}

func TestBoardInitializationAfterGame(t *testing.T) {
	b := Board{}
	b.initBoard()
	b.fields[1], b.fields[3] = "X", "X"
	b.initBoard()

	for _, field := range b.fields {
		if field != " " {
			t.Errorf("Field is has %v symbol instead of %v", field, " ")
		}
	}
}

func TestBoardFull(t *testing.T) {
	b := Board{}
	b.initBoard()
	for index := range b.fields {
		b.fields[index] = "X"
	}

	if b.isFull() {
		t.Errorf("Board is full but isFull method thinks otherwise %v", b.fields)
	}
}

func TestBoardPutSign(t *testing.T) {
	var player_x string = "X"
	// var player_o string = "O"
	b := Board{}
	b.currentPlayer = player_x
	b.putSign(1)
	b.putSign(2)

	if player_o != b.fields[1] {
		t.Errorf("Expected %v got %v", player_o, b.fields[1])
	}

	if player_x != b.fields[0] {
		t.Errorf("Expected %v current %v", player_x, b.fields[0])
	}
}

func TestBoardPutSignPlayerSwitch(t *testing.T) {
	b := Board{}
	var player_x string = "X"
	b.currentPlayer = player_x
	b.putSign(1)

	if b.currentPlayer == player_x {
		t.Errorf("After player putting sign on field there should be switch")
	}
}

func TestFieldValid(t *testing.T) {
	b := Board{}
	b.initBoard()
	b.putSign(1)
	// one less because its decreased in put sign
	if b.isValidField(0) == true {
		t.Errorf("After geting a mark in a field user shouldn't be able to mark the same field again")
	}
}