package robot

import (
	"fmt"
	"toy-robot-go/model"
)

func Left(robot *model.Robot) {
	directions := model.Directions
	fmt.Println("Previous Direction ", robot.Facing)
	if robot.Facing == directions[0] {
		robot.Facing = directions[len(directions)-1]
	} else {
		index := indexOf(robot.Facing, directions[:])
		robot.Facing = directions[index-1]
	}
	fmt.Println("Updated Direction ", robot.Facing)
}

func Right(robot *model.Robot) {
	directions := model.Directions
	fmt.Println("Previous Direction ", robot.Facing)
	if robot.Facing == directions[len(directions)-1] {
		robot.Facing = directions[0]
	} else {
		index := indexOf(robot.Facing, directions[:])
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
			fmt.Println("Robot will go out of the board \n")
		} else {
			fmt.Println("Previous Positions : ", positions.X_Coords, ",", positions.Y_Coords)
			positions.Y_Coords += 1
			fmt.Println("Updated Positions : ", positions.X_Coords, ",", positions.Y_Coords, "\n")
		}
	case "east":
		if positions.X_Coords == table.MaxX {
			fmt.Println("Robot will go out of the board \n")
		} else {
			fmt.Println("Previous Positions : ", positions.X_Coords, ",", positions.Y_Coords)
			positions.X_Coords += 1
			fmt.Println("Updated Positions : ", positions.X_Coords, ",", positions.Y_Coords, "\n")
		}
	case "south":
		if positions.Y_Coords == 0 {
			fmt.Println("Robot will go out of the board \n")
		} else {
			fmt.Println("Previous Positions : ", positions.X_Coords, ",", positions.Y_Coords)
			positions.Y_Coords -= 1
			fmt.Println("Updated Positions : ", positions.X_Coords, ",", positions.Y_Coords, "\n")
		}
	case "west":
		if positions.X_Coords == 0 {
			fmt.Println("Robot will go out of the board \n")
		} else {
			fmt.Println("Previous Positions : ", positions.X_Coords, ",", positions.Y_Coords)
			positions.X_Coords -= 1
			fmt.Println("Updated Positions : ", positions.X_Coords, ",", positions.Y_Coords, "\n")
		}
	}
	robot.Position = positions
	table.Robot = &robot
}

func Report(robot *model.Robot) {
	positions := robot.Position
	fmt.Println("Current Position of Robot: ", positions.X_Coords, ", ", positions.Y_Coords, ", ", robot.Facing)
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
