package main

import "fmt"

func main() {
	rect := Rectangle{
		Width:  10,
		Length: 5,
	}

	area := rect.Area()

	fmt.Println("Area:", area)
}

type Rectangle struct {
	Width  int
	Length int
}

func (this Rectangle) Area() int  {
	return this.Width * this.Length
}
