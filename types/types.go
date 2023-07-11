// Package types provides common game types.
package types

import (
	"fmt"
	"strings"
)

// Maximum fields constants.
const (
	MaxBoardFields = 16
	MaxTileFields  = MaxBoardFields / 2
)

// Target defines the target field of a game.
type Target struct {
	Symbol Symbol `json:"symbol,omitempty"`
	Color  Color  `json:"color,omitempty"`
}

func (t Target) String() string {
	if t.Symbol == Cosmic {
		return string(t.Symbol)
	}
	return fmt.Sprintf("%s%s", t.Color, t.Symbol)
}

// TilePosition defines the position of a tile on the board.
type TilePosition byte

// TilePosition constants.
const (
	TopLeft TilePosition = iota
	TopRight
	BottomLeft
	BottomRight
	numTilePosition
)

// IsValid returns true if the tile position is valid, false otherwise.
func (p TilePosition) IsValid() bool { return p < numTilePosition }

var tilePositionStrs = []string{"topleft", "topright", "bottomleft", "bottomright"}

func (p TilePosition) String() string {
	if !p.IsValid() {
		return fmt.Sprintf("invalid tile position %d", p)
	}
	return tilePositionStrs[p]
}

// TileID defines the id of a type (tile number & side).
type TileID struct {
	SetID  byte `json:"set"`
	TileNo byte `json:"no"`
	Front  bool `json:"front"`
}

func (id TileID) String() string {
	var side byte
	if id.Front {
		side = 'F'
	} else {
		side = 'B'
	}
	return fmt.Sprintf("%c%d%c", id.SetID, id.TileNo, side)
}

// Symbol is the type of a symbol.
type Symbol byte

// Symbol constants.
const (
	Pyramid Symbol = iota
	Star
	Moon
	Saturn
	Cosmic
	numSymbol
)

// IsValid checks if a symbol is valid.
func (s Symbol) IsValid() bool { return s < numSymbol }

var symbolStrs = []string{"Pyramid", "Star", "Moon", "Saturn", "Cosmic"}

func (s Symbol) String() string {
	if !s.IsValid() {
		return fmt.Sprintf("invalid symbol %d", s)
	}
	return symbolStrs[s]
}

// Symbols is the set of valid symbols.
var Symbols = []Symbol{Pyramid, Star, Moon, Saturn, Cosmic}

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

// Coordinate defines a two dimensional
type Coordinate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c Coordinate) String() string { return fmt.Sprintf("%d,%d", c.X, c.Y) }

// IsCenterField returns true if the coordinate refers a center field, false otherwise.
func IsCenterField(x, y int) bool {
	return x == MaxTileFields-1 && y == MaxTileFields-1 || x == MaxTileFields-1 && y == MaxTileFields || x == MaxTileFields && y == MaxTileFields-1 || x == MaxTileFields && y == MaxTileFields
}

// IsInRange checks if the x and y value of the coordinate are valid.
func IsInRange(x, y int) bool {
	return x >= 0 && x < MaxBoardFields && y >= 0 && y < MaxBoardFields
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
