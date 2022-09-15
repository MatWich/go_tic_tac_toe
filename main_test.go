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