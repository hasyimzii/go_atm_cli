package views

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/controllers"
)

func TransferMenu() {
	var amount int
	var id, pin string

	fmt.Print("Insert id you want to transfer: ")
	fmt.Scanln(&id)

	fmt.Print("Insert amount you want to transfer: Rp ")
	fmt.Scanln(&amount)

	fmt.Print("Insert your pin: ")
	fmt.Scanln(&pin)

	login := controllers.Login(controllers.LoginId, pin)

	if login {
		err := controllers.Transfer(id, amount)
		if err != nil {
			panic(err.Error())
		}
	} else {
		fmt.Println("Wrong pin!")
	}

	controllers.UserInput("[Press enter to go back] ", &controllers.Input)
	MainMenu()
}
