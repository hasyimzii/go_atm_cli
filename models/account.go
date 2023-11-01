package models

type Account struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Pin     string `json:"pin"`
	Balance int    `json:"balance"`
}
