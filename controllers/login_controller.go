package controllers

import (
	"github.com/hasyimzii/go_atm_cli/models"
	"github.com/hasyimzii/go_atm_cli/repositories"
	"github.com/sirupsen/logrus"
)

var LoginId, LoginName string
var Accounts []models.Account

func Login(id, pin string) bool {
	accountRepository := repositories.NewAccountRepo()

	var err error

	Accounts, err = accountRepository.FindAll()
	if err != nil {
		panic(err.Error())
	}

	for _, account := range Accounts {
		if account.Id == id && account.Pin == pin {
			LoginId = account.Id
			LoginName = account.Name
			return true
		}
	}

	logrus.Info("[Account Id: ", id, "] Login failed!")
	return false
}
