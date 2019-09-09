package main

import "fmt"

type Rover struct {
	Coordinates Coordinates `json:"Coordinates"`
	Direction Direction  `json:"direction"`
	MaxGridPlanet MaxGridPlanet `json:"maxGridPlanets"`
}

func createRover() Rover {
	var pos = Rover{Rover{}.Coordinates, Rover{}.Direction, MaxGridPlanet{}}
	return pos
}
func (pos *Rover) toString() string{
	return fmt.Sprint(pos.Coordinates.XPosition,pos.Coordinates.YPosition," ", pos.Direction.toString())
}










