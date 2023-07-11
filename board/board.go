// Package board provides the game board definition and methods.
package board

import (
	"encoding/json"
	"fmt"

	"github.com/go-ricrob/game/types"
)

/*
rotate tile clockwise
*/
type transform func(x int, y int) (int, int)

var posRotations = map[types.TilePosition]transform{
	types.TopLeft:     func(x, y int) (int, int) { return (types.MaxTileFields - 1) - y, x },                             // rotate 270 degree
	types.TopRight:    func(x, y int) (int, int) { return x, y },                                                         // rotate 0 degree
	types.BottomRight: func(x, y int) (int, int) { return y, (types.MaxTileFields - 1) - x },                             // rotate 90 degree
	types.BottomLeft:  func(x, y int) (int, int) { return (types.MaxTileFields - 1) - x, (types.MaxTileFields - 1) - y }, // rotate 180 degree
}

var posShifts = map[types.TilePosition]transform{
	types.TopLeft:     func(x, y int) (int, int) { return x, types.MaxTileFields + y },                       // shift up
	types.TopRight:    func(x, y int) (int, int) { return types.MaxTileFields + x, types.MaxTileFields + y }, // shift up and right
	types.BottomRight: func(x, y int) (int, int) { return types.MaxTileFields + x, y },                       // shift right
	types.BottomLeft:  func(x, y int) (int, int) { return x, y },                                             // no shift
}

func rotateWalls(d types.Direction, k int) types.Direction {
	r := d << k     // shift k bits
	r |= r >> 4     // add overlow / rotate
	return r & 0x0f //clear high bits
}

/*
rotate walls clockwise
*/
type rotate func(d types.Direction) types.Direction

var wallRotations = map[types.TilePosition]rotate{
	types.TopLeft:     func(d types.Direction) types.Direction { return rotateWalls(d, 3) }, // rotate 270 degree- right bit shift 3
	types.TopRight:    func(d types.Direction) types.Direction { return d },                 // rotate 0 degree - no shift
	types.BottomRight: func(d types.Direction) types.Direction { return rotateWalls(d, 1) }, // rotate 90 degree - right bit shift 1
	types.BottomLeft:  func(d types.Direction) types.Direction { return rotateWalls(d, 2) }, // rotate 180 degree - right bit shift 2
}

// Field is the type representing a field of a board.
type Field struct {
	walls   types.Direction
	target  types.Target
	targets map[types.Direction]types.Coordinate
}

func (f *Field) String() string {
	return fmt.Sprintf("{%s %s %v}", f.walls, f.target, f.targets)
}

// MarshalJSON implements the json.Marshaler interface.
// Use own marshaler to avoid public fields in struct Field.
func (f *Field) MarshalJSON() ([]byte, error) {
	targetFn := func(t types.Target) *types.Target {
		if t.Symbol == 0 && t.Color == 0 {
			return nil
		}
		return &t
	}

	return json.Marshal(struct {
		Walls types.Direction `json:"walls,omitempty"`
		// need to use pointer here, so that omitempy does work for target
		Target *types.Target `json:"target,omitempty"`
	}{
		Walls:  f.walls,
		Target: targetFn(f.target),
	})
}

func (f *Field) hasWall(d types.Direction) bool { return f.walls&d != 0 }

func (f *Field) addWall(dirs ...types.Direction) {
	for _, d := range dirs {
		f.walls |= d
	}
}

type moveFn func(robots map[types.Color]types.Coordinate, color types.Color) (x, y int, ok bool)

// Board is the type representing a board.
type Board struct {
	fields []*Field
	Moves  map[types.Direction]moveFn
}

func fieldIdx(x, y int) int { return y*types.MaxBoardFields + x }
func fieldCoord(idx int) (x int, y int) {
	return idx % types.MaxBoardFields, idx / types.MaxBoardFields
}

