package board

import (
	"testing"

	"github.com/go-ricrob/game/types"
)

func TestBoard(t *testing.T) {
	tests := []struct {
		name    string
		tileIDs map[types.TilePosition]types.TileID
		fields  map[types.Coordinate]Field
	}{
		{
			"default board",
			map[types.TilePosition]types.TileID{
				types.TopLeft:     {SetID: 'A', TileNo: 1, Front: true},
				types.TopRight:    {SetID: 'A', TileNo: 2, Front: true},
				types.BottomRight: {SetID: 'A', TileNo: 3, Front: true},
				types.BottomLeft:  {SetID: 'A', TileNo: 4, Front: true},
			},
			map[types.Coordinate]Field{

				/*
					{Set: 'A', No: 1, Front: true}
				*/

				{X: 0, Y: 8}: {walls: types.West | types.North},
				{X: 1, Y: 8}: {},
				{X: 2, Y: 8}: {},
				{X: 3, Y: 8}: {},
				{X: 4, Y: 8}: {},
				{X: 5, Y: 8}: {},
				{X: 6, Y: 8}: {walls: types.East},
				{X: 7, Y: 8}: {walls: types.West | types.North},

				{X: 0, Y: 9}: {walls: types.West | types.South},
				{X: 1, Y: 9}: {},
				{X: 2, Y: 9}: {walls: types.North},
				{X: 3, Y: 9}: {},
				{X: 4, Y: 9}: {},
				{X: 5, Y: 9}: {},
				{X: 6, Y: 9}: {},
				{X: 7, Y: 9}: {walls: types.South | types.North},

				{X: 0, Y: 10}: {walls: types.West},
				{X: 1, Y: 10}: {},
				{X: 2, Y: 10}: {walls: types.South | types.East, target: types.Targets[types.TnRedStar]},
				{X: 3, Y: 10}: {walls: types.West},
				{X: 4, Y: 10}: {},
				{X: 5, Y: 10}: {},
				{X: 6, Y: 10}: {},
				{X: 7, Y: 10}: {walls: types.South | types.East, target: types.Targets[types.TnCosmic]},

				{X: 0, Y: 11}: {walls: types.West},
				{X: 1, Y: 11}: {},
				{X: 2, Y: 11}: {},
				{X: 3, Y: 11}: {},
				{X: 4, Y: 11}: {walls: types.East},
				{X: 5, Y: 11}: {walls: types.West | types.North, target: types.Targets[types.TnGreenSaturn]},
				{X: 6, Y: 11}: {},
				{X: 7, Y: 11}: {},

				{X: 0, Y: 12}: {walls: types.West},
				{X: 1, Y: 12}: {walls: types.East | types.North, target: types.Targets[types.TnYellowPyramid]},
				{X: 2, Y: 12}: {walls: types.West},
				{X: 3, Y: 12}: {},
				{X: 4, Y: 12}: {},
				{X: 5, Y: 12}: {walls: types.South},
				{X: 6, Y: 12}: {},
				{X: 7, Y: 12}: {},

				{X: 0, Y: 13}: {walls: types.West},
				{X: 1, Y: 13}: {walls: types.South},
				{X: 2, Y: 13}: {},
				{X: 3, Y: 13}: {},
				{X: 4, Y: 13}: {},
				{X: 5, Y: 13}: {},
				{X: 6, Y: 13}: {walls: types.North},
				{X: 7, Y: 13}: {},

				{X: 0, Y: 14}: {walls: types.West},
				{X: 1, Y: 14}: {},
				{X: 2, Y: 14}: {},
				{X: 3, Y: 14}: {},
				{X: 4, Y: 14}: {},
				{X: 5, Y: 14}: {walls: types.East},
				{X: 6, Y: 14}: {walls: types.West | types.South, target: types.Targets[types.TnBlueMoon]},
				{X: 7, Y: 14}: {},

				{X: 0, Y: 15}: {walls: types.West | types.North},
				{X: 1, Y: 15}: {walls: types.North},
				{X: 2, Y: 15}: {walls: types.North},
				{X: 3, Y: 15}: {walls: types.East | types.North},
				{X: 4, Y: 15}: {walls: types.West | types.North},
				{X: 5, Y: 15}: {walls: types.North},
				{X: 6, Y: 15}: {walls: types.North},
				{X: 7, Y: 15}: {walls: types.North},

				/*
					{Set: 'A', No: 2, Front: true}
				*/

				{X: 8, Y: 8}:  {walls: types.East | types.North},
				{X: 9, Y: 8}:  {walls: types.West},
				{X: 10, Y: 8}: {},
				{X: 11, Y: 8}: {},
				{X: 12, Y: 8}: {},
				{X: 13, Y: 8}: {},
				{X: 14, Y: 8}: {},
				{X: 15, Y: 8}: {walls: types.East},

				{X: 8, Y: 9}:  {walls: types.South},
				{X: 9, Y: 9}:  {},
				{X: 10, Y: 9}: {},
				{X: 11, Y: 9}: {walls: types.East | types.North, target: types.Targets[types.TnYellowMoon]},
				{X: 12, Y: 9}: {walls: types.West},
				{X: 13, Y: 9}: {},
				{X: 14, Y: 9}: {walls: types.North},
				{X: 15, Y: 9}: {walls: types.East},

				{X: 8, Y: 10}:  {walls: types.West},
				{X: 9, Y: 10}:  {},
				{X: 10, Y: 10}: {},
				{X: 11, Y: 10}: {walls: types.South},
				{X: 12, Y: 10}: {},
				{X: 13, Y: 10}: {walls: types.East},
				{X: 14, Y: 10}: {walls: types.West | types.South, target: types.Targets[types.TnGreenStar]},
				{X: 15, Y: 10}: {walls: types.East},

				{X: 8, Y: 11}:  {},
				{X: 9, Y: 11}:  {},
				{X: 10, Y: 11}: {},
				{X: 11, Y: 11}: {},
				{X: 12, Y: 11}: {},
				{X: 13, Y: 11}: {},
				{X: 14, Y: 11}: {},
				{X: 15, Y: 11}: {walls: types.East | types.North},

				{X: 8, Y: 12}:  {},
				{X: 9, Y: 12}:  {walls: types.North},
				{X: 10, Y: 12}: {},
				{X: 11, Y: 12}: {},
				{X: 12, Y: 12}: {},
				{X: 13, Y: 12}: {},
				{X: 14, Y: 12}: {},
				{X: 15, Y: 12}: {walls: types.East | types.South},

				{X: 8, Y: 13}:  {},
				{X: 9, Y: 13}:  {walls: types.East | types.South, target: types.Targets[types.TnBluePyramid]},
				{X: 10, Y: 13}: {walls: types.West},
				{X: 11, Y: 13}: {},
				{X: 12, Y: 13}: {},
				{X: 13, Y: 13}: {},
				{X: 14, Y: 13}: {},
				{X: 15, Y: 13}: {walls: types.East},

				{X: 8, Y: 14}:  {},
				{X: 9, Y: 14}:  {},
				{X: 10, Y: 14}: {},
				{X: 11, Y: 14}: {},
				{X: 12, Y: 14}: {walls: types.East},
				{X: 13, Y: 14}: {walls: types.West | types.North, target: types.Targets[types.TnRedSaturn]},
				{X: 14, Y: 14}: {},
				{X: 15, Y: 14}: {walls: types.East},

				{X: 8, Y: 15}:  {walls: types.North},
				{X: 9, Y: 15}:  {walls: types.North},
				{X: 10, Y: 15}: {walls: types.North},
				{X: 11, Y: 15}: {walls: types.East | types.North},
				{X: 12, Y: 15}: {walls: types.West | types.North},
				{X: 13, Y: 15}: {walls: types.South | types.North},
				{X: 14, Y: 15}: {walls: types.North},
				{X: 15, Y: 15}: {walls: types.East | types.North},

				/*
					{Set: 'A', No: 3, Front: true}
				*/

				{X: 8, Y: 0}:  {walls: types.South},
				{X: 9, Y: 0}:  {walls: types.South},
				{X: 10, Y: 0}: {walls: types.South},
				{X: 11, Y: 0}: {walls: types.South | types.North},
				{X: 12, Y: 0}: {walls: types.South},
				{X: 13, Y: 0}: {walls: types.East | types.South},
				{X: 14, Y: 0}: {walls: types.West | types.South},
				{X: 15, Y: 0}: {walls: types.East | types.South},

				{X: 8, Y: 1}:  {},
				{X: 9, Y: 1}:  {},
				{X: 10, Y: 1}: {},
				{X: 11, Y: 1}: {walls: types.East | types.South, target: types.Targets[types.TnRedMoon]},
				{X: 12, Y: 1}: {walls: types.West},
				{X: 13, Y: 1}: {},
				{X: 14, Y: 1}: {walls: types.North},
				{X: 15, Y: 1}: {walls: types.East},

				{X: 8, Y: 2}:  {},
				{X: 9, Y: 2}:  {},
				{X: 10, Y: 2}: {},
				{X: 11, Y: 2}: {},
				{X: 12, Y: 2}: {},
				{X: 13, Y: 2}: {walls: types.East},
				{X: 14, Y: 2}: {walls: types.West | types.South, target: types.Targets[types.TnGreenPyramid]},
				{X: 15, Y: 2}: {walls: types.East},

				{X: 8, Y: 3}:  {walls: types.East},
				{X: 9, Y: 3}:  {walls: types.West | types.North, target: types.Targets[types.TnYellowStar]},
				{X: 10, Y: 3}: {},
				{X: 11, Y: 3}: {},
				{X: 12, Y: 3}: {},
				{X: 13, Y: 3}: {},
				{X: 14, Y: 3}: {},
				{X: 15, Y: 3}: {walls: types.East},

				{X: 8, Y: 4}:  {},
				{X: 9, Y: 4}:  {walls: types.South},
				{X: 10, Y: 4}: {},
				{X: 11, Y: 4}: {},
				{X: 12, Y: 4}: {},
				{X: 13, Y: 4}: {},
				{X: 14, Y: 4}: {},
				{X: 15, Y: 4}: {walls: types.East},

				{X: 8, Y: 5}:  {},
				{X: 9, Y: 5}:  {},
				{X: 10, Y: 5}: {},
				{X: 11, Y: 5}: {},
				{X: 12, Y: 5}: {},
				{X: 13, Y: 5}: {},
				{X: 14, Y: 5}: {},
				{X: 15, Y: 5}: {walls: types.East | types.North},

				{X: 8, Y: 6}:  {walls: types.North},
				{X: 9, Y: 6}:  {},
				{X: 10, Y: 6}: {},
				{X: 11, Y: 6}: {},
				{X: 12, Y: 6}: {walls: types.East | types.North, target: types.Targets[types.TnBlueSaturn]},
				{X: 13, Y: 6}: {walls: types.West},
				{X: 14, Y: 6}: {},
				{X: 15, Y: 6}: {walls: types.East | types.South},

				{X: 8, Y: 7}:  {walls: types.East | types.South},
				{X: 9, Y: 7}:  {walls: types.West},
				{X: 10, Y: 7}: {},
				{X: 11, Y: 7}: {},
				{X: 12, Y: 7}: {walls: types.South},
				{X: 13, Y: 7}: {},
				{X: 14, Y: 7}: {},
				{X: 15, Y: 7}: {walls: types.East},

				/*
					{Set: 'A', No: 4, Front: true}
				*/

				{X: 0, Y: 0}: {walls: types.West | types.South},
				{X: 1, Y: 0}: {walls: types.South | types.North},
				{X: 2, Y: 0}: {walls: types.South},
				{X: 3, Y: 0}: {walls: types.South},
				{X: 4, Y: 0}: {walls: types.South},
				{X: 5, Y: 0}: {walls: types.East | types.South},
				{X: 6, Y: 0}: {walls: types.West | types.South},
				{X: 7, Y: 0}: {walls: types.South},

				{X: 0, Y: 1}: {walls: types.West},
				{X: 1, Y: 1}: {walls: types.East | types.South, target: types.Targets[types.TnRedPyramid]},
				{X: 2, Y: 1}: {walls: types.West},
				{X: 3, Y: 1}: {},
				{X: 4, Y: 1}: {},
				{X: 5, Y: 1}: {},
				{X: 6, Y: 1}: {},
				{X: 7, Y: 1}: {},

				{X: 0, Y: 2}: {walls: types.West},
				{X: 1, Y: 2}: {},
				{X: 2, Y: 2}: {},
				{X: 3, Y: 2}: {},
				{X: 4, Y: 2}: {walls: types.East | types.North, target: types.Targets[types.TnBlueStar]},
				{X: 5, Y: 2}: {walls: types.West},
				{X: 6, Y: 2}: {},
				{X: 7, Y: 2}: {},

				{X: 0, Y: 3}: {walls: types.West | types.North},
				{X: 1, Y: 3}: {},
				{X: 2, Y: 3}: {},
				{X: 3, Y: 3}: {},
				{X: 4, Y: 3}: {walls: types.South},
				{X: 5, Y: 3}: {},
				{X: 6, Y: 3}: {},
				{X: 7, Y: 3}: {},

				{X: 0, Y: 4}: {walls: types.West | types.South},
				{X: 1, Y: 4}: {},
				{X: 2, Y: 4}: {},
				{X: 3, Y: 4}: {},
				{X: 4, Y: 4}: {},
				{X: 5, Y: 4}: {},
				{X: 6, Y: 4}: {},
				{X: 7, Y: 4}: {},

				{X: 0, Y: 5}: {walls: types.West},
				{X: 1, Y: 5}: {},
				{X: 2, Y: 5}: {},
				{X: 3, Y: 5}: {},
				{X: 4, Y: 5}: {},
				{X: 5, Y: 5}: {},
				{X: 6, Y: 5}: {},
				{X: 7, Y: 5}: {},

				{X: 0, Y: 6}: {walls: types.West},
				{X: 1, Y: 6}: {walls: types.East},
				{X: 2, Y: 6}: {walls: types.West | types.North, target: types.Targets[types.TnGreenMoon]},
				{X: 3, Y: 6}: {},
				{X: 4, Y: 6}: {},
				{X: 5, Y: 6}: {walls: types.North},
				{X: 6, Y: 6}: {},
				{X: 7, Y: 6}: {walls: types.North},

				{X: 0, Y: 7}: {walls: types.West},
				{X: 1, Y: 7}: {},
				{X: 2, Y: 7}: {walls: types.South},
				{X: 3, Y: 7}: {},
				{X: 4, Y: 7}: {walls: types.East},
				{X: 5, Y: 7}: {walls: types.West | types.South, target: types.Targets[types.TnYellowSaturn]},
				{X: 6, Y: 7}: {walls: types.East},
				{X: 7, Y: 7}: {walls: types.West | types.South},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			b := NewBoard(test.tileIDs)
			for p, f1 := range test.fields {
				f2 := b.Field(p.X, p.Y)
				// if !reflect.DeepEqual(f1, f2) { // no moves check yet
				if f1.walls != f2.walls || f1.target != f2.target {
					t.Fatalf("field %d,%d: %v - expected %v", p.X, p.Y, f2, f1)
				}
			}
		})
	}
}
