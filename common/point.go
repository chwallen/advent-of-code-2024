package common

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

type Direction Point

var Up = Direction{x: 0, y: -1}
var UpRight = Direction{x: 1, y: -1}
var Right = Direction{x: 1, y: 0}
var DownRight = Direction{x: 1, y: 1}
var Down = Direction{x: 0, y: 1}
var DownLeft = Direction{x: -1, y: 1}
var Left = Direction{x: -1, y: 0}
var UpLeft = Direction{x: -1, y: -1}
