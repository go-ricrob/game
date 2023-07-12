// Package board provides the game board definition and methods.
package board

import (
	"encoding/json"
	"fmt"
	"strings"

	. "github.com/go-ricrob/game/types"
)

const (
	numBoardFields = 16
	numTileFields  = numBoardFields / 2
)

// TilePosition defines the position of a tile on the board.
type TilePosition byte

// TilePosition constants.
const (
	TopLeft TilePosition = iota
	TopRight
	BottomLeft
	BottomRight
)

var tilePositionStrs = []string{"topleft", "topright", "bottomleft", "bottomright"}

func (p TilePosition) String() string {
	if int(p) >= len(tilePositionStrs) {
		return fmt.Sprintf("invalid tile position %d", p)
	}
	return tilePositionStrs[p]
}

// Direction is the type of the 'compass direction' of a wall.
type Direction byte

// Direction bitmap constants. Use bitmap to define field walls in one byte.
const (
	North Direction = 1 << iota
	East
	South
	West
)

var directionStrs = map[Direction]string{North: "north", East: "east", South: "south", West: "west"}

func (d Direction) String() string {
	dirs := []string{}
	for dir, s := range directionStrs {
		if dir&d != 0 {
			dirs = append(dirs, s)
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(dirs, "|"))
}

// Directions is the set of valid directions.
var Directions = []Direction{North, East, South, West}

// Symbol is the type of a symbol.
type Symbol byte

// Symbol constants.
const (
	NoSymbol = iota
	Pyramid
	Star
	Moon
	Saturn
	Cosmic
)

var symbolStrs = []string{"", "Pyramid", "Star", "Moon", "Saturn", "Cosmic"}

func (s Symbol) String() string {
	if int(s) >= len(symbolStrs) {
		return fmt.Sprintf("invalid symbol %d", s)
	}
	return symbolStrs[s]
}

// Symbols is the set of valid symbols.
var Symbols = []Symbol{Pyramid, Star, Moon, Saturn, Cosmic}

/*
rotate tile clockwise
*/
type transform func(x int, y int) (int, int)

var posRotations = map[TilePosition]transform{
	TopLeft:     func(x, y int) (int, int) { return (numTileFields - 1) - y, x },                       // rotate 270 degree
	TopRight:    func(x, y int) (int, int) { return x, y },                                             // rotate 0 degree
	BottomRight: func(x, y int) (int, int) { return y, (numTileFields - 1) - x },                       // rotate 90 degree
	BottomLeft:  func(x, y int) (int, int) { return (numTileFields - 1) - x, (numTileFields - 1) - y }, // rotate 180 degree
}

var posShifts = map[TilePosition]transform{
	TopLeft:     func(x, y int) (int, int) { return x, numTileFields + y },                 // shift up
	TopRight:    func(x, y int) (int, int) { return numTileFields + x, numTileFields + y }, // shift up and right
	BottomRight: func(x, y int) (int, int) { return numTileFields + x, y },                 // shift right
	BottomLeft:  func(x, y int) (int, int) { return x, y },                                 // no shift
}

func rotateWalls(d Direction, k int) Direction {
	r := d << k     // shift k bits
	r |= r >> 4     // add overlow / rotate
	return r & 0x0f //clear high bits
}

/*
rotate walls clockwise
*/
type rotate func(d Direction) Direction

var wallRotations = map[TilePosition]rotate{
	TopLeft:     func(d Direction) Direction { return rotateWalls(d, 3) }, // rotate 270 degree- right bit shift 3
	TopRight:    func(d Direction) Direction { return d },                 // rotate 0 degree - no shift
	BottomRight: func(d Direction) Direction { return rotateWalls(d, 1) }, // rotate 90 degree - right bit shift 1
	BottomLeft:  func(d Direction) Direction { return rotateWalls(d, 2) }, // rotate 180 degree - right bit shift 2
}

// Field is the type representing a field of a board.
type Field struct {
	walls   Direction
	symbol  Symbol
	color   Color
	targets map[Direction]Coordinate
}

func (f *Field) String() string {
	return fmt.Sprintf("{%s %s %s %v}", f.walls, f.symbol, f.color, f.targets)
}

// MarshalJSON implements the json.Marshaler interface.
// Use own marshaler to avoid public fields in struct Field.
func (f *Field) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Walls  Direction `json:"walls,omitempty"`
		Symbol Symbol    `json:"symbol,omitempty"`
		Color  Color     `json:"color,omitempty"`
	}{
		Walls:  f.walls,
		Symbol: f.symbol,
		Color:  f.color,
	})
}

