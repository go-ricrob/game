package types

import (
	"fmt"
)

// Coordinate defines a two dimensional coordinate.
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c Coordinate) String() string { return fmt.Sprintf("%d,%d", c.X, c.Y) }
