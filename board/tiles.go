package board

import (
	"github.com/go-ricrob/game/coord"
	"github.com/go-ricrob/game/types"
)

// tiles orientation: 'center' field position x,y = 0,0
var tiles = map[string]map[byte]Field{
	"A1F": {
		coord.Ctob(0, 7): {Walls: EastWall},
		coord.Ctob(2, 0): {Walls: SouthWall | WestWall, Symbol: Cosmic},
		coord.Ctob(2, 5): {Walls: SouthWall | WestWall, Symbol: Star, Color: types.Red},
		coord.Ctob(3, 2): {Walls: NorthWall | EastWall, Symbol: Saturn, Color: types.Green},
		coord.Ctob(4, 6): {Walls: SouthWall | EastWall, Symbol: Pyramid, Color: types.Yellow},
		coord.Ctob(6, 1): {Walls: NorthWall | WestWall, Symbol: Moon, Color: types.Blue},
		coord.Ctob(7, 3): {Walls: NorthWall},
	},
	"A1B": {
		coord.Ctob(0, 4): {Walls: SouthWall | WestWall, Symbol: Cosmic},
		coord.Ctob(1, 2): {Walls: SouthWall | EastWall, Symbol: Moon, Color: types.Blue},
		coord.Ctob(2, 7): {Walls: EastWall},
		coord.Ctob(3, 1): {Walls: NorthWall | EastWall, Symbol: Pyramid, Color: types.Yellow},
		coord.Ctob(4, 6): {Walls: NorthWall | WestWall, Symbol: Saturn, Color: types.Green},
		coord.Ctob(6, 5): {Walls: SouthWall | WestWall, Symbol: Star, Color: types.Red},
		coord.Ctob(7, 2): {Walls: NorthWall},
	},
	"A2F": {
		coord.Ctob(1, 5): {Walls: SouthWall | EastWall, Symbol: Pyramid, Color: types.Blue},
		coord.Ctob(3, 1): {Walls: NorthWall | EastWall, Symbol: Moon, Color: types.Yellow},
		coord.Ctob(3, 7): {Walls: EastWall},
		coord.Ctob(5, 6): {Walls: NorthWall | WestWall, Symbol: Saturn, Color: types.Red},
		coord.Ctob(6, 2): {Walls: SouthWall | WestWall, Symbol: Star, Color: types.Green},
		coord.Ctob(7, 3): {Walls: NorthWall},
	},
	"A2B": {
		coord.Ctob(1, 6): {Walls: SouthWall | WestWall, Symbol: Star, Color: types.Green},
		coord.Ctob(2, 3): {Walls: SouthWall | EastWall, Symbol: Saturn, Color: types.Red},
		coord.Ctob(2, 7): {Walls: EastWall},
		coord.Ctob(4, 1): {Walls: NorthWall | WestWall, Symbol: Pyramid, Color: types.Blue},
		coord.Ctob(6, 5): {Walls: NorthWall | EastWall, Symbol: Moon, Color: types.Yellow},
		coord.Ctob(7, 2): {Walls: NorthWall},
	},
	"A3F": {
		coord.Ctob(1, 4): {Walls: NorthWall | WestWall, Symbol: Saturn, Color: types.Blue},
		coord.Ctob(1, 7): {Walls: EastWall},
		coord.Ctob(4, 1): {Walls: SouthWall | WestWall, Symbol: Star, Color: types.Yellow},
		coord.Ctob(5, 6): {Walls: SouthWall | EastWall, Symbol: Pyramid, Color: types.Green},
		coord.Ctob(6, 3): {Walls: NorthWall | EastWall, Symbol: Moon, Color: types.Red},
		coord.Ctob(7, 5): {Walls: NorthWall},
	},
	"A3B": {
		coord.Ctob(1, 4): {Walls: NorthWall | WestWall, Symbol: Moon, Color: types.Red},
		coord.Ctob(1, 7): {Walls: EastWall},
		coord.Ctob(2, 1): {Walls: SouthWall | EastWall, Symbol: Saturn, Color: types.Blue},
		coord.Ctob(5, 6): {Walls: NorthWall | EastWall, Symbol: Pyramid, Color: types.Green},
		coord.Ctob(6, 1): {Walls: SouthWall | WestWall, Symbol: Star, Color: types.Yellow},
		coord.Ctob(7, 2): {Walls: NorthWall},
	},
	"A4F": {
		coord.Ctob(1, 7): {Walls: EastWall},
		coord.Ctob(2, 0): {Walls: NorthWall | EastWall, Symbol: Saturn, Color: types.Yellow},
		coord.Ctob(3, 5): {Walls: SouthWall | WestWall, Symbol: Star, Color: types.Blue},
		coord.Ctob(5, 1): {Walls: SouthWall | EastWall, Symbol: Moon, Color: types.Green},
		coord.Ctob(6, 6): {Walls: NorthWall | WestWall, Symbol: Pyramid, Color: types.Red},
		coord.Ctob(7, 3): {Walls: NorthWall},
	},
	"A4B": {
		coord.Ctob(1, 6): {Walls: NorthWall | EastWall, Symbol: Saturn, Color: types.Yellow},
		coord.Ctob(2, 0): {Walls: NorthWall | WestWall, Symbol: Pyramid, Color: types.Red},
		coord.Ctob(2, 7): {Walls: EastWall},
		coord.Ctob(3, 5): {Walls: SouthWall | EastWall, Symbol: Moon, Color: types.Green},
		coord.Ctob(5, 2): {Walls: SouthWall | WestWall, Symbol: Star, Color: types.Blue},
		coord.Ctob(7, 3): {Walls: NorthWall},
	},
}