// Symbol returns the symbol of a field.
func (f *Field) Symbol() Symbol { return f.symbol }

// Color returns the color of a field.
func (f *Field) Color() Color { return f.color }

func (f *Field) hasWall(d Direction) bool { return f.walls&d != 0 }

func (f *Field) addWall(dirs ...Direction) {
	for _, d := range dirs {
		f.walls |= d
	}
}

type moveFn func(robots map[Color]Coordinate, color Color) (x, y int, ok bool)

// Board is the type representing a board.
type Board struct {
	fields []*Field
	Moves  map[Direction]moveFn
}

func fieldIdx(x, y int) int             { return y*numBoardFields + x }
func fieldCoord(idx int) (x int, y int) { return idx % numBoardFields, idx / numBoardFields }

func (b *Board) calculateTarget(x0, y0 int, direction Direction) (int, int, bool) {
	x, y := x0, y0

	switch direction {
	default:
		panic(fmt.Sprintf("invalid direction %s", direction))
	case East:
		for b.Field(x, y0).walls&East == 0 {
			x++
		}
	case West:
		for b.Field(x, y0).walls&West == 0 {
			x--
		}
	case North:
		for b.Field(x0, y).walls&North == 0 {
			y++
		}
	case South:
		for b.Field(x0, y).walls&South == 0 {
			y--
		}
	}

	if x == x0 && y == y0 {
		return x0, y0, false
	}
	return x, y, true
}

// New creates a new board instance. Parameter tiles needs to be valid - if not NewBoard will panic.
func New(tileIDMap map[TilePosition]string) *Board {
	numField := numBoardFields * numBoardFields
	b := &Board{fields: make([]*Field, numField)}
	b.Moves = map[Direction]moveFn{
		North: b.moveNorth,
		East:  b.moveEast,
		South: b.moveSouth,
		West:  b.moveWest,
	}
	// init fields
	for i := 0; i < numField; i++ {
		b.fields[i] = &Field{targets: map[Direction]Coordinate{}}
	}

	// set tile fields
	for p, id := range tileIDMap {
		fields, ok := tiles[id]
		if !ok {
			panic(fmt.Errorf("invalid tileID: %v", id))
		}
		for c, f := range fields {
			x, y := posShifts[p](posRotations[p](c.X, c.Y)) // rotate and shift
			field := b.fields[fieldIdx(x, y)]
			field.walls = wallRotations[p](f.walls)
			field.symbol = f.symbol
			field.color = f.color
		}
	}

	// set outer walls
	for x := 0; x < numBoardFields; x++ {
		b.fields[fieldIdx(x, numBoardFields-1)].addWall(North) // top border
		b.fields[fieldIdx(x, 0)].addWall(South)                // bottom border
	}
	for y := 0; y < numBoardFields; y++ {
		b.fields[fieldIdx(0, y)].addWall(West)                // left border
		b.fields[fieldIdx(numBoardFields-1, y)].addWall(East) // right border
	}

	// set center walls
	b.fields[fieldIdx(numTileFields-1, numTileFields-1)].addWall(West, South) // bottom left center field
	b.fields[fieldIdx(numTileFields-1, numTileFields)].addWall(West, North)   // top left center field
	b.fields[fieldIdx(numTileFields, numTileFields-1)].addWall(East, South)   // bottom right center field
	b.fields[fieldIdx(numTileFields, numTileFields)].addWall(East, North)     // top right center field

	// finally: set neighbor walls
	for x := 0; x < numBoardFields; x++ {
		for y := 0; y < numBoardFields; y++ {
			idx := fieldIdx(x, y)
			if !b.fields[idx].hasWall(North) && y < numBoardFields-1 && b.fields[fieldIdx(x, y+1)].hasWall(South) { // north neighbor
				b.fields[idx].addWall(North)
			}
			if !b.fields[idx].hasWall(East) && x < numBoardFields-1 && b.fields[fieldIdx(x+1, y)].hasWall(West) { // east neighbor
				b.fields[idx].addWall(East)
			}
			if !b.fields[idx].hasWall(South) && y > 0 && b.fields[fieldIdx(x, y-1)].hasWall(North) { // south neighbor
				b.fields[idx].addWall(South)
			}
			if !b.fields[idx].hasWall(West) && x > 0 && b.fields[fieldIdx(x-1, y)].hasWall(East) { // west neighbor
				b.fields[idx].addWall(West)
			}
		}
	}

	// calculate routes
	for x0 := 0; x0 < numBoardFields; x0++ {
		for y0 := 0; y0 < numBoardFields; y0++ {
			field := b.Field(x0, y0)
			for _, direction := range Directions {
				if x, y, ok := b.calculateTarget(x0, y0, direction); ok {
					field.targets[direction] = Coordinate{X: x, Y: y}
				}
			}
		}
	}

	return b
}

