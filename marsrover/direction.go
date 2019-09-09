package main
type Direction interface {
	leftMove() Direction
	rightMove() Direction
	moveForward(pos Rover) (int,int)
	toString() string
}
type direction struct {
	W Direction
	E Direction
	N Direction
	S Direction
}
var (
	Enum  = &direction{
		W: West{},
		E: East{},
		N: North{},
		S: South{},
	}
)
