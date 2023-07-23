package types

import "fmt"

// Dir is the type of the 'compass direction' movement.
type Dir byte

// Dir constants.
const (
	North = iota
	East
	South
	West
	NumDir
)

var dirStrs = []string{"north", "east", "south", "west"}

func (d Dir) String() string {
	if int(d) >= len(dirStrs) {
		return fmt.Sprintf("invalid direction %d", d)
	}
	return dirStrs[d]
}
