package main

type West struct {
}
func (West) rightMove() Direction {
	return Enum.N
}
func (West) leftMove() Direction {
	return Enum.S
}
func (West) moveForward(coordinate Coordinates) Coordinates {
	return Coordinates{XPosition:coordinate.XPosition-1,YPosition:coordinate.YPosition}
}
func (West) toString() string {
	return "West"
}
