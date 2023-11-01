package controllers

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var Input string

func UserInput(text string, input *string) {
	fmt.Print(text)
	fmt.Scanln(input)
	ClearScreen()
}

func WrongInput() {
	fmt.Println("[Alert: Wrong type input!]")
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
