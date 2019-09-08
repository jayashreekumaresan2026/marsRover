package main

import "fmt"

type Rover struct {
	XPosition int `json:"xPosition"`
	YPosition int `json:"yPosition"`
	Direction string  `json:"direction"`
	MaxGridPlanet MaxGridPlanet `json:"maxGridPlanets"`
}
func (pos *Rover) getDirection() string {
	return pos.Direction
}

func (pos *Rover) setDirection(direction string) {
	pos.Direction = direction
}

func (pos Rover) toString() string {
	return fmt.Sprint(pos.XPosition, " ",pos.YPosition, " ", pos.Direction)
}


func (pos *Rover) move() {
	pos.moveForward()
}









