package board

import (
	"fmt"
	"strings"
)

// Wall defines b oard field walls.
type Wall byte

// Wall constants.
const (
	NorthWall Wall = 1 << iota
	EastWall
	SouthWall
	WestWall
)

var wallStrs = map[Wall]string{
	NorthWall: "northWall",
	EastWall:  "eastWall",
	SouthWall: "southWall",
	WestWall:  "westWall"}

func (w Wall) String() string {
	walls := []string{}
	for wall, s := range wallStrs {
		if wall&w != 0 {
			walls = append(walls, s)
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(walls, "|"))
}
