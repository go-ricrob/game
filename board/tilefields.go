package board

import "github.com/go-ricrob/game/types"

// tile fields orientation: 'center' field position x,y = 0,0
var tileFields = map[types.TileID]map[types.Coordinate]Field{
	{SetID: 'A', TileNo: 1, Front: true}: {
		{X: 0, Y: 7}: {walls: types.East},
		{X: 2, Y: 0}: {walls: types.South | types.West, target: types.Targets[types.TnCosmic]},
		{X: 2, Y: 5}: {walls: types.South | types.West, target: types.Targets[types.TnRedStar]},
		{X: 3, Y: 2}: {walls: types.North | types.East, target: types.Targets[types.TnGreenSaturn]},
		{X: 4, Y: 6}: {walls: types.South | types.East, target: types.Targets[types.TnYellowPyramid]},
		{X: 6, Y: 1}: {walls: types.North | types.West, target: types.Targets[types.TnBlueMoon]},
		{X: 7, Y: 3}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 1, Front: false}: {
		{X: 0, Y: 4}: {walls: types.South | types.West, target: types.Targets[types.TnCosmic]},
		{X: 1, Y: 2}: {walls: types.South | types.East, target: types.Targets[types.TnBlueMoon]},
		{X: 2, Y: 7}: {walls: types.East},
		{X: 3, Y: 1}: {walls: types.North | types.East, target: types.Targets[types.TnYellowPyramid]},
		{X: 4, Y: 6}: {walls: types.North | types.West, target: types.Targets[types.TnGreenSaturn]},
		{X: 6, Y: 5}: {walls: types.South | types.West, target: types.Targets[types.TnRedStar]},
		{X: 7, Y: 2}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 2, Front: true}: {
		{X: 1, Y: 5}: {walls: types.South | types.East, target: types.Targets[types.TnBluePyramid]},
		{X: 3, Y: 1}: {walls: types.North | types.East, target: types.Targets[types.TnYellowMoon]},
		{X: 3, Y: 7}: {walls: types.East},
		{X: 5, Y: 6}: {walls: types.North | types.West, target: types.Targets[types.TnRedSaturn]},
		{X: 6, Y: 2}: {walls: types.South | types.West, target: types.Targets[types.TnGreenStar]},
		{X: 7, Y: 3}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 2, Front: false}: {
		{X: 1, Y: 6}: {walls: types.South | types.West, target: types.Targets[types.TnGreenStar]},
		{X: 2, Y: 3}: {walls: types.South | types.East, target: types.Targets[types.TnRedSaturn]},
		{X: 2, Y: 7}: {walls: types.East},
		{X: 4, Y: 1}: {walls: types.North | types.West, target: types.Targets[types.TnBluePyramid]},
		{X: 6, Y: 5}: {walls: types.North | types.East, target: types.Targets[types.TnYellowMoon]},
		{X: 7, Y: 2}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 3, Front: true}: {
		{X: 1, Y: 4}: {walls: types.North | types.West, target: types.Targets[types.TnBlueSaturn]},
		{X: 1, Y: 7}: {walls: types.East},
		{X: 4, Y: 1}: {walls: types.South | types.West, target: types.Targets[types.TnYellowStar]},
		{X: 5, Y: 6}: {walls: types.South | types.East, target: types.Targets[types.TnGreenPyramid]},
		{X: 6, Y: 3}: {walls: types.North | types.East, target: types.Targets[types.TnRedMoon]},
		{X: 7, Y: 5}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 3, Front: false}: {
		{X: 1, Y: 4}: {walls: types.North | types.West, target: types.Targets[types.TnRedMoon]},
		{X: 1, Y: 7}: {walls: types.East},
		{X: 2, Y: 1}: {walls: types.South | types.East, target: types.Targets[types.TnBlueSaturn]},
		{X: 5, Y: 6}: {walls: types.North | types.East, target: types.Targets[types.TnGreenPyramid]},
		{X: 6, Y: 1}: {walls: types.South | types.West, target: types.Targets[types.TnYellowStar]},
		{X: 7, Y: 2}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 4, Front: true}: {
		{X: 1, Y: 7}: {walls: types.East},
		{X: 2, Y: 0}: {walls: types.North | types.East, target: types.Targets[types.TnYellowSaturn]},
		{X: 3, Y: 5}: {walls: types.South | types.West, target: types.Targets[types.TnBlueStar]},
		{X: 5, Y: 1}: {walls: types.South | types.East, target: types.Targets[types.TnGreenMoon]},
		{X: 6, Y: 6}: {walls: types.North | types.West, target: types.Targets[types.TnRedPyramid]},
		{X: 7, Y: 3}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 4, Front: false}: {
		{X: 1, Y: 6}: {walls: types.North | types.East, target: types.Targets[types.TnYellowSaturn]},
		{X: 2, Y: 0}: {walls: types.North | types.West, target: types.Targets[types.TnRedPyramid]},
		{X: 2, Y: 7}: {walls: types.East},
		{X: 3, Y: 5}: {walls: types.South | types.East, target: types.Targets[types.TnGreenMoon]},
		{X: 5, Y: 2}: {walls: types.South | types.West, target: types.Targets[types.TnBlueStar]},
		{X: 7, Y: 3}: {walls: types.North},
	},
}
