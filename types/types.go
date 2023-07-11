// Package types provides common game types.
package types

import (
	"fmt"
	"strconv"
	"strings"
)

// Maximum fields constants.
const (
	MaxBoardFields = 16
	MaxTileFields  = MaxBoardFields / 2
)

// Parameter name constants.
const (
	TopLeftTilePrm     = "ttl"
	TopRightTilePrm    = "ttr"
	BottomRightTilePrm = "tbr"
	BottomLeftTilePrm  = "tbl"

	YellowRobotPrm = "ry"
	RedRobotPrm    = "rr"
	GreenRobotPrm  = "rg"
	BlueRobotPrm   = "rb"
	SilverRobotPrm = "rs"

	TargetPrm = "t"
)

// Default parameters.
var (
	DefTopLeftTile     = TileID{SetID: 'A', TileNo: 1, Front: true}
	DefTopRightTile    = TileID{SetID: 'A', TileNo: 2, Front: true}
	DefBottomRightTile = TileID{SetID: 'A', TileNo: 3, Front: true}
	DefBottomLeftTile  = TileID{SetID: 'A', TileNo: 4, Front: true}

	DefYellowRobot = Coordinate{X: 0, Y: 0}
	DefRedRobot    = Coordinate{X: 1, Y: 0}
	DefGreenRobot  = Coordinate{X: 2, Y: 0}
	DefBlueRobot   = Coordinate{X: 3, Y: 0}
	DefSilverRobot = Coordinate{X: -1, Y: -1}

	DefTarget = Targets[TnCosmic]
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

// Target name constants.
const (
	TnYellowPyramid = "yellowPyramid"
	TnRedPyramid    = "redPyramid"
	TnGreenPyramid  = "greenPyramid"
	TnBluePyramid   = "bluePyramid"

	TnYellowStar = "yellowStar"
	TnRedStar    = "redStar"
	TnGreenStar  = "greenStar"
	TnBlueStar   = "blueStar"

	TnYellowMoon = "yellowMoon"
	TnRedMoon    = "redMoon"
	TnGreenMoon  = "greenMoon"
	TnBlueMoon   = "blueMoon"

	TnYellowSaturn = "yellowSaturn"
	TnRedSaturn    = "redSaturn"
	TnGreenSaturn  = "greenSaturn"
	TnBlueSaturn   = "blueSaturn"

	TnCosmic = "Cosmic"
)

// Targets is a map of valid targets.
var Targets = map[string]Target{
	TnYellowPyramid: {Symbol: Pyramid, Color: Yellow},
	TnRedPyramid:    {Symbol: Pyramid, Color: Red},
	TnGreenPyramid:  {Symbol: Pyramid, Color: Green},
	TnBluePyramid:   {Symbol: Pyramid, Color: Blue},

	TnYellowStar: {Symbol: Star, Color: Yellow},
	TnRedStar:    {Symbol: Star, Color: Red},
	TnGreenStar:  {Symbol: Star, Color: Green},
	TnBlueStar:   {Symbol: Star, Color: Blue},

	TnYellowMoon: {Symbol: Moon, Color: Yellow},
	TnRedMoon:    {Symbol: Moon, Color: Red},
	TnGreenMoon:  {Symbol: Moon, Color: Green},
	TnBlueMoon:   {Symbol: Moon, Color: Blue},

	TnYellowSaturn: {Symbol: Saturn, Color: Yellow},
	TnRedSaturn:    {Symbol: Saturn, Color: Red},
	TnGreenSaturn:  {Symbol: Saturn, Color: Green},
	TnBlueSaturn:   {Symbol: Saturn, Color: Blue},

	TnCosmic: {Symbol: Cosmic, Color: CosmicColor},
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

// SetID is the id of the tile set.
type SetID byte

// TileNo is the type of the tile number.
type TileNo byte

// TileID defines the id of a type (tile number & side).
type TileID struct {
	SetID  SetID  `json:"set"`
	TileNo TileNo `json:"no"`
	Front  bool   `json:"front"`
}

func convertBoolToSide(front bool) byte {
	if front {
		return 'F'
	}
	return 'B'
}

func convertSideToBool(b byte) (bool, error) {
	switch b {
	case 'F', 'f':
		return true, nil
	case 'B', 'b':
		return false, nil
	default:
		return false, fmt.Errorf("invalid tile side: %c", b)
	}
}

// ParseTileID interprets a string s as tile ID.
func ParseTileID(s string) (id TileID, err error) {
	if len(s) != 3 {
		return id, fmt.Errorf("invalid tile: %s", s)
	}

	no64, err := strconv.ParseInt(s[1:2], 10, 0)
	if err != nil {
		return id, err
	}
	front, err := convertSideToBool(s[2])
	if err != nil {
		return id, err
	}
	return TileID{SetID: SetID(s[0]), TileNo: TileNo(no64), Front: front}, nil
}

func (id TileID) String() string {
	return fmt.Sprintf("%c%d%c", id.SetID, id.TileNo, convertBoolToSide(id.Front))
}

// Set implements the flag.Value interface.
func (id *TileID) Set(value string) (err error) {
	if *id, err = ParseTileID(value); err != nil {
		return err
	}
	return nil
}

// IsValid returns true if the tile ID is valid, false otherwise.
func (id TileID) IsValid() bool {
	if id.SetID != 'A' && id.SetID != 'B' {
		return false
	}
	if id.TileNo < 1 || id.TileNo > 4 {
		return false
	}
	return true
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

// ParseCoordinate interprets a string s in the form of x,y as coordinate.
func ParseCoordinate(s string) (Coordinate, error) {
	pos := Coordinate{X: -1, Y: -1}

	p := strings.Split(s, ",")
	if len(p) != 2 {
		return pos, strconv.ErrSyntax
	}
	x64, err := strconv.ParseInt(p[0], 10, 0)
	if err != nil {
		return pos, err
	}
	y64, err := strconv.ParseInt(p[1], 10, 0)
	if err != nil {
		return pos, err
	}

	pos.X = int(x64)
	pos.Y = int(y64)

	return pos, nil
}

func (c Coordinate) String() string { return fmt.Sprintf("%d,%d", c.X, c.Y) }

// Set implements the flag.Value interface.
func (c *Coordinate) Set(value string) (err error) {
	if *c, err = ParseCoordinate(value); err != nil {
		return err
	}
	if c.IsCenterField() {
		return fmt.Errorf("invalid position: center field %s", c)
	}
	return nil
}

// IsCenterField returns true if the coordinate refers a center field, false otherwise.
func (c Coordinate) IsCenterField() bool {
	return c.X == MaxTileFields-1 && c.Y == MaxTileFields-1 || c.X == MaxTileFields-1 && c.Y == MaxTileFields || c.X == MaxTileFields && c.Y == MaxTileFields-1 || c.X == MaxTileFields && c.Y == MaxTileFields
}

// IsInRange checks if the x and y value of the coordinate are valid.
func (c Coordinate) IsInRange() bool {
	return c.X >= 0 && c.X < MaxBoardFields && c.Y >= 0 && c.Y < MaxBoardFields
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
