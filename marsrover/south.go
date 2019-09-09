package main

type South struct {
}

func (South) rightMove() Direction {
	return Enum.W
}
func (South) leftMove() Direction {
	return Enum.E
}
func (South) moveForward(coordinate Coordinates) Coordinates {
	return Coordinates{XPosition:coordinate.XPosition,YPosition:coordinate.YPosition-1}
}
func (South) toString() string {
	return "South"
}
