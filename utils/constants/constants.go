package constants

import "../types"

const (
	Root     = 3
	SideLen  = Root * Root
	NumTiles = SideLen * SideLen

	PresenceError = EmptyPresence
	EmptyPresence = types.Presence(0)
	FullPresence  = types.Presence(511) // calc'd cuz first 9 bits = 2^9 - 1

	EmptyTile     = types.Tile(0)
	EmptyTileChar = `.`

	onePresence   = types.Presence(1 << uint(1-1))
	twoPresence   = types.Presence(1 << uint(2-1))
	threePresence = types.Presence(1 << uint(3-1))
	fourPresence  = types.Presence(1 << uint(4-1))
	fivePresence  = types.Presence(1 << uint(5-1))
	sixPresence   = types.Presence(1 << uint(6-1))
	sevenPresence = types.Presence(1 << uint(7-1))
	eightPresence = types.Presence(1 << uint(8-1))
	ninePresence  = types.Presence(1 << uint(9-1))

	EmptyPuzzle = `.................................................................................`
)

var (
	AllEntries       = []types.Entry{1, 2, 3, 4, 5, 6, 7, 8, 9}

	allEntriesPresence = []types.Presence{
		onePresence,
		twoPresence,
		threePresence,
		fourPresence,
		fivePresence,
		sixPresence,
		sevenPresence,
		eightPresence,
		ninePresence,
	}
)
