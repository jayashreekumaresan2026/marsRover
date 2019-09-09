package main

type South struct {
}

func (South) rightMove() Direction {
	return Enum.N
}
func (South) leftMove() Direction {
	return Enum.S
}
func (South) moveForward(pos Rover) (int, int) {
	return pos.XPosition - 1, pos.YPosition
}
func (South) toString() string {
	return "South"
}
