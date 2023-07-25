package board

import (
	"fmt"
)

// Color is the type of a color.
type Color byte

// Color bitmap constants.
const (
	_ Color = iota // no color
	Yellow
	Red
	Green
	Blue
	Silver
)

var colorStrs = []string{"", "yellow", "red", "green", "blue", "silver"}

func (c Color) String() string {
	if int(c) >= len(colorStrs) {
		return fmt.Sprintf("invalid symbol %d", c)
	}
	return colorStrs[c]
}
