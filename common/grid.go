package common

// A Grid represents a static 2D slice of items.
type Grid[V any] struct {
	data  [][]V
	sizeY int
	sizeX int
}

// Allocate2DSlice efficiently allocates one large slice instead of multiple
// smaller ones which is faster to allocate and more cache efficient.
func Allocate2DSlice[T any](maxX, maxY int) [][]T {
	rows := make([][]T, maxY)
	columns := make([]T, maxY*maxX)
	for i := range rows {
		rows[i], columns = columns[:maxX], columns[maxX:]
	}
	return rows
}

func NewGrid[V any](maxX, maxY int) Grid[V] {
	data := Allocate2DSlice[V](maxX, maxY)
	return Grid[V]{data: data, sizeY: maxY, sizeX: maxX}
}

// Clone creates a shallow copy of g.
func (g Grid[V]) Clone() Grid[V] {
	clone := NewGrid[V](g.sizeX, g.sizeY)
	clone.CopyFrom(g)
	return clone
}

// Get retrieves the data at x,y. Panics if out of bounds.
func (g Grid[V]) Get(x, y int) V {
	return g.data[y][x]
}

// GetPoint retrieves the data at the point's coordinates.
// Panics if out of bounds.
func (g Grid[V]) GetPoint(p Point) V {
	return g.Get(p.X(), p.Y())
}

func (g Grid[V]) IsWithinBounds(x, y int) bool {
	return 0 <= x && x < g.sizeX && 0 <= y && y < g.sizeY
}

func (g Grid[V]) IsPointWithinBounds(p Point) bool {
	return g.IsWithinBounds(p.X(), p.Y())
}

func (g Grid[V]) Size() (sizeX, sizeY int) {
	return g.sizeX, g.sizeY
}

func (g Grid[V]) Set(x, y int, v V) {
	g.data[y][x] = v
}

// CopyFrom shallowly copies all data from src into g.
// Panics if g or src have no rows or if g and src differ in size.
func (g Grid[V]) CopyFrom(src Grid[V]) {
	n := cap(g.data[0])
	if n != cap(src.data[0]) {
		panic("cannot copy src into g as dimensions differ")
	}
	// Since all rows are actually one contiguous memory space, we can copy all
	// data in a single call.
	copy(g.data[0][:n], src.data[0][:n])
}
