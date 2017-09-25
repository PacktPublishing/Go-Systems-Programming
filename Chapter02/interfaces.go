package main

import (
	"fmt"
)

type coordinates interface {
	xaxis() int
	yaxis() int
}

type point2D struct {
	X int
	Y int
}

func (s point2D) xaxis() int {
	return s.X
}

func (s point2D) yaxis() int {
	return s.Y
}

func findCoordinates(a coordinates) {
	fmt.Println("X:", a.xaxis(), "Y:", a.yaxis())
}

type coordinate int

func (s coordinate) xaxis() int {
	return int(s)
}

func (s coordinate) yaxis() int {
	return 0
}

func main() {

	x := point2D{X: -1, Y: 12}
	fmt.Println(x)
	findCoordinates(x)

	y := coordinate(10)
	findCoordinates(y)
}
