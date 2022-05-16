package utils

import (
	"Game-Battleship/global"
	"fmt"
)

// AnalyzeCommand is used to check the player's command and print out related message
func AnalyzeCommand(coordinate string, flag bool) {
	// flag = true - player 1
	// flag = false - player 2

	if flag {
		// Player1's command
		CheckBattleShip(coordinate, global.Player2ShipPosition)

	} else {
		// Player2's command
		CheckBattleShip(coordinate, global.Player1ShipPosition)
	}

}

// CheckBattleShip is used to check the status of the player's battleship
func CheckBattleShip(coordinate string, shipPosition map[int][]string) {
	for key, value := range shipPosition {
		for index, item := range value {
			if item == coordinate {
				fmt.Println(global.GetMessage("MESSAGE_HIT_BATTLESHIP"))

				// Set the coordinate to empty
				value[index] = ""

				// Check all position of this battleship
				for i := 0; i < len(value); i++ {
					if value[i] != "" {
						return
					}
				}

				// Remove sank battleship
				delete(shipPosition, key)

				fmt.Println(global.GetMessage("MESSAGE_SANK_BATTLESHIP"))

				return
			}
		}

	}

	fmt.Println(global.GetMessage("MESSAGE_MISS_TARGET"))
}

// CheckWinner is used to check who is the winner.
func CheckWinner() bool {
	if len(global.Player1ShipPosition) == 0 {
		fmt.Println(global.GetMessage("MESSAGE_ANNOUNCE_WINNER") + " Player 2! ")
		fmt.Println(global.GetMessage("MESSAGE_ANNOUNCE_WINNER") + " Player 2! ")
		fmt.Println(global.GetMessage("MESSAGE_ANNOUNCE_WINNER") + " Player 2! \n")

		return true
	}

	if len(global.Player2ShipPosition) == 0 {
		fmt.Println(global.GetMessage("MESSAGE_ANNOUNCE_WINNER") + " Player 1! ")
		fmt.Println(global.GetMessage("MESSAGE_ANNOUNCE_WINNER") + " Player 1! ")
		fmt.Println(global.GetMessage("MESSAGE_ANNOUNCE_WINNER") + " Player 1! \n")

		return true
	}

	// The game is not over
	return false
}
