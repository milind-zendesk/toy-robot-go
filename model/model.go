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

func IndexOf(element string) int {
	for k, v := range Directions {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
