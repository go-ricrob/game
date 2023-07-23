package board

import "fmt"

// Symbol is the type of a symbol.
type Symbol byte

// Symbol constants.
const (
	NoSymbol = iota
	Pyramid
	Star
	Moon
	Saturn
	Cosmic
)

var symbolStrs = []string{"", "Pyramid", "Star", "Moon", "Saturn", "Cosmic"}

func (s Symbol) String() string {
	if int(s) >= len(symbolStrs) {
		return fmt.Sprintf("invalid symbol %d", s)
	}
	return symbolStrs[s]
}

// Symbols is the set of valid symbols.
var Symbols = []Symbol{Pyramid, Star, Moon, Saturn, Cosmic}
