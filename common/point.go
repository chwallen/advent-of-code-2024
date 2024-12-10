package common

import (
	"fmt"
)

// A Point represents a point in a 2D space.
type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{x, y}
}

// Add returns a new point with one step in direction d.
func (c Point) Add(d Direction) Point {
	return Point{x: c.x + d.x, y: c.y + d.y}
}

// Neighbors returns all neighboring points.
func (c Point) Neighbors() []Point {
	return []Point{
		c.Add(Left),
		c.Add(Right),
		c.Add(Up),
		c.Add(Down),
	}
}

func (c Point) X() int {
	return c.x
}

func (c Point) Y() int {
	return c.y
}

func (c Point) XY() (x, y int) {
	return c.x, c.y
}

func (c Point) IsWithinBounds(minX, minY, maxX, maxY int) bool {
	return minX <= c.x && c.x < maxX && minY <= c.y && c.y < maxY
}

type Direction Point

var Up = Direction{x: 0, y: -1}
var UpRight = Direction{x: 1, y: -1}
var Right = Direction{x: 1, y: 0}
var DownRight = Direction{x: 1, y: 1}
var Down = Direction{x: 0, y: 1}
var DownLeft = Direction{x: -1, y: 1}
var Left = Direction{x: -1, y: 0}
var UpLeft = Direction{x: -1, y: -1}

// TurnRight returns a new point which is turned 90 degrees right.
func (d Direction) TurnRight() Direction {
	return Direction{-d.y, d.x}
}

// TurnLeft returns the point which is turned 90 degrees left.
func (d Direction) TurnLeft() Direction {
	return Direction{d.y, -d.x}
}

// GetCardinalIndex gets the int which represents the direction.
// Only works for the directions Up, Right, Down, and Left.
func (d Direction) GetCardinalIndex() int {
	switch d {
	case Up:
		return 0
	case Right:
		return 1
	case Down:
		return 2
	case Left:
		return 3
	default:
		panic(fmt.Errorf("invalid direction %v", d))
	}
}
