package board

import "fmt"

// Tile defines the position of a tile on the board.
type Tile byte

// Tile constants.
const (
	TopLeft Tile = iota
	TopRight
	BottomLeft
	BottomRight
	NumTile
)

var tileStrs = []string{"topleft", "topright", "bottomleft", "bottomright"}

func (p Tile) String() string {
	if int(p) >= len(tileStrs) {
		return fmt.Sprintf("invalid tile position %d", p)
	}
	return tileStrs[p]
}
