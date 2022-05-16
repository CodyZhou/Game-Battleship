package test

import (
	"Game-Battleship/bootstrap"
	"Game-Battleship/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.InitializeEnv()

	bootstrap.StartGame()

	code := m.Run()

	os.Exit(code)
}
