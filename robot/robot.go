package robot

import (
	"fmt"
	"toy-robot-go/model"
)

var directions = model.Directions

func Left(robot *model.Robot) {
	fmt.Println("Previous Direction ", robot.Facing)
	if robot.Facing == directions[0] {
		robot.Facing = directions[len(directions)-1]
	} else {
		index := model.IndexOf(robot.Facing)
		robot.Facing = directions[index-1]
	}
	fmt.Println("Updated Direction ", robot.Facing)
}

func Right(robot *model.Robot) {
	fmt.Println("Previous Direction ", robot.Facing)
	if robot.Facing == directions[len(directions)-1] {
		robot.Facing = directions[0]
	} else {
		index := model.IndexOf(robot.Facing)
		robot.Facing = directions[index+1]
	}
	fmt.Println("Updated Direction ", robot.Facing)
}

func Move(table *model.Table) {
	robot := *table.Robot
	positions := robot.Position
	switch robot.Facing {
	case "north":
		if positions.Y_Coords == table.MaxY {
			fmt.Println("Robot will go out of the board")
		} else {
			fmt.Println("Previous Positions : ", positions.X_Coords, ",", positions.Y_Coords)
			positions.Y_Coords += 1
			fmt.Println("Updated Positions : ", positions.X_Coords, ",", positions.Y_Coords)
		}
	case "east":
		if positions.X_Coords == table.MaxX {
			fmt.Println("Robot will go out of the board")
		} else {
			fmt.Println("Previous Positions : ", positions.X_Coords, ",", positions.Y_Coords)
			positions.X_Coords += 1
			fmt.Println("Updated Positions : ", positions.X_Coords, ",", positions.Y_Coords)
		}
	case "south":
		if positions.Y_Coords == 0 {
			fmt.Println("Robot will go out of the board")
		} else {
			fmt.Println("Previous Positions : ", positions.X_Coords, ",", positions.Y_Coords)
			positions.Y_Coords -= 1
			fmt.Println("Updated Positions : ", positions.X_Coords, ",", positions.Y_Coords)
		}
	case "west":
		if positions.X_Coords == 0 {
			fmt.Println("Robot will go out of the board")
		} else {
			fmt.Println("Previous Positions : ", positions.X_Coords, ",", positions.Y_Coords)
			positions.X_Coords -= 1
			fmt.Println("Updated Positions : ", positions.X_Coords, ",", positions.Y_Coords)
		}
	}
	robot.Position = positions
	table.Robot = &robot
}

func Report(robot *model.Robot) {
	positions := robot.Position
	fmt.Println("Current Position of Robot: ", positions.X_Coords, ", ", positions.Y_Coords, ", ", robot.Facing)
}
