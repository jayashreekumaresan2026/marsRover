package main

func (Rover) start(command string) (string, error) {
	pos := createRover()
	for i := 0; i < len(command); i++ {
		switch command[i] {
		case 'L':
			pos.moveLeft()
		case 'R':
			pos.moveRight()
		case 'M':
			pos.move()
		}
	}
	return rover.toString(), nil
}

func (Rover) moveLeft() {
	rover.Direction = rover.Direction.leftMove()
}
func (Rover) moveRight() {
	rover.Direction = rover.Direction.rightMove()
}
func (Rover) move() {
	rover.Coordinates = rover.Direction.moveForward(rover.Coordinates)
}
