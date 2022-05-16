package test

import (
	"Game-Battleship/bootstrap"
	"fmt"
	"testing"
)

func TestPlayerInput(t *testing.T) {
	fmt.Println("Test player input")

	t.Run("Player 1 input command 'HELP'", func(t *testing.T) {
		command := "HELP"
		flag := true

		returnFlag := bootstrap.CommandAction(command, flag)

		if flag != returnFlag {
			t.Errorf("Expect player is %s, but got %s!", bootstrap.GetPlayerContent(flag), bootstrap.GetPlayerContent(returnFlag))
		}

		fmt.Println("")
		fmt.Println("")
	})

	t.Run("Player 2 input command 'HELP'", func(t *testing.T) {
		command := "HELP"
		flag := false

		returnFlag := bootstrap.CommandAction(command, flag)

		if flag != returnFlag {
			t.Errorf("Expect player is %s, but got %s!", bootstrap.GetPlayerContent(flag), bootstrap.GetPlayerContent(returnFlag))
		}

		fmt.Println("")
		fmt.Println("")
	})

	t.Run("Player 1 input command 'GG'", func(t *testing.T) {
		command := "GG"
		flag := true

		returnFlag := bootstrap.CommandAction(command, flag)

		if flag != returnFlag {
			t.Errorf("Expect player is %s, but got %s!", bootstrap.GetPlayerContent(flag), bootstrap.GetPlayerContent(returnFlag))
		}

		fmt.Println("")
		fmt.Println("")
	})

	t.Run("Player 2 input command 'GG'", func(t *testing.T) {
		command := "GG"
		flag := false

		returnFlag := bootstrap.CommandAction(command, flag)

		if flag != returnFlag {
			t.Errorf("Expect player is %s, but got %s!", bootstrap.GetPlayerContent(!flag), bootstrap.GetPlayerContent(returnFlag))
		}

		fmt.Println("")
		fmt.Println("")
	})

	t.Run("Player 1 input command 'B1'", func(t *testing.T) {
		command := "B1"
		flag := true

		returnFlag := bootstrap.CommandAction(command, flag)

		if flag == returnFlag {
			t.Errorf("Expect player is %s, but got %s!", bootstrap.GetPlayerContent(!flag), bootstrap.GetPlayerContent(returnFlag))
		}

		fmt.Println("")
		fmt.Println("")
	})

	t.Run("Player 2 input command 'G1'", func(t *testing.T) {
		command := "G1"
		flag := false

		returnFlag := bootstrap.CommandAction(command, flag)

		if flag == returnFlag {
			t.Errorf("Expect player is %s, but got %s!", bootstrap.GetPlayerContent(!flag), bootstrap.GetPlayerContent(returnFlag))
		}

		fmt.Println("")
		fmt.Println("")
	})

}
