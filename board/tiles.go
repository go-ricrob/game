package board

import . "github.com/go-ricrob/game/types"

// tiles orientation: 'center' field position x,y = 0,0
var tiles = map[string]map[Coordinate]Field{
	"A1F": {
		{X: 0, Y: 7}: {walls: East},
		{X: 2, Y: 0}: {walls: South | West, symbol: Cosmic, color: CosmicColor},
		{X: 2, Y: 5}: {walls: South | West, symbol: Star, color: Red},
		{X: 3, Y: 2}: {walls: North | East, symbol: Saturn, color: Green},
		{X: 4, Y: 6}: {walls: South | East, symbol: Pyramid, color: Yellow},
		{X: 6, Y: 1}: {walls: North | West, symbol: Moon, color: Blue},
		{X: 7, Y: 3}: {walls: North},
	},
	"A1B": {
		{X: 0, Y: 4}: {walls: South | West, symbol: Cosmic, color: CosmicColor},
		{X: 1, Y: 2}: {walls: South | East, symbol: Moon, color: Blue},
		{X: 2, Y: 7}: {walls: East},
		{X: 3, Y: 1}: {walls: North | East, symbol: Pyramid, color: Yellow},
		{X: 4, Y: 6}: {walls: North | West, symbol: Saturn, color: Green},
		{X: 6, Y: 5}: {walls: South | West, symbol: Star, color: Red},
		{X: 7, Y: 2}: {walls: North},
	},
	"A2F": {
		{X: 1, Y: 5}: {walls: South | East, symbol: Pyramid, color: Blue},
		{X: 3, Y: 1}: {walls: North | East, symbol: Moon, color: Yellow},
		{X: 3, Y: 7}: {walls: East},
		{X: 5, Y: 6}: {walls: North | West, symbol: Saturn, color: Red},
		{X: 6, Y: 2}: {walls: South | West, symbol: Star, color: Green},
		{X: 7, Y: 3}: {walls: North},
	},
	"A2B": {
		{X: 1, Y: 6}: {walls: South | West, symbol: Star, color: Green},
		{X: 2, Y: 3}: {walls: South | East, symbol: Saturn, color: Red},
		{X: 2, Y: 7}: {walls: East},
		{X: 4, Y: 1}: {walls: North | West, symbol: Pyramid, color: Blue},
		{X: 6, Y: 5}: {walls: North | East, symbol: Moon, color: Yellow},
		{X: 7, Y: 2}: {walls: North},
	},
	"A3F": {
		{X: 1, Y: 4}: {walls: North | West, symbol: Saturn, color: Blue},
		{X: 1, Y: 7}: {walls: East},
		{X: 4, Y: 1}: {walls: South | West, symbol: Star, color: Yellow},
		{X: 5, Y: 6}: {walls: South | East, symbol: Pyramid, color: Green},
		{X: 6, Y: 3}: {walls: North | East, symbol: Moon, color: Red},
		{X: 7, Y: 5}: {walls: North},
	},
	"A3B": {
		{X: 1, Y: 4}: {walls: North | West, symbol: Moon, color: Red},
		{X: 1, Y: 7}: {walls: East},
		{X: 2, Y: 1}: {walls: South | East, symbol: Saturn, color: Blue},
		{X: 5, Y: 6}: {walls: North | East, symbol: Pyramid, color: Green},
		{X: 6, Y: 1}: {walls: South | West, symbol: Star, color: Yellow},
		{X: 7, Y: 2}: {walls: North},
	},
	"A4F": {
		{X: 1, Y: 7}: {walls: East},
		{X: 2, Y: 0}: {walls: North | East, symbol: Saturn, color: Yellow},
		{X: 3, Y: 5}: {walls: South | West, symbol: Star, color: Blue},
		{X: 5, Y: 1}: {walls: South | East, symbol: Moon, color: Green},
		{X: 6, Y: 6}: {walls: North | West, symbol: Pyramid, color: Red},
		{X: 7, Y: 3}: {walls: North},
	},
	"A4B": {
		{X: 1, Y: 6}: {walls: North | East, symbol: Saturn, color: Yellow},
		{X: 2, Y: 0}: {walls: North | West, symbol: Pyramid, color: Red},
		{X: 2, Y: 7}: {walls: East},
		{X: 3, Y: 5}: {walls: South | East, symbol: Moon, color: Green},
		{X: 5, Y: 2}: {walls: South | West, symbol: Star, color: Blue},
		{X: 7, Y: 3}: {walls: North},
	},
}
