package parsingcommands

import (
	"fmt"
	"strconv"
	"strings"
	"toy-robot-go/model"
	"toy-robot-go/robot"
	"toy-robot-go/table"
)

func ParseCommands(command string, main_table *model.Table) {
	directions := model.Directions
	action := strings.ToLower(command)

	action_split := strings.Split(action, " ")
	if len(action_split) > 1 && action_split[0] == "place" {

		place := strings.Split(action_split[1], ",")
		x_axis, err := strconv.Atoi(place[0])
		if err != nil || x_axis > main_table.MaxX {
			fmt.Println("Please Enter digit from 0 to 4 ro place the robot on X axis")
		}

		y_axis, err := strconv.Atoi(place[1])
		if err != nil || y_axis > main_table.MaxY {
			fmt.Println("Please Enter digit from 0 to 4 ro place the robot on Y axis")
		}

		if indexOf(place[2], directions) == -1 {
			fmt.Println("Please enter a valid Direction")
		}
		table.PlaceRobotOnTable(main_table, x_axis, y_axis, place[2])

	} else if action == "move" {
		robot.Move(main_table)
	} else if action == "left" {
		robot.Left(main_table.Robot)
	} else if action == "right" {
		robot.Right(main_table.Robot)
	} else if action == "report" {
		robot.Report(main_table.Robot)
	} else {
		fmt.Println("Please enter a valid command")
	}
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
