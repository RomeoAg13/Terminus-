package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"terminus_plus/core"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	SimulateCommand("Hello")
}

func SimulateCommand(input string) core.Terminus {
	t := core.Terminus{
		Input:  input,
		Output: "World",
	}

	fmt.Printf("Input: %s\n", t.Input)
	fmt.Printf("Output: %s\n", t.Output)

	return t
}

func SaveToFile(t core.Terminus) core.Terminus {
	c := core.Terminus{
		Input:  "Hello",
		Output: "World",
	}

	return c
}
