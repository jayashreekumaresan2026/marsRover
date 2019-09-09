package main

import "fmt"

type Rover struct {
	XPosition int `json:"xPosition"`
	YPosition int `json:"yPosition"`
	Direction Direction  `json:"direction"`
	MaxGridPlanet MaxGridPlanet `json:"maxGridPlanets"`
}

func (pos Rover) toString() string {
	return fmt.Sprint(pos.XPosition, " ",pos.YPosition, " ", pos.Direction.toString())
}
func createRover() Rover {
	var pos = Rover{Rover{}.XPosition, Rover{}.YPosition, Rover{}.Direction, MaxGridPlanet{}}

	return pos
}










