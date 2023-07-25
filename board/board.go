// Package board provides the game board definition and methods.
package board

import (
	"fmt"

	"github.com/go-ricrob/game/coord"
)

const (
	numBoardField = 16
	numTileField  = numBoardField / 2
)

// NumField is the number of board fields.
const NumField = numBoardField * numBoardField

/*
rotate tile clockwise
*/
type transform func(x int, y int) (int, int)

var posRotations = [NumTile]transform{
	TopLeft:     func(x, y int) (int, int) { return (numTileField - 1) - y, x },                      // rotate 270 degree
	TopRight:    func(x, y int) (int, int) { return x, y },                                           // rotate 0 degree
	BottomRight: func(x, y int) (int, int) { return y, (numTileField - 1) - x },                      // rotate 90 degree
	BottomLeft:  func(x, y int) (int, int) { return (numTileField - 1) - x, (numTileField - 1) - y }, // rotate 180 degree
}

var posShifts = [NumTile]transform{
	TopLeft:     func(x, y int) (int, int) { return x, numTileField + y },                // shift up
	TopRight:    func(x, y int) (int, int) { return numTileField + x, numTileField + y }, // shift up and right
	BottomRight: func(x, y int) (int, int) { return numTileField + x, y },                // shift right
	BottomLeft:  func(x, y int) (int, int) { return x, y },                               // no shift
}

func rotateWalls(w Wall, k int) Wall {
	r := w << k     // shift k bits
	r |= r >> 4     // add overlow / rotate
	return r & 0x0f //clear high bits
}

/*
rotate walls clockwise
*/
type rotate func(w Wall) Wall

var wallRotations = [NumTile]rotate{
	TopLeft:     func(w Wall) Wall { return rotateWalls(w, 3) }, // rotate 270 degree- right bit shift 3
	TopRight:    func(w Wall) Wall { return w },                 // rotate 0 degree - no shift
	BottomRight: func(w Wall) Wall { return rotateWalls(w, 1) }, // rotate 90 degree - right bit shift 1
	BottomLeft:  func(w Wall) Wall { return rotateWalls(w, 2) }, // rotate 180 degree - right bit shift 2
}

// Targets represents the target fields for a robot source field combination.
type Targets struct {
	North, South, East, West coord.XY
}

func (t *Targets) calcTargets(b *Board, x0, y0 int) {
	// north
	y := y0
	for b.Field(x0, y).Walls&NorthWall == 0 {
		y++
	}
	t.North.X, t.North.Y = x0, y
	// south
	y = y0
	for b.Field(x0, y).Walls&SouthWall == 0 {
		y--
	}
	t.South.X, t.South.Y = x0, y
	// east
	x := x0
	for b.Field(x, y).Walls&EastWall == 0 {
		x++
	}
	t.East.X, t.East.Y = x, y0
	// west
	x = x0
	for b.Field(x, y).Walls&WestWall == 0 {
		x--
	}
	t.West.X, t.West.Y = x, y0
}

// Field is the type representing a field of a board.
type Field struct {
	Walls   Wall   `json:"walls,omitempty"`
	Symbol  Symbol `json:"symbol,omitempty"`
	Color   Color  `json:"color,omitempty"`
	Targets Targets
}

func (f *Field) String() string {
	return fmt.Sprintf("{%s %s %s %v}", f.Walls, f.Symbol, f.Color, f.Targets)
}

func (f *Field) hasWall(w Wall) bool { return f.Walls&w != 0 }

func (f *Field) addWall(walls ...Wall) {
	for _, w := range walls {
		f.Walls |= w
	}
}

// Board is the type representing a board.
type Board struct {
	Fields [NumField]*Field
}

