package main

import (
	"fmt"
)

func main() {
	for {
		var command string = ""
		_, err := fmt.Scanln(&command)
		if err != nil {
			panic(err)
		}
		switch command {
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
	fmt.Printf("Done")
}
