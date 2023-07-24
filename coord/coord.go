// Package coord defines types and functions for coordinate conversions.
package coord

type XY struct {
	X, Y int
}

// Ctob converts a x,y coordinate with x,y < 16 to a coordinate byte.
func Ctob(x, y int) byte { return byte(x)<<4 | byte(y) }

// Btoc converts a coordinate byte to x,y coordinates.
func Btoc(b byte) (x, y int) { return int(b >> 4), int(b & 0x0f) }

// X returns the x coordinate of a coordinate byte.
func X(b byte) int { return int(b >> 4) }

// Y returns the y coordinate of a coordinate byte.
func Y(b byte) int { return int(b & 0x0f) }
