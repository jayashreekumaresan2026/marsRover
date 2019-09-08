package main
type Direction interface {
	moveLeft() Direction
	moveRight() Direction
	moveForward() Direction
}

