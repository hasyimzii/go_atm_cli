package views

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/controllers"
)

func LoginMenu() {
	var id, pin string

	fmt.Println("Login into your account")
	fmt.Print("Type your id: ")
	fmt.Scanln(&id)
	fmt.Println(id)

	fmt.Print("Type your pin: ")
	fmt.Scanln(&pin)
	fmt.Println(pin)

	// check account

	controllers.UserInput("ENTER", &controllers.Input)
	MainMenu()
}
