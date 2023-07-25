package main

import (
	"fmt"
	"strings"
)

func main() {
	max_width := 5
	max_height := 5
	directions := [4]string{"east", "west", "north", "sputh"}
	robot_x := 0
	robot_y := 0
	fmt.Println(max_height, max_width, directions, robot_x, robot_y)
	fmt.Println("Place your marker on 5x5 metrics")
	fmt.Println("Accepted input : PLACE <x_coordinate>, <y_coordinate>, <direction>")
	for {
		var action string
		_, err := fmt.Scan(&action)
		if err != nil {
			panic(err)
		}
		action = strings.ToLower(action)
		fmt.Println(action)
		action_split := strings.Split(action, " ")
		fmt.Println(action_split)
		if len(action_split) > 0 {
			asd := strings.Split(action_split[1], ",")
			fmt.Println(asd)
			// }
			// place_match, err := regexp.MatchString("place", action)
			// if place_match {
			// 	splitted_action := strings.Split(action, ",")
			// 	fmt.Println(splitted_action)
		} else {
			switch action {
			case "move":
				fmt.Printf("Move")
			case "left":
				fmt.Printf("Left")
			case "right":
				fmt.Printf("Right")
			case "report":
				fmt.Printf("Left")
			case "exit":
				break
			}
		}
	}
	fmt.Printf("Done")
}
