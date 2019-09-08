package main

const n string = "N"
const e string = "E"
const s string = "S"
const w string = "W"

func (pos *Rover) moveLeft() string {
	switch pos.Direction {
	case n:
		pos.Direction = w
		break
	case w:
		pos.Direction = s
		break
	case s:
		pos.Direction = e
		break
	case e:
		pos.Direction = n
		break
	}
	return pos.Direction
}

func (pos *Rover) moveRight() string{
	switch pos.Direction {
	case n:
		pos.Direction = e
		break
	case e:
		pos.Direction = s
		break
	case s:
		pos.Direction = w
		break
	case w:
		pos.Direction = n
		break
	}
	return pos.Direction
}

func (pos *Rover) moveForward() (int,int){
	switch pos.Direction {
	case "N":
		pos.YPosition = pos.YPosition + 1
	case "S":
		pos.YPosition = pos.YPosition - 1
	case "E":
		pos.XPosition = pos.XPosition + 1
	case "W":
		pos.XPosition = pos.XPosition - 1
	}
	return pos.XPosition,pos.YPosition
}
