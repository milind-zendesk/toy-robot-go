package parsingcommands

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"toy-robot-go/model"
	"toy-robot-go/robot"
	"toy-robot-go/table"
)

func ParseCommands(command string, main_table *model.Table) {

	action := strings.ToLower(command)

	switch command {
	case "move":
		robot.Move(main_table)

	case "left":
		robot.Left(main_table.Robot)

	case "right":
		robot.Right(main_table.Robot)

	case "report":
		robot.Report(main_table.Robot)

	case "exit":
		os.Exit(0)

	default:
		parsePlaceCommand(action, main_table)
	}
}

func parsePlaceCommand(action string, main_table *model.Table) {
	re := regexp.MustCompile(`^(place)\s(\d+),(\d+),(\w+)$`)
	match := re.FindStringSubmatch(action)
	if len(match) == 5 {
		x_axis, err := strconv.Atoi(match[2])
		if err != nil || x_axis > main_table.MaxX {
			fmt.Println("Please Enter digit from 0 to 4 ro place the robot on X axis")
		}

		y_axis, err := strconv.Atoi(match[3])
		if err != nil || y_axis > main_table.MaxY {
			fmt.Println("Please Enter digit from 0 to 4 ro place the robot on Y axis")
		}

		if model.IndexOf(match[4]) == -1 {
			fmt.Println("Please enter a valid Direction")
		} else {
			table.PlaceRobotOnTable(main_table, x_axis, y_axis, match[4])
		}
	} else {
		fmt.Println("Please Enter a Valid Command")
	}
}
