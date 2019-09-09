package main

type East struct {
}

func (East) rightMove() Direction {
	return Enum.N
}
func (East) leftMove() Direction {
	return Enum.S
}
func (East) moveForward(pos Rover) (int, int) {
	return pos.XPosition + 1, pos.YPosition
}
func (East) toString() string {
	return "East"
}
