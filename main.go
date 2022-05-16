package main

import (
	"Game-Battleship/bootstrap"
	"Game-Battleship/config"
	"fmt"
)

func main() {
	fmt.Println("***************************************************************")
	fmt.Println("**                                                           **")
	fmt.Println("**                   Welcome to Battleship                   **")
	fmt.Println("**                     Author: Cody Zhou                     **")
	fmt.Println("**                codysbusiness050917@gmail.com              **")
	fmt.Println("**                                                           **")
	fmt.Println("***************************************************************")

	config.InitializeEnv()

	bootstrap.StartGame()

}
