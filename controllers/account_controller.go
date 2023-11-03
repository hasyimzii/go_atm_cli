package controllers

import (
	"github.com/hasyimzii/go_atm_cli/models"
)

func GetAccount() models.Account {
	for _, account := range Accounts {
		if account.Id == LoginId {
			return account
		}
	}

	return models.Account{}
}
