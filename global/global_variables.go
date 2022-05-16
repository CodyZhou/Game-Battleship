package global

var Columns = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
var Rows = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

// Player1ShipPosition is used to save the list of battleship position of player 1
var Player1ShipPosition map[int][]string

// Player2ShipPosition is used to save the list of battleship position of player 2
var Player2ShipPosition map[int][]string

// DebugMode is used to set up debug mode for the game. Debug mode will display more information.
var DebugMode bool = false
