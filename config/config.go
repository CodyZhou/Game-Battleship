package config

import (
	"Game-Battleship/global"
	"Game-Battleship/utils"
	"fmt"
)

// InitializeEnv is used to initialize the game environment.
func InitializeEnv() {

	// Empty players' battleship position
	global.Player1ShipPosition = make(map[int][]string, 0)
	global.Player2ShipPosition = make(map[int][]string, 0)

	InitializeBattleShipPosition("player1")
	InitializeBattleShipPosition("player2")

	// For debug mode
	if global.DebugMode {
		fmt.Println("Player 1 ship position: ", global.Player1ShipPosition, " Battleship quantity: ", len(global.Player1ShipPosition))
		fmt.Println("Player 2 ship position: ", global.Player2ShipPosition, " Battleship quantity: ", len(global.Player2ShipPosition))
	}

	fmt.Println("Game initialize ENV ... successfully!")

}

//InitializeBattleShipPosition is used to set up the players' battleship position automatically.
func InitializeBattleShipPosition(playerFlag string) {
	// Setup battleship coordinate from 5 to 2
	shipSize := 5
	for shipSize >= 2 {

		for {
			tempShipPosition := utils.GetShipPosition(shipSize)

			// Verify that position overlap
			if playerFlag == "player1" {
				flag := utils.VerifyCoordinates(global.Player1ShipPosition, tempShipPosition)
				if flag == false {
					global.Player1ShipPosition[shipSize] = tempShipPosition
					break
				}
			}

			if playerFlag == "player2" {
				flag := utils.VerifyCoordinates(global.Player2ShipPosition, tempShipPosition)
				if flag == false {
					global.Player2ShipPosition[shipSize] = tempShipPosition
					break
				}
			}

		}

		shipSize--
	}
}
