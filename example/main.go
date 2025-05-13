package main

import (
	"fmt"
	"github.com/nexidian/gocliselect"
)

func main() {
	menu := gocliselect.NewMenu(fmt.Sprintf("Chose a colour\nThis is new line"))

	menu.AddItem("Red", 1)
	menu.AddItem("Blue", 2)
	menu.AddItem("Green", 3)
	menu.AddItem("Yellow", 4)
	menu.AddItem("Red", 5)
	menu.AddItem("Blue", 6)
	menu.AddItem("Green", 7)
	menu.AddItem("Yellow", 8)
	menu.EnableSkip(nil)

	result, err := menu.Display()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	if _, ok := result.(int); ok {
		fmt.Printf("Selected option: %d\n", result)
	} else if _, ok := result.(string); ok {
		fmt.Printf("Selected option: %s\n", result)
	} else if result == nil {
		fmt.Printf("Select is skipped.")
	} else {
		fmt.Printf("Selected option of unexpected type: %T with value: %v\n", result, result)
	}
}
