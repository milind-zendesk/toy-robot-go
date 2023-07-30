package main

import (
	"bufio"
	"fmt"
	"os"
	"toy-robot-go/model"
	"toy-robot-go/parsingcommands"
)

func main() {
	table := model.Table{
		MaxX: 5,
		MaxY: 5,
	}

	// directions := [4]string{"east", "south", "west", "north"}

	fmt.Println("Place your marker on 5x5 metrics")
	fmt.Println("Accepted input : PLACE <x_coordinate>,<y_coordinate>,<direction>")
	fmt.Println("Then Move the Robot using Left, Right and Move Commands")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		var action string
		scanner.Scan()
		action = scanner.Text()

		parsingcommands.ParseCommands(action, &table)
		//3. Execute as per the given input
		//3. output
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
