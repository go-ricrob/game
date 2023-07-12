package board

import (
	"testing"

	"github.com/go-ricrob/game/types"
)

func TestBoard(t *testing.T) {
	tests := []struct {
		name   string
		tiles  map[TilePosition]string
		fields map[types.Coordinate]Field
	}{
		{
			"default board",
			map[TilePosition]string{
				TopLeft:     "A1F",
				TopRight:    "A2F",
				BottomRight: "A3F",
				BottomLeft:  "A4F",
			},
			map[types.Coordinate]Field{

				/*
					{Set: 'A', No: 1, Front: true}
				*/

				{X: 0, Y: 8}: {walls: West | North},
				{X: 1, Y: 8}: {},
				{X: 2, Y: 8}: {},
				{X: 3, Y: 8}: {},
				{X: 4, Y: 8}: {},
				{X: 5, Y: 8}: {},
				{X: 6, Y: 8}: {walls: East},
				{X: 7, Y: 8}: {walls: West | North},

				{X: 0, Y: 9}: {walls: West | South},
				{X: 1, Y: 9}: {},
				{X: 2, Y: 9}: {walls: North},
				{X: 3, Y: 9}: {},
				{X: 4, Y: 9}: {},
				{X: 5, Y: 9}: {},
				{X: 6, Y: 9}: {},
				{X: 7, Y: 9}: {walls: South | North},

				{X: 0, Y: 10}: {walls: West},
				{X: 1, Y: 10}: {},
				{X: 2, Y: 10}: {walls: South | East, symbol: Star, color: types.Red},
				{X: 3, Y: 10}: {walls: West},
				{X: 4, Y: 10}: {},
				{X: 5, Y: 10}: {},
				{X: 6, Y: 10}: {},
				{X: 7, Y: 10}: {walls: South | East, symbol: Cosmic, color: types.CosmicColor},

				{X: 0, Y: 11}: {walls: West},
				{X: 1, Y: 11}: {},
				{X: 2, Y: 11}: {},
				{X: 3, Y: 11}: {},
				{X: 4, Y: 11}: {walls: East},
				{X: 5, Y: 11}: {walls: West | North, symbol: Saturn, color: types.Green},
				{X: 6, Y: 11}: {},
				{X: 7, Y: 11}: {},

				{X: 0, Y: 12}: {walls: West},
				{X: 1, Y: 12}: {walls: East | North, symbol: Pyramid, color: types.Yellow},
				{X: 2, Y: 12}: {walls: West},
				{X: 3, Y: 12}: {},
				{X: 4, Y: 12}: {},
				{X: 5, Y: 12}: {walls: South},
				{X: 6, Y: 12}: {},
				{X: 7, Y: 12}: {},

				{X: 0, Y: 13}: {walls: West},
				{X: 1, Y: 13}: {walls: South},
				{X: 2, Y: 13}: {},
				{X: 3, Y: 13}: {},
				{X: 4, Y: 13}: {},
				{X: 5, Y: 13}: {},
				{X: 6, Y: 13}: {walls: North},
				{X: 7, Y: 13}: {},

				{X: 0, Y: 14}: {walls: West},
				{X: 1, Y: 14}: {},
				{X: 2, Y: 14}: {},
				{X: 3, Y: 14}: {},
				{X: 4, Y: 14}: {},
				{X: 5, Y: 14}: {walls: East},
				{X: 6, Y: 14}: {walls: West | South, symbol: Moon, color: types.Blue},
				{X: 7, Y: 14}: {},

				{X: 0, Y: 15}: {walls: West | North},
				{X: 1, Y: 15}: {walls: North},
				{X: 2, Y: 15}: {walls: North},
				{X: 3, Y: 15}: {walls: East | North},
				{X: 4, Y: 15}: {walls: West | North},
				{X: 5, Y: 15}: {walls: North},
				{X: 6, Y: 15}: {walls: North},
				{X: 7, Y: 15}: {walls: North},

				/*
					{Set: 'A', No: 2, Front: true}
				*/

				{X: 8, Y: 8}:  {walls: East | North},
				{X: 9, Y: 8}:  {walls: West},
				{X: 10, Y: 8}: {},
				{X: 11, Y: 8}: {},
				{X: 12, Y: 8}: {},
				{X: 13, Y: 8}: {},
				{X: 14, Y: 8}: {},
				{X: 15, Y: 8}: {walls: East},

				{X: 8, Y: 9}:  {walls: South},
				{X: 9, Y: 9}:  {},
				{X: 10, Y: 9}: {},
				{X: 11, Y: 9}: {walls: East | North, symbol: Moon, color: types.Yellow},
				{X: 12, Y: 9}: {walls: West},
				{X: 13, Y: 9}: {},
				{X: 14, Y: 9}: {walls: North},
				{X: 15, Y: 9}: {walls: East},

				{X: 8, Y: 10}:  {walls: West},
				{X: 9, Y: 10}:  {},
				{X: 10, Y: 10}: {},
				{X: 11, Y: 10}: {walls: South},
				{X: 12, Y: 10}: {},
				{X: 13, Y: 10}: {walls: East},
				{X: 14, Y: 10}: {walls: West | South, symbol: Star, color: types.Green},
				{X: 15, Y: 10}: {walls: East},

				{X: 8, Y: 11}:  {},
				{X: 9, Y: 11}:  {},
				{X: 10, Y: 11}: {},
				{X: 11, Y: 11}: {},
				{X: 12, Y: 11}: {},
				{X: 13, Y: 11}: {},
				{X: 14, Y: 11}: {},
				{X: 15, Y: 11}: {walls: East | North},

				{X: 8, Y: 12}:  {},
				{X: 9, Y: 12}:  {walls: North},
				{X: 10, Y: 12}: {},
				{X: 11, Y: 12}: {},
				{X: 12, Y: 12}: {},
				{X: 13, Y: 12}: {},
				{X: 14, Y: 12}: {},
				{X: 15, Y: 12}: {walls: East | South},

				{X: 8, Y: 13}:  {},
				{X: 9, Y: 13}:  {walls: East | South, symbol: Pyramid, color: types.Blue},
				{X: 10, Y: 13}: {walls: West},
				{X: 11, Y: 13}: {},
				{X: 12, Y: 13}: {},
				{X: 13, Y: 13}: {},
				{X: 14, Y: 13}: {},
				{X: 15, Y: 13}: {walls: East},

				{X: 8, Y: 14}:  {},
				{X: 9, Y: 14}:  {},
				{X: 10, Y: 14}: {},
				{X: 11, Y: 14}: {},
				{X: 12, Y: 14}: {walls: East},
				{X: 13, Y: 14}: {walls: West | North, symbol: Saturn, color: types.Red},
				{X: 14, Y: 14}: {},
				{X: 15, Y: 14}: {walls: East},

				{X: 8, Y: 15}:  {walls: North},
				{X: 9, Y: 15}:  {walls: North},
				{X: 10, Y: 15}: {walls: North},
				{X: 11, Y: 15}: {walls: East | North},
				{X: 12, Y: 15}: {walls: West | North},
				{X: 13, Y: 15}: {walls: South | North},
				{X: 14, Y: 15}: {walls: North},
				{X: 15, Y: 15}: {walls: East | North},

				/*
					{Set: 'A', No: 3, Front: true}
				*/

				{X: 8, Y: 0}:  {walls: South},
				{X: 9, Y: 0}:  {walls: South},
				{X: 10, Y: 0}: {walls: South},
				{X: 11, Y: 0}: {walls: South | North},
				{X: 12, Y: 0}: {walls: South},
				{X: 13, Y: 0}: {walls: East | South},
				{X: 14, Y: 0}: {walls: West | South},
				{X: 15, Y: 0}: {walls: East | South},

				{X: 8, Y: 1}:  {},
				{X: 9, Y: 1}:  {},
				{X: 10, Y: 1}: {},
				{X: 11, Y: 1}: {walls: East | South, symbol: Moon, color: types.Red},
				{X: 12, Y: 1}: {walls: West},
				{X: 13, Y: 1}: {},
				{X: 14, Y: 1}: {walls: North},
				{X: 15, Y: 1}: {walls: East},

				{X: 8, Y: 2}:  {},
				{X: 9, Y: 2}:  {},
				{X: 10, Y: 2}: {},
				{X: 11, Y: 2}: {},
				{X: 12, Y: 2}: {},
				{X: 13, Y: 2}: {walls: East},
				{X: 14, Y: 2}: {walls: West | South, symbol: Pyramid, color: types.Green},
				{X: 15, Y: 2}: {walls: East},

				{X: 8, Y: 3}:  {walls: East},
				{X: 9, Y: 3}:  {walls: West | North, symbol: Star, color: types.Yellow},
				{X: 10, Y: 3}: {},
				{X: 11, Y: 3}: {},
				{X: 12, Y: 3}: {},
				{X: 13, Y: 3}: {},
				{X: 14, Y: 3}: {},
				{X: 15, Y: 3}: {walls: East},

				{X: 8, Y: 4}:  {},
				{X: 9, Y: 4}:  {walls: South},
				{X: 10, Y: 4}: {},
				{X: 11, Y: 4}: {},
				{X: 12, Y: 4}: {},
				{X: 13, Y: 4}: {},
				{X: 14, Y: 4}: {},
				{X: 15, Y: 4}: {walls: East},

				{X: 8, Y: 5}:  {},
				{X: 9, Y: 5}:  {},
				{X: 10, Y: 5}: {},
				{X: 11, Y: 5}: {},
				{X: 12, Y: 5}: {},
				{X: 13, Y: 5}: {},
				{X: 14, Y: 5}: {},
				{X: 15, Y: 5}: {walls: East | North},

				{X: 8, Y: 6}:  {walls: North},
				{X: 9, Y: 6}:  {},
				{X: 10, Y: 6}: {},
				{X: 11, Y: 6}: {},
				{X: 12, Y: 6}: {walls: East | North, symbol: Saturn, color: types.Blue},
				{X: 13, Y: 6}: {walls: West},
				{X: 14, Y: 6}: {},
				{X: 15, Y: 6}: {walls: East | South},

				{X: 8, Y: 7}:  {walls: East | South},
				{X: 9, Y: 7}:  {walls: West},
				{X: 10, Y: 7}: {},
				{X: 11, Y: 7}: {},
				{X: 12, Y: 7}: {walls: South},
				{X: 13, Y: 7}: {},
				{X: 14, Y: 7}: {},
				{X: 15, Y: 7}: {walls: East},

				/*
					{Set: 'A', No: 4, Front: true}
				*/

				{X: 0, Y: 0}: {walls: West | South},
				{X: 1, Y: 0}: {walls: South | North},
				{X: 2, Y: 0}: {walls: South},
				{X: 3, Y: 0}: {walls: South},
				{X: 4, Y: 0}: {walls: South},
				{X: 5, Y: 0}: {walls: East | South},
				{X: 6, Y: 0}: {walls: West | South},
				{X: 7, Y: 0}: {walls: South},

				{X: 0, Y: 1}: {walls: West},
				{X: 1, Y: 1}: {walls: East | South, symbol: Pyramid, color: types.Red},
				{X: 2, Y: 1}: {walls: West},
				{X: 3, Y: 1}: {},
				{X: 4, Y: 1}: {},
				{X: 5, Y: 1}: {},
				{X: 6, Y: 1}: {},
				{X: 7, Y: 1}: {},

				{X: 0, Y: 2}: {walls: West},
				{X: 1, Y: 2}: {},
				{X: 2, Y: 2}: {},
				{X: 3, Y: 2}: {},
				{X: 4, Y: 2}: {walls: East | North, symbol: Star, color: types.Blue},
				{X: 5, Y: 2}: {walls: West},
				{X: 6, Y: 2}: {},
				{X: 7, Y: 2}: {},

				{X: 0, Y: 3}: {walls: West | North},
				{X: 1, Y: 3}: {},
				{X: 2, Y: 3}: {},
				{X: 3, Y: 3}: {},
				{X: 4, Y: 3}: {walls: South},
				{X: 5, Y: 3}: {},
				{X: 6, Y: 3}: {},
				{X: 7, Y: 3}: {},

				{X: 0, Y: 4}: {walls: West | South},
				{X: 1, Y: 4}: {},
				{X: 2, Y: 4}: {},
				{X: 3, Y: 4}: {},
				{X: 4, Y: 4}: {},
				{X: 5, Y: 4}: {},
				{X: 6, Y: 4}: {},
				{X: 7, Y: 4}: {},

				{X: 0, Y: 5}: {walls: West},
				{X: 1, Y: 5}: {},
				{X: 2, Y: 5}: {},
				{X: 3, Y: 5}: {},
				{X: 4, Y: 5}: {},
				{X: 5, Y: 5}: {},
				{X: 6, Y: 5}: {},
				{X: 7, Y: 5}: {},

				{X: 0, Y: 6}: {walls: West},
				{X: 1, Y: 6}: {walls: East},
				{X: 2, Y: 6}: {walls: West | North, symbol: Moon, color: types.Green},
				{X: 3, Y: 6}: {},
				{X: 4, Y: 6}: {},
				{X: 5, Y: 6}: {walls: North},
				{X: 6, Y: 6}: {},
				{X: 7, Y: 6}: {walls: North},

				{X: 0, Y: 7}: {walls: West},
				{X: 1, Y: 7}: {},
				{X: 2, Y: 7}: {walls: South},
				{X: 3, Y: 7}: {},
				{X: 4, Y: 7}: {walls: East},
				{X: 5, Y: 7}: {walls: West | South, symbol: Saturn, color: types.Yellow},
				{X: 6, Y: 7}: {walls: East},
				{X: 7, Y: 7}: {walls: West | South},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := New(test.tiles)
			for p, f1 := range test.fields {
				f2 := b.Field(p.X, p.Y)
				// if !reflect.DeepEqual(f1, f2) { // no moves check yet
				if f1.walls != f2.walls || f1.symbol != f2.symbol || f1.color != f2.color {
					t.Fatalf("field %d,%d: %v - expected %v", p.X, p.Y, f2, f1)
				}
			}
		})
	}
}
