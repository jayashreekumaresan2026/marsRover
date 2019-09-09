package main

type East struct {
}

func (East) rightMove() Direction {
	return Enum.S
}
func (East) leftMove() Direction {
	return Enum.N
}
func (East) moveForward(coordinate Coordinates) Coordinates {
	return Coordinates{XPosition:coordinate.XPosition+1,YPosition:coordinate.YPosition}
}
func (East) toString() string {
	return "East"
}
