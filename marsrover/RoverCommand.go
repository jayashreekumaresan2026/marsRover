package main

import (
	"strings"
)

func (*Rover) start(command string) (string, error) {
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

func (pos *Rover) moveLeft() {
	rover.Direction = rover.Direction.leftMove()
}
func (pos *Rover) moveRight() {
	rover.Direction = rover.Direction.rightMove()
}
func (pos *Rover) move() {
  pos.Direction.moveForward(Rover{})
}
