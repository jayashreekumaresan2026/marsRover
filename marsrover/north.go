package main

type North struct {
}

func (North) rightMove() Direction {
	return Enum.N
}
func (North) leftMove() Direction {
	return Enum.S
}
func (North) moveForward(pos Rover) (int, int) {
	return pos.XPosition, pos.YPosition + 1
}
func (North) toString() string {
	return "North"
}
