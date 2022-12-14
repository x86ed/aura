package util

// Renderable is the interface all objects drawable on the screen must use
type Renderable interface {
	// Draw draws the method the first Coord value represents the offset.
	// The second represents
	Draw(...Coord)
	Hide(...Coord)
	Interrupt() error
}

// Coord is a struct which represents
type Coord struct {
	X, Y int
}

// RenderObj is the base object for objects renderable on the screen
type RenderObj struct {
	Dims     Coord
	Offset   Coord
	IsOffset bool
	Label    string
}
