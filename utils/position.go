package utils

import (
	"Game-Battleship/global"
	"math/rand"
	"strconv"
	"time"
)

// GetShipPosition is used to get and return the list of battleship's coordinate
// size is the battleship's length. For example: 2, 3, 4, 5
// []string is the list of battleship's coordinate. For example: ["B5", "B6", "B7"]
func GetShipPosition(size int) []string {
	coordinates := make([]string, 0)

	// tempCoordinateMap is used to save the coordinate list temporary
	// To get the final coordinate for the battleship
	tempCoordinateMap := make(map[int][]string)
	tempMapKey := 0

	// Get base coordinate
	rand.Seed(time.Now().UnixNano())
	columnIndex := rand.Intn(len(global.Columns))
	column := global.Columns[columnIndex]
	row := global.Rows[rand.Intn(len(global.Rows))]

	// For test
	//baseCoordinate := column + row
	//fmt.Println("Base coordinate: ", baseCoordinate)

	// for up. column is unchanged, row reduced size
	upCoordinates := make([]string, 0)
	tempRowInt, _ := strconv.Atoi(row)
	if tempRowInt-size >= 0 {
		for i := 0; i < size; i++ {
			upCoordinates = append(upCoordinates, column+strconv.Itoa(tempRowInt-i))
		}

		tempCoordinateMap[tempMapKey] = upCoordinates
		tempMapKey++
	}
	// For test
	//fmt.Println("upCoordinates: ", upCoordinates)

	// for down. column is unchanged, row increased size
	downCoordinates := make([]string, 0)
	tempRowInt, _ = strconv.Atoi(row)
	if tempRowInt+size <= 11 {
		for i := 0; i < size; i++ {
			downCoordinates = append(downCoordinates, column+strconv.Itoa(tempRowInt+i))
		}

		tempCoordinateMap[tempMapKey] = downCoordinates
		tempMapKey++
	}
	// For test
	//fmt.Println("downCoordinates: ", downCoordinates)

	// for right. row is unchanged, column increase size
	rightCoordinates := make([]string, 0)
	if columnIndex+size <= len(global.Columns) {
		for i := 0; i < size; i++ {
			rightCoordinates = append(rightCoordinates, global.Columns[columnIndex+i]+row)
		}

		tempCoordinateMap[tempMapKey] = rightCoordinates
		tempMapKey++
	}
	// For test
	//fmt.Println("rightCoordinates: ", rightCoordinates)

	// for left. row is unchanged, column reduce size
	leftCoordinates := make([]string, 0)
	if columnIndex-size >= 1 {
		for i := 0; i < size; i++ {
			leftCoordinates = append(leftCoordinates, global.Columns[columnIndex-i]+row)
		}

		tempCoordinateMap[tempMapKey] = leftCoordinates
	}
	// For test
	//fmt.Println("leftCoordinates: ", leftCoordinates)

	// For test
	//fmt.Println("tempCoordinateMap: ", tempCoordinateMap)

	coordinates = tempCoordinateMap[rand.Intn(len(tempCoordinateMap))]

	return coordinates
}

// VerifyCoordinates is used to check if coordinates overlap
func VerifyCoordinates(player map[int][]string, coordinates []string) bool {

	if len(player) <= 0 {
		return false
	}

	for _, value := range coordinates {
		for _, v := range player {
			for _, item := range v {
				if item == value {
					return true
				}
			}
		}
	}

	//
	//for _, value := range coordinates {
	//	if _, ok := player[value]; !ok {
	//
	//	}
	//}1

	return false
}

// []string == [c1,d1,......J10]

// map[string]{}
// map[A2]{} map[A3]{}

//key
