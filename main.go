package main

import (
	"bufio"
	"fmt"
	"os"
	"toy-robot-go/model"
	parsingcommands "toy-robot-go/parsingCommands"
)

func main() {
	table := model.Table{
		MaxX: 5,
		MaxY: 5,
	}

	fmt.Println("Place your marker on 5x5 metrics")
	fmt.Println("Accepted input : PLACE <x_coordinate>,<y_coordinate>,<direction>")
	fmt.Println("Then Move the Robot using Left, Right and Move Commands")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		var action string
		scanner.Scan()
		action = scanner.Text()
		parsingcommands.ParseCommands(action, &table)
	}
}