func (b *Board) calculateTarget(x0, y0 int, direction types.Direction) (int, int, bool) {
	x, y := x0, y0

	switch direction {
	default:
		panic(fmt.Sprintf("invalid direction %s", direction))
	case types.East:
		for b.Field(x, y0).walls&types.East == 0 {
			x++
		}
	case types.West:
		for b.Field(x, y0).walls&types.West == 0 {
			x--
		}
	case types.North:
		for b.Field(x0, y).walls&types.North == 0 {
			y++
		}
	case types.South:
		for b.Field(x0, y).walls&types.South == 0 {
			y--
		}
	}

	if x == x0 && y == y0 {
		return x0, y0, false
	}
	return x, y, true
}

// NewBoard creates a new board instance. Parameter tiles needs to be valid - if not NewBoard will panic.
func NewBoard(tiles map[types.TilePosition]types.TileID) *Board {
	numField := types.MaxBoardFields * types.MaxBoardFields
	b := &Board{fields: make([]*Field, numField)}
	b.Moves = map[types.Direction]moveFn{
		types.North: b.moveNorth,
		types.East:  b.moveEast,
		types.South: b.moveSouth,
		types.West:  b.moveWest,
	}
	// init fields
	for i := 0; i < numField; i++ {
		b.fields[i] = &Field{targets: map[types.Direction]types.Coordinate{}}
	}

	// set tile fields
	for p, id := range tiles {
		fields, ok := tileFields[id]
		if !ok {
			panic(fmt.Errorf("invalid tileID: %v", id))
		}
		for c, f := range fields {
			x, y := posShifts[p](posRotations[p](c.X, c.Y)) // rotate and shift
			field := b.fields[fieldIdx(x, y)]
			field.walls = wallRotations[p](f.walls)
			field.target = f.target
		}
	}

	// set outer walls
	for x := 0; x < types.MaxBoardFields; x++ {
		b.fields[fieldIdx(x, types.MaxBoardFields-1)].addWall(types.North) // top border
		b.fields[fieldIdx(x, 0)].addWall(types.South)                      // bottom border
	}
	for y := 0; y < types.MaxBoardFields; y++ {
		b.fields[fieldIdx(0, y)].addWall(types.West)                      // left border
		b.fields[fieldIdx(types.MaxBoardFields-1, y)].addWall(types.East) // right border
	}

	// set center walls
	b.fields[fieldIdx(types.MaxTileFields-1, types.MaxTileFields-1)].addWall(types.West, types.South) // bottom left center field
	b.fields[fieldIdx(types.MaxTileFields-1, types.MaxTileFields)].addWall(types.West, types.North)   // top left center field
	b.fields[fieldIdx(types.MaxTileFields, types.MaxTileFields-1)].addWall(types.East, types.South)   // bottom right center field
	b.fields[fieldIdx(types.MaxTileFields, types.MaxTileFields)].addWall(types.East, types.North)     // top right center field

	// finally: set neighbor walls
	for x := 0; x < types.MaxBoardFields; x++ {
		for y := 0; y < types.MaxBoardFields; y++ {
			idx := fieldIdx(x, y)
			if !b.fields[idx].hasWall(types.North) && y < types.MaxBoardFields-1 && b.fields[fieldIdx(x, y+1)].hasWall(types.South) { // north neighbor
				b.fields[idx].addWall(types.North)
			}
			if !b.fields[idx].hasWall(types.East) && x < types.MaxBoardFields-1 && b.fields[fieldIdx(x+1, y)].hasWall(types.West) { // east neighbor
				b.fields[idx].addWall(types.East)
			}
			if !b.fields[idx].hasWall(types.South) && y > 0 && b.fields[fieldIdx(x, y-1)].hasWall(types.North) { // south neighbor
				b.fields[idx].addWall(types.South)
			}
			if !b.fields[idx].hasWall(types.West) && x > 0 && b.fields[fieldIdx(x-1, y)].hasWall(types.East) { // west neighbor
				b.fields[idx].addWall(types.West)
			}
		}
	}

	// calculate routes
	for x0 := 0; x0 < types.MaxBoardFields; x0++ {
		for y0 := 0; y0 < types.MaxBoardFields; y0++ {
			field := b.Field(x0, y0)
			for _, direction := range types.Directions {
				if x, y, ok := b.calculateTarget(x0, y0, direction); ok {
					field.targets[direction] = types.Coordinate{X: x, Y: y}
				}
			}
		}
	}

	return b
}

