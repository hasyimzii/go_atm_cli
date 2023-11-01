package main

import (
	"github.com/hasyimzii/go_atm_cli/controllers"
	"github.com/hasyimzii/go_atm_cli/views"
)

func init() {
	controllers.ClearScreen()
}

func main() {
	views.LoginMenu()
}
