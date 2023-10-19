package main

type Position struct {
	X int
	Y int
}

func (this *Position) Move(dx, dy int)  {

	this.X+= dx
	this.Y+= dy
}

func main() {
	p := Position{X: 10, Y: 20}
	p.Move(5, -5)
}
