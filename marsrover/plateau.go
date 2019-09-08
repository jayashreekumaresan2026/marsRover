package main

type MaxGridPlanet struct {
	XMax int `json:"XMax"`
	YMax int `json:"YMax"`
}

var maxGrid MaxGridPlanet

func (max *MaxGridPlanet) isLesser(pos *Rover) bool {
	return (pos.XPosition <= max.XMax) && (pos.YPosition <= max.YMax)
}

