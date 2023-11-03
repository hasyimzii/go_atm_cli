package repository

import (
	"encoding/json"
	"os"

	"github.com/hasyimzii/go_atm_cli/models"
)

type AccountRepository interface {
	FindAll() ([]models.Account, error)
	Update(accounts []models.Account) error
}

type accountRepositoryImpl struct {
	file string
}

func NewAccountRepository() AccountRepository {
	return &accountRepositoryImpl{file: "./models/data.json"}
}

func (repo *accountRepositoryImpl) FindAll() ([]models.Account, error) {
	file, err := os.ReadFile(repo.file)
	if err != nil {
		return nil, err
	}

	var accounts []models.Account
	err = json.Unmarshal([]byte(file), &accounts)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (repo *accountRepositoryImpl) Update(accounts []models.Account) error {
	data, err := json.Marshal(accounts)
	if err != nil {
		return err
	}

	err = os.WriteFile(repo.file, data, 0666)
	if err != nil {
		return err
	}

	return nil
}
