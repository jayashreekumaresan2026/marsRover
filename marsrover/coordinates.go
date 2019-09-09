package main

type Coordinates struct {
	XPosition int `json:"xPosition"`
	YPosition int `json:"yPosition"`
}

func (Coordinates) coordinatesValues(XPosition int, YPosition int) Coordinates {
	return Coordinates{XPosition, YPosition}
}
