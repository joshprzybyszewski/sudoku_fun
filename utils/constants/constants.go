package constants

import (
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

const (
	Root     = 3
	SideLen  = Root * Root
	NumTiles = SideLen * SideLen

	PresenceError = EmptyPresence
	EmptyPresence = types.Presence(0)
	FullPresence  = types.Presence(511) // calc'd cuz first 9 bits = 2^9 - 1

	EmptyTile     = types.Tile(0)
	EmptyTileChar = `.`

	EmptyPuzzle = `.................................................................................`
)

var (
	AllEntries = []types.Entry{1, 2, 3, 4, 5, 6, 7, 8, 9}
)
