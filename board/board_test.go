package board

import (
	"testing"

	"github.com/go-ricrob/game/coord"
)

func TestBoard(t *testing.T) {
	tests := []struct {
		name   string
		tiles  [NumTile]string
		fields map[byte]Field
	}{
		{
			"default board",
			[NumTile]string{
				TopLeft:     "A1F",
				TopRight:    "A2F",
				BottomRight: "A3F",
				BottomLeft:  "A4F",
			},
			map[byte]Field{

				/*
					{Set: 'A', No: 1, Front: true}
				*/

				coord.Ctob(0, 8): {Walls: WestWall | NorthWall},
				coord.Ctob(6, 8): {Walls: EastWall},
				coord.Ctob(7, 8): {Walls: WestWall | NorthWall},

				coord.Ctob(0, 9): {Walls: WestWall | SouthWall},
				coord.Ctob(2, 9): {Walls: NorthWall},
				coord.Ctob(7, 9): {Walls: SouthWall | NorthWall},

				coord.Ctob(0, 10): {Walls: WestWall},
				coord.Ctob(2, 10): {Walls: SouthWall | EastWall, Symbol: Star, Color: Red},
				coord.Ctob(3, 10): {Walls: WestWall},
				coord.Ctob(7, 10): {Walls: SouthWall | EastWall, Symbol: Cosmic},

				coord.Ctob(0, 11): {Walls: WestWall},
				coord.Ctob(4, 11): {Walls: EastWall},
				coord.Ctob(5, 11): {Walls: WestWall | NorthWall, Symbol: Saturn, Color: Green},

				coord.Ctob(0, 12): {Walls: WestWall},
				coord.Ctob(1, 12): {Walls: EastWall | NorthWall, Symbol: Pyramid, Color: Yellow},
				coord.Ctob(2, 12): {Walls: WestWall},
				coord.Ctob(5, 12): {Walls: SouthWall},

				coord.Ctob(0, 13): {Walls: WestWall},
				coord.Ctob(1, 13): {Walls: SouthWall},
				coord.Ctob(6, 13): {Walls: NorthWall},

				coord.Ctob(0, 14): {Walls: WestWall},
				coord.Ctob(5, 14): {Walls: EastWall},
				coord.Ctob(6, 14): {Walls: WestWall | SouthWall, Symbol: Moon, Color: Blue},

				coord.Ctob(0, 15): {Walls: WestWall | NorthWall},
				coord.Ctob(1, 15): {Walls: NorthWall},
				coord.Ctob(2, 15): {Walls: NorthWall},
				coord.Ctob(3, 15): {Walls: EastWall | NorthWall},
				coord.Ctob(4, 15): {Walls: WestWall | NorthWall},
				coord.Ctob(5, 15): {Walls: NorthWall},
				coord.Ctob(6, 15): {Walls: NorthWall},
				coord.Ctob(7, 15): {Walls: NorthWall},

				/*
					{Set: 'A', No: 2, Front: true}
				*/

				coord.Ctob(8, 8):  {Walls: EastWall | NorthWall},
				coord.Ctob(9, 8):  {Walls: WestWall},
				coord.Ctob(15, 8): {Walls: EastWall},

				coord.Ctob(8, 9):  {Walls: SouthWall},
				coord.Ctob(11, 9): {Walls: EastWall | NorthWall, Symbol: Moon, Color: Yellow},
				coord.Ctob(12, 9): {Walls: WestWall},
				coord.Ctob(14, 9): {Walls: NorthWall},
				coord.Ctob(15, 9): {Walls: EastWall},

				coord.Ctob(8, 10):  {Walls: WestWall},
				coord.Ctob(11, 10): {Walls: SouthWall},
				coord.Ctob(13, 10): {Walls: EastWall},
				coord.Ctob(14, 10): {Walls: WestWall | SouthWall, Symbol: Star, Color: Green},
				coord.Ctob(15, 10): {Walls: EastWall},

				coord.Ctob(15, 11): {Walls: EastWall | NorthWall},

				coord.Ctob(9, 12):  {Walls: NorthWall},
				coord.Ctob(15, 12): {Walls: EastWall | SouthWall},

				coord.Ctob(9, 13):  {Walls: EastWall | SouthWall, Symbol: Pyramid, Color: Blue},
				coord.Ctob(10, 13): {Walls: WestWall},
				coord.Ctob(15, 13): {Walls: EastWall},

				coord.Ctob(12, 14): {Walls: EastWall},
				coord.Ctob(13, 14): {Walls: WestWall | NorthWall, Symbol: Saturn, Color: Red},
				coord.Ctob(15, 14): {Walls: EastWall},

				coord.Ctob(8, 15):  {Walls: NorthWall},
				coord.Ctob(9, 15):  {Walls: NorthWall},
				coord.Ctob(10, 15): {Walls: NorthWall},
				coord.Ctob(11, 15): {Walls: EastWall | NorthWall},
				coord.Ctob(12, 15): {Walls: WestWall | NorthWall},
				coord.Ctob(13, 15): {Walls: SouthWall | NorthWall},
				coord.Ctob(14, 15): {Walls: NorthWall},
				coord.Ctob(15, 15): {Walls: EastWall | NorthWall},

				/*
					{Set: 'A', No: 3, Front: true}
				*/

				coord.Ctob(8, 0):  {Walls: SouthWall},
				coord.Ctob(9, 0):  {Walls: SouthWall},
				coord.Ctob(10, 0): {Walls: SouthWall},
				coord.Ctob(11, 0): {Walls: SouthWall | NorthWall},
				coord.Ctob(12, 0): {Walls: SouthWall},
				coord.Ctob(13, 0): {Walls: EastWall | SouthWall},
				coord.Ctob(14, 0): {Walls: WestWall | SouthWall},
				coord.Ctob(15, 0): {Walls: EastWall | SouthWall},

				coord.Ctob(11, 1): {Walls: EastWall | SouthWall, Symbol: Moon, Color: Red},
				coord.Ctob(12, 1): {Walls: WestWall},
				coord.Ctob(14, 1): {Walls: NorthWall},
				coord.Ctob(15, 1): {Walls: EastWall},

				coord.Ctob(13, 2): {Walls: EastWall},
				coord.Ctob(14, 2): {Walls: WestWall | SouthWall, Symbol: Pyramid, Color: Green},
				coord.Ctob(15, 2): {Walls: EastWall},

				coord.Ctob(8, 3):  {Walls: EastWall},
				coord.Ctob(9, 3):  {Walls: WestWall | NorthWall, Symbol: Star, Color: Yellow},
				coord.Ctob(15, 3): {Walls: EastWall},

				coord.Ctob(9, 4):  {Walls: SouthWall},
				coord.Ctob(15, 4): {Walls: EastWall},

				coord.Ctob(15, 5): {Walls: EastWall | NorthWall},

				coord.Ctob(8, 6):  {Walls: NorthWall},
				coord.Ctob(12, 6): {Walls: EastWall | NorthWall, Symbol: Saturn, Color: Blue},
				coord.Ctob(13, 6): {Walls: WestWall},
				coord.Ctob(15, 6): {Walls: EastWall | SouthWall},

				coord.Ctob(8, 7):  {Walls: EastWall | SouthWall},
				coord.Ctob(9, 7):  {Walls: WestWall},
				coord.Ctob(12, 7): {Walls: SouthWall},
				coord.Ctob(15, 7): {Walls: EastWall},

				/*
					{Set: 'A', No: 4, Front: true}
				*/

				coord.Ctob(0, 0): {Walls: WestWall | SouthWall},
				coord.Ctob(1, 0): {Walls: SouthWall | NorthWall},
				coord.Ctob(2, 0): {Walls: SouthWall},
				coord.Ctob(3, 0): {Walls: SouthWall},
				coord.Ctob(4, 0): {Walls: SouthWall},
				coord.Ctob(5, 0): {Walls: EastWall | SouthWall},
				coord.Ctob(6, 0): {Walls: WestWall | SouthWall},
				coord.Ctob(7, 0): {Walls: SouthWall},

				coord.Ctob(0, 1): {Walls: WestWall},
				coord.Ctob(1, 1): {Walls: EastWall | SouthWall, Symbol: Pyramid, Color: Red},
				coord.Ctob(2, 1): {Walls: WestWall},

				coord.Ctob(0, 2): {Walls: WestWall},
				coord.Ctob(4, 2): {Walls: EastWall | NorthWall, Symbol: Star, Color: Blue},
				coord.Ctob(5, 2): {Walls: WestWall},

				coord.Ctob(0, 3): {Walls: WestWall | NorthWall},
				coord.Ctob(4, 3): {Walls: SouthWall},

				coord.Ctob(0, 4): {Walls: WestWall | SouthWall},

				coord.Ctob(0, 5): {Walls: WestWall},

				coord.Ctob(0, 6): {Walls: WestWall},
				coord.Ctob(1, 6): {Walls: EastWall},
				coord.Ctob(2, 6): {Walls: WestWall | NorthWall, Symbol: Moon, Color: Green},
				coord.Ctob(5, 6): {Walls: NorthWall},
				coord.Ctob(7, 6): {Walls: NorthWall},

				coord.Ctob(0, 7): {Walls: WestWall},
				coord.Ctob(2, 7): {Walls: SouthWall},
				coord.Ctob(4, 7): {Walls: EastWall},
				coord.Ctob(5, 7): {Walls: WestWall | SouthWall, Symbol: Saturn, Color: Yellow},
				coord.Ctob(6, 7): {Walls: EastWall},
				coord.Ctob(7, 7): {Walls: WestWall | SouthWall},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := New(test.tiles)
			for p, f1 := range b.Fields {
				x, y := coord.Btoc(byte(p))
				if f2, ok := test.fields[byte(p)]; ok {
					if f1.Walls != f2.Walls || f1.Symbol != f2.Symbol || f1.Color != f2.Color {
						t.Fatalf("field %d,%d: %s %s %s - expected %s %s %s", x, y, f1.Walls, f1.Symbol, f1.Color, f2.Walls, f2.Symbol, f2.Color)
					}
				} else {
					if f1.Walls != 0 || f1.Symbol != 0 || f1.Color != 0 {
						t.Fatalf("field %d,%d: %s %s %s - expected %s %s %s", x, y, f1.Walls, f1.Symbol, f1.Color, f2.Walls, f2.Symbol, f2.Color)
					}
				}
			}
		})
	}
}
