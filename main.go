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
	robot_x := 0
	robot_y := 0
	marker_face := "north"
	fmt.Println("Place your marker on 5x5 metrics")
	fmt.Println("Accepted input : PLACE <x_coordinate>,<y_coordinate>,<direction>")
	fmt.Println("Then Move the Robot using Left, Right and Move Commands \n")

	for {
		var action string = ""
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		action = scanner.Text()
		action = strings.ToLower(action)

		action_split := strings.Split(action, " ")
		if len(action_split) > 1 && action_split[0] == "place" {
			place := strings.Split(action_split[1], ",")
			x_axis, err := strconv.Atoi(place[0])
			if err != nil || x_axis > 4 {
				fmt.Println("Please Enter digit from 0 to 4 ro place the robot on X axis")
			} else {
				robot_x = x_axis
			}

			y_axis, err := strconv.Atoi(place[1])
			if err != nil || y_axis > 4 {
				fmt.Println("Please Enter digit from 0 to 4 ro place the robot on Y axis")
			} else {
				robot_y = y_axis
			}

			if indexOf(place[2], directions[:]) == -1 {
				fmt.Println("Please enter a valid Direction")
			} else {
				marker_face = place[2]
			}
		} else if action == "move" {
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
		} else if action == "left" {
			fmt.Println("Previous Direction ", marker_face)
			if marker_face == directions[0] {
				marker_face = directions[len(directions)-1]
			} else {
				index := indexOf(marker_face, directions[:])
				marker_face = directions[index-1]
			}
			fmt.Println("Updated Direction ", marker_face, "\n")
		} else if action == "right" {
			fmt.Println("Previous Direction ", marker_face)
			if marker_face == directions[len(directions)-1] {
				marker_face = directions[0]
			} else {
				index := indexOf(marker_face, directions[:])
				marker_face = directions[index+1]
			}
			fmt.Println("Updated Direction ", marker_face, "\n")
		} else if action == "report" {
			fmt.Println("Current Position of the Robot: ", robot_x, ",", robot_y, " Facing: ", marker_face, "\n")
		} else if action == "exit" {
			break
		} else {
			fmt.Println("Please enter a valid command\n")
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
