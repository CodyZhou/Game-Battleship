package bootstrap

import (
	"Game-Battleship/global"
	"Game-Battleship/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartScanInput() {
	// Set up flag for players
	// true - player 1
	// false - player 2
	flag := true

	input := bufio.NewScanner(os.Stdin)

	fmt.Print(GetPlayerContent(flag), "! ", global.GetMessage("MESSAGE_INPUT_NOTICE_COMMAND_LINE"))

	for input.Scan() {

		line := input.Text()
		line = strings.ToUpper(line)

		flag = CommandAction(line, flag)

		if global.DebugMode {
			fmt.Println("Player 1 ship position: ", global.Player1ShipPosition, " Battleship quantity: ", len(global.Player1ShipPosition))
			fmt.Println("Player 2 ship position: ", global.Player2ShipPosition, " Battleship quantity: ", len(global.Player2ShipPosition))
		}

		// Print the notice command
		fmt.Print(GetPlayerContent(flag), "! ", global.GetMessage("MESSAGE_INPUT_NOTICE_COMMAND_LINE"))

	}

}

// GetPlayerContent is used to set up player content for display message.
func GetPlayerContent(flag bool) string {
	if flag {
		return "Player 1"
	} else {
		return "Player 2"
	}
}

func CommandAction(line string, flag bool) bool {
	// Verify the command of the player
	err := utils.PlayerInputValidation(line)
	if err != nil {
		fmt.Println(err.Error())

	} else {
		switch line {
		case "HELP":
			fmt.Println(global.GetMessage("MESSAGE_HELP_INFO"))
		case "ON":
			global.DebugMode = true
			fmt.Println(global.GetMessage("MESSAGE_DEBUG_MODE_ON"))
		case "OFF":
			global.DebugMode = false
			fmt.Println(global.GetMessage("MESSAGE_DEBUG_MODE_OFF"))
		case "BYE":
			fmt.Println(global.GetMessage("MESSAGE_QUIT_GAME"))
			// quit from game
			os.Exit(0)
		case "GG":
			fmt.Println(GetPlayerContent(flag), global.GetMessage("MESSAGE_ADMIT_DEFEAT"))
			fmt.Println(global.GetMessage("MESSAGE_ANNOUNCE_WINNER") + GetPlayerContent(!flag) + "! \n")
			RestartGame()
		default:
			fmt.Println("You input: ", line)
			utils.AnalyzeCommand(line, flag)

			endGameFlag := utils.CheckWinner()

			if endGameFlag {
				RestartGame()
			}

			flag = !flag
		}

	}

	return flag
}
