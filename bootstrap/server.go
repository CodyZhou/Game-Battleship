package bootstrap

import (
	"Game-Battleship/config"
	"Game-Battleship/global"
	"fmt"
	"time"
)

func StartGame() {
	// Game start
	StartScanInput()

}

func RestartGame() {
	fmt.Println(global.GetMessage("MESSAGE_RESTART_GAME"))
	time.Sleep(3 * time.Second)

	config.InitializeEnv()

	fmt.Println(global.GetMessage("MESSAGE_RESTART_GAME_DONE"))

	StartGame()
}
