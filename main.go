package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	max_width := 4
	max_height := 4
	directions := [4]string{"west", "south", "east", "north"}
	var robot_x int = 0
	var robot_y int = 0
	marker_face := "north"
	fmt.Println("Place your marker on 5x5 metrics")
	fmt.Println("Accepted input : PLACE <x_coordinate>,<y_coordinate>,<direction>")
	fmt.Println("Then Move the marker using Left, Right and Move Commands \n")

	for {
		var action string = ""
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		action = scanner.Text()
		action = strings.ToLower(action)

		action_split := strings.Split(action, " ")
		if len(action_split) > 1 {
			place := strings.Split(action_split[1], ",")
			x_axis, err := strconv.Atoi(place[0])
			if err != nil {
				panic(err)
			}
			y_axis, err := strconv.Atoi(place[1])
			if err != nil {
				panic(err)
			}
			robot_x = x_axis
			robot_y = y_axis
			marker_face = place[2]
		}

		switch action {
		case "move":
			switch marker_face {
			case "north":
				if robot_y == max_height {
					fmt.Println("Robot will go out of the board \n")
				} else {
					fmt.Println("Previous Positions : ", robot_x, ",", robot_y)
					robot_y = robot_y + 1
					fmt.Println("Updated Positions : ", robot_x, ",", robot_y, "\n")
				}
			case "west":
				if robot_x == max_width {
					fmt.Println("Robot will go out of the board \n")
				} else {
					fmt.Println("Previous Positions : ", robot_x, ",", robot_y)
					robot_x = robot_x + 1
					fmt.Println("Updated Positions : ", robot_x, ",", robot_y, "\n")
				}
			case "south":
				if robot_y == 0 {
					fmt.Println("Robot will go out of the board \n")
				} else {
					fmt.Println("Previous Positions : ", robot_x, ",", robot_y)
					robot_y = robot_y - 1
					fmt.Println("Updated Positions : ", robot_x, ",", robot_y, "\n")
				}
			case "east":
				if robot_x == 0 {
					fmt.Println("Robot will go out of the board \n")
				} else {
					fmt.Println("Previous Positions : ", robot_x, ",", robot_y)
					robot_x = robot_x - 1
					fmt.Println("Updated Positions : ", robot_x, ",", robot_y, "\n")
				}
			}
		case "left":
			fmt.Println("Previous Direction ", marker_face)
			if marker_face == directions[0] {
				marker_face = directions[len(directions)-1]
			} else {
				index := indexOf(marker_face, directions[:])
				marker_face = directions[index-1]
			}
			fmt.Println("Updated Direction ", marker_face, "\n")

		case "right":
			fmt.Println("Previous Direction ", marker_face)
			if marker_face == directions[len(directions)-1] {
				marker_face = directions[0]
			} else {
				index := indexOf(marker_face, directions[:])
				marker_face = directions[index+1]
			}
			fmt.Println("Updated Direction ", marker_face, "\n")

		case "report":
			fmt.Println("Current Position of the Robot: ", robot_x, ",", robot_y, " Facing: ", marker_face, "\n")
		}
		if action == "exit" {
			break
		}
	}
	fmt.Println("Program Terminated")
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
