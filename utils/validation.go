package utils

import (
	"Game-Battleship/global"
	"errors"
	"strconv"
	"strings"
)

func PlayerInputValidation(command string) error {
	// Verify the command
	commandList := []string{
		"HELP",
		"BYE",
		"GG",
		"ON",
		"OFF",
	}

	for _, value := range commandList {
		if value == command {
			return nil
		}
	}

	// Transfer to upper
	command = strings.ToUpper(command)

	// Verify the coordinate
	if len(command) == 2 || len(command) == 3 {
		// Verify the columns.
		// Accept from A to J
		if command[0] < 65 || command[0] > 74 {
			return errors.New(global.GetMessage("ERROR_COMMAND_INPUT_ERROR"))
		}

		// Verify the rows.
		// Accept from 1 to 10
		tempRowsInt, err := strconv.Atoi(command[1:len(command)])
		if err != nil {
			return errors.New(global.GetMessage("ERROR_COMMAND_INPUT_ERROR"))
		}

		if tempRowsInt < 1 || tempRowsInt > 10 {
			return errors.New(global.GetMessage("ERROR_COMMAND_INPUT_ERROR"))
		}

		return nil
	} else {
		return errors.New(global.GetMessage("ERROR_COMMAND_INPUT_ERROR"))
	}

}
