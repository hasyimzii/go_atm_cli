package main

import (
	"github.com/hasyimzii/go_atm_cli/controllers"
	"github.com/hasyimzii/go_atm_cli/logs"
	"github.com/hasyimzii/go_atm_cli/views"
)

func init() {
	controllers.ClearScreen()
	logs.SetLogger()
}

func main() {
	views.LoginMenu()
}
