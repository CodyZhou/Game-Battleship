package global

var MessageDictionary = map[string]string{
	// System information
	"SYSTEM_INPUT": "Please Input: ",

	// Player Input
	"MESSAGE_INPUT_NOTICE_COMMAND_LINE": "Please input the command (Example: B5. input 'help' for help): ",

	// Help Information
	"MESSAGE_HELP_INFO": "\t help \t Print help information \n" +
		"\t bye \t Quit the game \n" +
		"\t gg \t Admit defeat and give up the game \n" +
		"\t A-J \t Columns (not case sensitive) \n" +
		"\t 1-10 \t Rows \n" +
		"\n" +
		"\t on \t Turn on the DEBUG mode, game will display more information\n" +
		"\t off \t Turn off the DEBUG mode\n" +
		"\n " +
		"\t NOTICE: This game is not case-sensitive. So, 'help' is the same as 'HELP'! ",

	"MESSAGE_ADMIT_DEFEAT": " admit defeat! Thank you for playing!\n",

	"MESSAGE_QUIT_GAME": "Thank you for playing...\n",

	"MESSAGE_RESTART_GAME":      "Game restarting ... \n",
	"MESSAGE_RESTART_GAME_DONE": "Game has been restarted! Enjoy it!\n",

	"MESSAGE_MISS_TARGET":     "You missed the target\n",
	"MESSAGE_HIT_BATTLESHIP":  "You hit an enemy's battleship!\n",
	"MESSAGE_SANK_BATTLESHIP": "You SANK an enemy's battleship! What an amazing operation!\n",

	"MESSAGE_ANNOUNCE_WINNER": "Congratulations! The WINNER is ",

	"MESSAGE_DEBUG_MODE_ON":  "DEBUG mode is on!",
	"MESSAGE_DEBUG_MODE_OFF": "DEBUG mode is off!",

	// Error message
	"ERROR_COMMAND_INPUT_ERROR": "Command not recognized! Please input 'help' for more information.",
}

func GetMessage(messageFlag string) string {
	return MessageDictionary[messageFlag]
}
