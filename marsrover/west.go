package main

type West struct {
}
func (West) rightMove() Direction {
	return Enum.N
}
func (West) leftMove() Direction {
	return Enum.S
}
func (West) moveForward(pos Rover) (int, int) {
	return pos.XPosition - 1, pos.YPosition
}
func (West) toString() string {
	return "West"
}
