package controllers

import (
	"fmt"

	"github.com/hasyimzii/go_atm_cli/models"
	"github.com/hasyimzii/go_atm_cli/repository"
	"github.com/sirupsen/logrus"
)

func Withdraw(amount int) error {
	for index, account := range Accounts {
		if account.Id == LoginId {
			if amount > account.Balance {
				msg := "Withdraw failed, insufficient balance!"
				logrus.Warn("[Account: ", LoginName, "] ", msg)
				return models.ValidationError(msg)
			} else {
				(&Accounts[index]).Balance -= amount

				// update data
				accountRepository := repository.NewAccountRepo()
				err := accountRepository.Update(Accounts)
				if err != nil {
					return err
				}

				msg := fmt.Sprint("Success withdraw Rp ", amount)
				logrus.Info("[Account: ", LoginName, "] ", msg)
				fmt.Println(msg)
				return nil
			}
		}
	}

	return models.NotFoundError("Withdraw failed, id not found!")
}
