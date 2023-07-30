package model

type Robot struct {
	Position Positioning
	Facing   string
}

type Positioning struct {
	X_Coords int
	Y_Coords int
}

type Table struct {
	MaxX  int
	MaxY  int
	Robot *Robot
}

var Directions = []string{"east", "south", "west", "north"}
