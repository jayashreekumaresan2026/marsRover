package main

type MaxGridPlanet struct {
	XMax int `json:"XMax"`
	YMax int `json:"YMax"`
}

var _ MaxGridPlanet

func (max *MaxGridPlanet) isLesser(coordinates Coordinates) bool {
	return (coordinates.XPosition <= max.XMax) && (coordinates.YPosition <= max.YMax)
}

