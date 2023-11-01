package views

import (
	"fmt"
	"os"
	"strings"

	"github.com/hasyimzii/go_atm_cli/controllers"
)

func ExitMenu() {
	controllers.UserInput("Are you sure want to exit? [y/n]: ", &controllers.Input)

	switch strings.ToLower(controllers.Input) {
	case "y":
		fmt.Println("Good bye!")
		os.Exit(1)
	case "n":
		MainMenu()
	default:
		controllers.WrongInput()
		MainMenu()
	}
}
