package speed

import (
	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

const (
	onePresence   = types.Presence(1 << uint(1-1))
	twoPresence   = types.Presence(1 << uint(2-1))
	threePresence = types.Presence(1 << uint(3-1))
	fourPresence  = types.Presence(1 << uint(4-1))
	fivePresence  = types.Presence(1 << uint(5-1))
	sixPresence   = types.Presence(1 << uint(6-1))
	sevenPresence = types.Presence(1 << uint(7-1))
	eightPresence = types.Presence(1 << uint(8-1))
	ninePresence  = types.Presence(1 << uint(9-1))
)

var (
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

	presenceToNumFree [constants.FullPresence + 1]numFreeAndPossibleEntries
	inited = false
)

type numFreeAndPossibleEntries struct{
	numFree uint8
	possibleEntries []types.Entry
}

func InitUtils() {
	if inited {
		return
	}
	for p := types.Presence(0); p <= constants.FullPresence; p++ {
		presenceToNumFree[p] = numFreeAndPossibleEntries{
			uint8(GetNumFreeSpots(p)),
			GetPossibleEntriesQuickly(p, constants.EmptyPresence, constants.EmptyPresence),
		}
	}
	inited = true
}

func GetCachedNumFreeAndPosEntries(presence types.Presence) (int, []types.Entry) {
	if !inited {
		InitUtils()
	}
	value := presenceToNumFree[presence]
	return int(value.numFree), value.possibleEntries
}

func GetPossibleEntries(entries []types.Entry, presence types.Presence) []types.Entry {
	posEntries := make([]types.Entry, 0, constants.SideLen)
	for _, entry := range entries {
		ePresence := PresenceOf(entry)
		if !IsPresent(presence, ePresence) {
			posEntries = append(posEntries, entry)
		}
	}
	return posEntries
}

func GetPossibleEntriesQuickly(rowP, colP, boxP types.Presence) []types.Entry {
	posEntries := make([]types.Entry, 0, constants.SideLen)
	presence := rowP | colP | boxP

	if !IsPresent(presence, onePresence) {
		posEntries = append(posEntries, types.Entry(1))
	}
	if !IsPresent(presence, twoPresence) {
		posEntries = append(posEntries, types.Entry(2))
	}
	if !IsPresent(presence, threePresence) {
		posEntries = append(posEntries, types.Entry(3))
	}
	if !IsPresent(presence, fourPresence) {
		posEntries = append(posEntries, types.Entry(4))
	}
	if !IsPresent(presence, fivePresence) {
		posEntries = append(posEntries, types.Entry(5))
	}
	if !IsPresent(presence, sixPresence) {
		posEntries = append(posEntries, types.Entry(6))
	}
	if !IsPresent(presence, sevenPresence) {
		posEntries = append(posEntries, types.Entry(7))
	}
	if !IsPresent(presence, eightPresence) {
		posEntries = append(posEntries, types.Entry(8))
	}
	if !IsPresent(presence, ninePresence) {
		posEntries = append(posEntries, types.Entry(9))
	}

	return posEntries
}

func PresenceOf(entry types.Entry) types.Presence {
	return presenceOfSpeed(entry)
}
func presenceOfSpeed(entry types.Entry) types.Presence {
	return types.Presence(1 << uint(entry-1))
}

func IsPresent(existing, entryPresence types.Presence) bool {
	return existing&entryPresence > 0
}

func GetBox(row, col int) (int, error) {
	return getBoxSpeed(row, col)
}
func getBoxSpeed(row, col int) (int, error) {
	return ((row / constants.Root) * constants.Root) + (col / constants.Root), nil
}

func GetNumFreeSpots(presence types.Presence) int {
	numFree := 0

	for _, ePresence := range allEntriesPresence {
		if !IsPresent(presence, ePresence) {
			numFree++
		}
	}

	return numFree
}