// New creates a new board instance. Parameter tiles needs to be valid - if not NewBoard will panic.
func New(tileIDs [NumTile]string) *Board {
	b := &Board{}
	// init fields
	for i := 0; i < NumField; i++ {
		b.Fields[i] = new(Field)
	}

	// set tile fields
	for p, id := range tileIDs {
		fields, ok := tiles[id]
		if !ok {
			panic(fmt.Errorf("invalid tileID: %v", id))
		}
		for c, f := range fields {
			x, y := posShifts[p](posRotations[p](coord.X(c), coord.Y(c))) // rotate and shift
			field := b.Fields[coord.Ctob(x, y)]
			field.Walls = wallRotations[p](f.Walls)
			field.Symbol = f.Symbol
			field.Color = f.Color
		}
	}

	// set outer walls
	for x := 0; x < numBoardField; x++ {
		b.Fields[coord.Ctob(x, numBoardField-1)].addWall(NorthWall) // top border
		b.Fields[coord.Ctob(x, 0)].addWall(SouthWall)               // bottom border
	}
	for y := 0; y < numBoardField; y++ {
		b.Fields[coord.Ctob(0, y)].addWall(WestWall)               // left border
		b.Fields[coord.Ctob(numBoardField-1, y)].addWall(EastWall) // right border
	}

	// set center walls
	b.Fields[coord.Ctob(numTileField-1, numTileField-1)].addWall(WestWall, SouthWall) // bottom left center field
	b.Fields[coord.Ctob(numTileField-1, numTileField)].addWall(WestWall, NorthWall)   // top left center field
	b.Fields[coord.Ctob(numTileField, numTileField-1)].addWall(EastWall, SouthWall)   // bottom right center field
	b.Fields[coord.Ctob(numTileField, numTileField)].addWall(EastWall, NorthWall)     // top right center field

	// finally: set neighbor walls
	for x := 0; x < numBoardField; x++ {
		for y := 0; y < numBoardField; y++ {
			// if west field has east wall -> set west wall
			if x > 0 && b.Field(x-1, y).hasWall(EastWall) {
				b.Field(x, y).addWall(WestWall)
			}
			// if east field has west wall -> set east wall
			if x < (numBoardField-1) && b.Field(x+1, y).hasWall(WestWall) {
				b.Field(x, y).addWall(EastWall)
			}
			// if south field has north wall -> set south wall
			if y > 0 && b.Field(x, y-1).hasWall(NorthWall) {
				b.Field(x, y).addWall(SouthWall)
			}
			// if north field has south wall -> set north wall
			if y < (numBoardField-1) && b.Field(x, y+1).hasWall(SouthWall) {
				b.Field(x, y).addWall(NorthWall)
			}
		}
	}

	// calculate routes
	for x := 0; x < numBoardField; x++ {
		for y := 0; y < numBoardField; y++ {
			field := b.Field(x, y)
			field.Targets.calcTargets(b, x, y)
		}
	}

	return b
}

// IsValidCoordinate returns true if the coordinate represents a valid board field, false otherwise.
func (b *Board) IsValidCoordinate(x, y int) bool {
	// in range ?
	if x < 0 || x >= numBoardField || y < 0 || y >= numBoardField {
		return false
	}
	// center field ?
	c1 := numTileField - 1
	c2 := numTileField
	if x >= c1 && x <= c2 && y >= c1 && y <= c2 {
		return false
	}
	return true
}

// Field returns the board field at position x,y.
func (b *Board) Field(x, y int) *Field { return b.Fields[coord.Ctob(x, y)] }

// TargetCoord returns the coordinate of the target.
func (b *Board) TargetCoord(symbol Symbol, color Color) byte {
	for idx, field := range b.Fields {
		if field.Symbol == symbol && field.Color == color {
			return byte(idx)
		}
	}
	panic(fmt.Errorf("invalid target: symbol %s color %s", symbol, color)) // should never happen
}

// MinMoves calculates the minimal moves from each field to the target field.
func (b *Board) MinMoves(cr byte) [NumField]int {
	var minMoves [NumField]int

	// init
	for i := 0; i < NumField; i++ {
		minMoves[i] = -1
	}

	minMoves[cr] = 0
	hsource := []byte{cr}
	vsource := []byte{cr}
	htarget := []byte{}
	vtarget := []byte{}

	moves := 0
	for len(hsource) != 0 || len(vsource) != 0 {
		moves++
		for _, ch := range hsource {
			x0, y0 := coord.Btoc(ch)

			x := x0
			for !b.Field(x, y0).hasWall(WestWall) {
				x--
				c := coord.Ctob(x, y0)
				if minMoves[c] == -1 {
					minMoves[c] = moves
					vtarget = append(vtarget, c)
				}
			}
			x = x0
			for !b.Field(x, y0).hasWall(EastWall) {
				x++
				c := coord.Ctob(x, y0)
				if minMoves[c] == -1 {
					minMoves[c] = moves
					vtarget = append(vtarget, c)
				}
			}
		}

		for _, cv := range vsource {
			x0, y0 := coord.Btoc(cv)

			y := y0
			for !b.Field(x0, y).hasWall(SouthWall) {
				y--
				c := coord.Ctob(x0, y)
				if minMoves[c] == -1 {
					minMoves[c] = moves
					htarget = append(htarget, c)
				}
			}
			y = y0
			for !b.Field(x0, y).hasWall(NorthWall) {
				y++
				c := coord.Ctob(x0, y)
				if minMoves[c] == -1 {
					minMoves[c] = moves
					htarget = append(htarget, c)
				}
			}
		}

		hsource, vsource = htarget, vtarget
		htarget, vtarget = nil, nil
	}
	return minMoves
}
