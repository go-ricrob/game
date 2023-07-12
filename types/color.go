package types

import (
	"fmt"
	"strings"
)

// Color is the type of a color.
type Color byte

// Color bitmap constants.
const (
	Yellow Color = 1 << iota
	Red
	Green
	Blue
	Silver

	CosmicColor = Yellow | Red | Green | Blue | Silver
)

// IsValid checks if a color is valid.
func (c Color) IsValid() bool {
	return c == Yellow || c == Red || c == Green || c == Blue || c == Silver
}

var colorStrs = map[Color]string{Yellow: "yellow", Red: "red", Green: "green", Blue: "blue", Silver: "silver"}

func (c Color) String() string {
	colors := []string{}
	for color, s := range colorStrs {
		if color&c != 0 {
			colors = append(colors, s)
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(colors, "|"))
}

// Colors is the set of valid colors.
var Colors = []Color{Yellow, Red, Green, Blue, Silver}
