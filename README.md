# Go ATM CLI
Simple ATM (Automated Teller Machine) app in CLI using Go

## How to install (using Go)
(Must've installed: Go Programming Language)

1. Clone this repository
```sh
git clone https://github.com/hasyimzii/go_atm_cli.git
```
2. Open repository in terminal
```sh
cd go_atm_cli
```
3. Build Go app
```go
go build
```
4. Excecute binary file
```sh
./go_atm_cli
```
5. Now the app is running

## How to install (using Docker)
(Must've installed: Docker)

1. Clone this repository
```sh
git clone https://github.com/hasyimzii/go_atm_cli.git
```
2. Open repository in terminal
```sh
cd go_atm_cli
```
3. Excecute docker compose command
```sh
docker composer up
```
4. Now the code is running
5. (If you want to remove the container & image) 
```sh
docker compose down
docker image rm go_atm_cli:1.0.0
```

## Login credentials
- Credential file [./models/data.json](https://github.com/hasyimzii/go_atm_cli/tree/main/models/data.json)
- The credential for login is ```id``` and ```pin``` of the user

## Features
- This app has some features, such as transfer, deposit, withdraw, account
- You must type the menu number to proceed to the menu
- Transfer: send money to target account
- Deposit: increase the balance of your account
- Withdraw: decrease the balance of your account
- Account: get you account details

## Logging
- Logging folder [./logs](https://github.com/hasyimzii/go_atm_cli/tree/main/logs)
- The logging output is stored in ```app.log``` file