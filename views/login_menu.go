package views

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/controllers"
)

func LoginMenu() {
	fmt.Println("Login into your account")

	var id, pin string
	fmt.Print("Type your id: ")
	fmt.Scanln(&id)
	fmt.Print("Type your pin: ")
	fmt.Scanln(&pin)

	login := controllers.Login(id, pin)

	if login {
		controllers.UserInput("[Login success! Press enter to continue] ", &controllers.Input)
		MainMenu()
	} else {
		controllers.UserInput("[Login failed! Press enter to go back] ", &controllers.Input)
		LoginMenu()
	}
}
