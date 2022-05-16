package main

import (
	"Game-Battleship/bootstrap"
	"Game-Battleship/config"
	"fmt"
)

func main() {
	fmt.Println("********** Welcome to Battleship **********")

	config.InitializeEnv()

	bootstrap.StartGame()

}