// Field returns the board field at position x,y.
func (b *Board) Field(x, y int) *Field { return b.fields[fieldIdx(x, y)] }

// TargetCoordinate returns the coordinate of the target.
func (b *Board) TargetCoordinate(target types.Target) (x, y int) {
	for idx, field := range b.fields {
		if field.target == target {
			return fieldCoord(idx)
		}
	}
	panic(fmt.Errorf("invalid target %s", target)) // should never happen
}

// Move returns the coordinate of a target field moving a robot in one of the directions.

func (b *Board) moveNorth(robots map[types.Color]types.Coordinate, color types.Color) (x, y int, ok bool) {
	// handle redirects here
	// field needs to handle redirection as a boarder
	// routes needs to deliver redirect field coords
	robot := robots[color]
	field := b.Field(robot.X, robot.Y)
	target, ok := field.targets[types.North]
	if !ok {
		return robot.X, robot.Y, false
	}
	for otherColor, otherRobot := range robots {
		if otherColor != color && otherRobot.X == robot.X && otherRobot.Y > robot.Y && otherRobot.Y <= target.Y {
			target.Y = otherRobot.Y - 1
		}
	}
	if target.X == robot.X && target.Y == robot.Y {
		return robot.X, robot.Y, false
	}
	return target.X, target.Y, true
}

func (b *Board) moveEast(robots map[types.Color]types.Coordinate, color types.Color) (x, y int, ok bool) {
	// handle redirects here
	// field needs to handle redirection as a boarder
	// routes needs to deliver redirect field coords
	robot := robots[color]
	field := b.Field(robot.X, robot.Y)
	target, ok := field.targets[types.East]
	if !ok {
		return robot.X, robot.Y, false
	}
	for otherColor, otherRobot := range robots {
		if otherColor != color && otherRobot.Y == robot.Y && otherRobot.X > robot.X && otherRobot.X <= target.X {
			target.X = otherRobot.X - 1
		}
	}
	if target.X == robot.X && target.Y == robot.Y {
		return robot.X, robot.Y, false
	}
	return target.X, target.Y, true
}

func (b *Board) moveSouth(robots map[types.Color]types.Coordinate, color types.Color) (x, y int, ok bool) {
	// handle redirects here
	// field needs to handle redirection as a boarder
	// routes needs to deliver redirect field coords
	robot := robots[color]
	field := b.Field(robot.X, robot.Y)
	target, ok := field.targets[types.South]
	if !ok {
		return robot.X, robot.Y, false
	}
	for otherColor, otherRobot := range robots {
		if otherColor != color && otherRobot.X == robot.X && otherRobot.Y < robot.Y && otherRobot.Y >= target.Y {
			target.Y = otherRobot.Y + 1
		}
	}
	if target.X == robot.X && target.Y == robot.Y {
		return robot.X, robot.Y, false
	}
	return target.X, target.Y, true
}

func (b *Board) moveWest(robots map[types.Color]types.Coordinate, color types.Color) (x, y int, ok bool) {
	// handle redirects here
	// field needs to handle redirection as a boarder
	// routes needs to deliver redirect field coords
	robot := robots[color]
	field := b.Field(robot.X, robot.Y)
	target, ok := field.targets[types.West]
	if !ok {
		return robot.X, robot.Y, false
	}
	for otherColor, otherRobot := range robots {
		if otherColor != color && otherRobot.Y == robot.Y && otherRobot.X < robot.X && otherRobot.X >= target.X {
			target.X = otherRobot.X + 1
		}
	}
	if target.X == robot.X && target.Y == robot.Y {
		return robot.X, robot.Y, false
	}
	return target.X, target.Y, true
}

// MarshalJSON implements the json.Marshaler interface.
func (b *Board) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.fields)
}
