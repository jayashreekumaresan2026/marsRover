package main

import (
	"strconv"
	"strings"
)

func getRoverPositionAndDirection() string {
	return "12N"
}
func(*Rover) start(command string)(string,error) {
	pos := createRover()
	roverCommand := strings.Split(command, "")
	for i := 0; i < len(command); i++ {
		switch roverCommand[i] {
		case "L":
			pos.moveLeft()
		case "R":
			pos.moveRight()
		case "M":
			pos.move()
		}
	}
	return rover.toString(), nil
}

func createRover() Rover {
	var maxGrid MaxGridPlanet
	maxGrid.XMax, maxGrid.YMax = 5, 5
	roverPosition := getRoverPositionAndDirection()
	splitInstruction := strings.Split(roverPosition, "")
	xPosition, _ := strconv.Atoi(splitInstruction[0])
	yPosition, _ := strconv.Atoi(splitInstruction[1])
	direction := splitInstruction[2]
	var pos = Rover{xPosition, yPosition, direction, MaxGridPlanet{}}
	return pos
}
