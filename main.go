package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"terminus_plus/core"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	prompt := "Term -> "

	for {
		fmt.Print(prompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur de lecture :", err)
			continue
		}

		input = strings.TrimSpace(input)
		SimulateCommand(input)
	}
}

func SimulateCommand(input string) core.Terminus {

	var cmdExec *exec.Cmd

	t := core.Terminus{
		Input:  input,
		Output: "",
	}

	parts := strings.Fields(t.Input)

	if len(parts) == 0 {
		fmt.Println("No input provided.")
		return t
	}
	cmd := parts[0]
	arg := parts[1:]

	exec.Command(cmd, arg...)

	if runtime.GOOS == "windows" {
		cmdExec = exec.Command("cmd", append([]string{"/C"}, parts...)...)
	} else {
		cmdExec = exec.Command(parts[0], parts[1:]...)
	}

	out, err := cmdExec.Output()
	if err != nil {
		fmt.Println("Erreur d'exécution de la commande :", err)
	} else {
		fmt.Println(string(out))
	}

	fmt.Println(string(out))

	return t
}

func SaveToFile(t core.Terminus) core.Terminus {
	c := core.Terminus{
		Input:  "Hello",
		Output: "World",
	}

	return c
}
