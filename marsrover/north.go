package main

type North struct {
}

func (North) rightMove() Direction {
	return Enum.E
}
func (North) leftMove() Direction {
	return Enum.W
}
func (North) moveForward(coordinate Coordinates) Coordinates {
	return Coordinates{XPosition:coordinate.XPosition,YPosition:coordinate.YPosition+1}
}
func (North) toString() string {
	return "North"
}
