package board

import "github.com/go-ricrob/game/types"

// tile fields orientation: 'center' field position x,y = 0,0
var tileFields = map[types.TileID]map[types.Coordinate]Field{
	{SetID: 'A', TileNo: 1, Front: true}: {
		{X: 0, Y: 7}: {walls: types.East},
		{X: 2, Y: 0}: {walls: types.South | types.West, symbol: types.Cosmic, color: types.CosmicColor},
		{X: 2, Y: 5}: {walls: types.South | types.West, symbol: types.Star, color: types.Red},
		{X: 3, Y: 2}: {walls: types.North | types.East, symbol: types.Saturn, color: types.Green},
		{X: 4, Y: 6}: {walls: types.South | types.East, symbol: types.Pyramid, color: types.Yellow},
		{X: 6, Y: 1}: {walls: types.North | types.West, symbol: types.Moon, color: types.Blue},
		{X: 7, Y: 3}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 1, Front: false}: {
		{X: 0, Y: 4}: {walls: types.South | types.West, symbol: types.Cosmic, color: types.CosmicColor},
		{X: 1, Y: 2}: {walls: types.South | types.East, symbol: types.Moon, color: types.Blue},
		{X: 2, Y: 7}: {walls: types.East},
		{X: 3, Y: 1}: {walls: types.North | types.East, symbol: types.Pyramid, color: types.Yellow},
		{X: 4, Y: 6}: {walls: types.North | types.West, symbol: types.Saturn, color: types.Green},
		{X: 6, Y: 5}: {walls: types.South | types.West, symbol: types.Star, color: types.Red},
		{X: 7, Y: 2}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 2, Front: true}: {
		{X: 1, Y: 5}: {walls: types.South | types.East, symbol: types.Pyramid, color: types.Blue},
		{X: 3, Y: 1}: {walls: types.North | types.East, symbol: types.Moon, color: types.Yellow},
		{X: 3, Y: 7}: {walls: types.East},
		{X: 5, Y: 6}: {walls: types.North | types.West, symbol: types.Saturn, color: types.Red},
		{X: 6, Y: 2}: {walls: types.South | types.West, symbol: types.Star, color: types.Green},
		{X: 7, Y: 3}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 2, Front: false}: {
		{X: 1, Y: 6}: {walls: types.South | types.West, symbol: types.Star, color: types.Green},
		{X: 2, Y: 3}: {walls: types.South | types.East, symbol: types.Saturn, color: types.Red},
		{X: 2, Y: 7}: {walls: types.East},
		{X: 4, Y: 1}: {walls: types.North | types.West, symbol: types.Pyramid, color: types.Blue},
		{X: 6, Y: 5}: {walls: types.North | types.East, symbol: types.Moon, color: types.Yellow},
		{X: 7, Y: 2}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 3, Front: true}: {
		{X: 1, Y: 4}: {walls: types.North | types.West, symbol: types.Saturn, color: types.Blue},
		{X: 1, Y: 7}: {walls: types.East},
		{X: 4, Y: 1}: {walls: types.South | types.West, symbol: types.Star, color: types.Yellow},
		{X: 5, Y: 6}: {walls: types.South | types.East, symbol: types.Pyramid, color: types.Green},
		{X: 6, Y: 3}: {walls: types.North | types.East, symbol: types.Moon, color: types.Red},
		{X: 7, Y: 5}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 3, Front: false}: {
		{X: 1, Y: 4}: {walls: types.North | types.West, symbol: types.Moon, color: types.Red},
		{X: 1, Y: 7}: {walls: types.East},
		{X: 2, Y: 1}: {walls: types.South | types.East, symbol: types.Saturn, color: types.Blue},
		{X: 5, Y: 6}: {walls: types.North | types.East, symbol: types.Pyramid, color: types.Green},
		{X: 6, Y: 1}: {walls: types.South | types.West, symbol: types.Star, color: types.Yellow},
		{X: 7, Y: 2}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 4, Front: true}: {
		{X: 1, Y: 7}: {walls: types.East},
		{X: 2, Y: 0}: {walls: types.North | types.East, symbol: types.Saturn, color: types.Yellow},
		{X: 3, Y: 5}: {walls: types.South | types.West, symbol: types.Star, color: types.Blue},
		{X: 5, Y: 1}: {walls: types.South | types.East, symbol: types.Moon, color: types.Green},
		{X: 6, Y: 6}: {walls: types.North | types.West, symbol: types.Pyramid, color: types.Red},
		{X: 7, Y: 3}: {walls: types.North},
	},
	{SetID: 'A', TileNo: 4, Front: false}: {
		{X: 1, Y: 6}: {walls: types.North | types.East, symbol: types.Saturn, color: types.Yellow},
		{X: 2, Y: 0}: {walls: types.North | types.West, symbol: types.Pyramid, color: types.Red},
		{X: 2, Y: 7}: {walls: types.East},
		{X: 3, Y: 5}: {walls: types.South | types.East, symbol: types.Moon, color: types.Green},
		{X: 5, Y: 2}: {walls: types.South | types.West, symbol: types.Star, color: types.Blue},
		{X: 7, Y: 3}: {walls: types.North},
	},
}