// IsValidCoordinate returns true if the coordinate represents a valid board field, false otherwise.
func (b *Board) IsValidCoordinate(x, y int) bool {
	// in range ?
	if x < 0 || x >= numBoardFields || y < 0 || y >= numBoardFields {
		return false
	}
	// center field ?
	c1 := numTileFields - 1
	c2 := numTileFields
	if x >= c1 && x <= c2 && y >= c1 && y <= c2 {
		return false
	}
	return true
}

// Field returns the board field at position x,y.
func (b *Board) Field(x, y int) *Field { return b.fields[fieldIdx(x, y)] }

// TargetCoordinate returns the coordinate of the target.
func (b *Board) TargetCoordinate(symbol Symbol, color Color) (x, y int) {
	for idx, field := range b.fields {
		if field.symbol == symbol && field.color == color {
			return fieldCoord(idx)
		}
	}
	panic(fmt.Errorf("invalid target: symbol %s color %s", symbol, color)) // should never happen
}

// Move returns the coordinate of a target field moving a robot in one of the directions.

func (b *Board) moveNorth(robots map[Color]Coordinate, color Color) (x, y int, ok bool) {
	// handle redirects here
	// field needs to handle redirection as a boarder
	// routes needs to deliver redirect field coords
	robot := robots[color]
	field := b.Field(robot.X, robot.Y)
	target, ok := field.targets[North]
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

func (b *Board) moveEast(robots map[Color]Coordinate, color Color) (x, y int, ok bool) {
	// handle redirects here
	// field needs to handle redirection as a boarder
	// routes needs to deliver redirect field coords
	robot := robots[color]
	field := b.Field(robot.X, robot.Y)
	target, ok := field.targets[East]
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

func (b *Board) moveSouth(robots map[Color]Coordinate, color Color) (x, y int, ok bool) {
	// handle redirects here
	// field needs to handle redirection as a boarder
	// routes needs to deliver redirect field coords
	robot := robots[color]
	field := b.Field(robot.X, robot.Y)
	target, ok := field.targets[South]
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

func (b *Board) moveWest(robots map[Color]Coordinate, color Color) (x, y int, ok bool) {
	// handle redirects here
	// field needs to handle redirection as a boarder
	// routes needs to deliver redirect field coords
	robot := robots[color]
	field := b.Field(robot.X, robot.Y)
	target, ok := field.targets[West]
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
