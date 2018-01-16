package slow

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

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

func PresenceOf(entry types.Entry) types.Presence {
	if entry <= 0 || entry > constants.SideLen {
		return constants.PresenceError
	}

	return types.Presence(1 << uint(entry-1))
}

func IsPresent(existing, entryPresence types.Presence) bool {
	return existing&entryPresence > 0
}

func GetBox(row, col int) (int, error) {
	if row >= constants.SideLen || row < 0 ||
		col >= constants.SideLen || col < 0 {
		return -1, errors.New(`Row or col was out of bounds!`)
	}

	// 0 | 1 | 2
	// 3 | 4 | 5
	// 6 | 7 | 8

	box := ((row / constants.Root) * constants.Root) + (col / constants.Root)

	return box, nil
}

func GetNumFreeSpots(presence types.Presence) int {
	numFree := 0

	for i := 1; i <= constants.SideLen; i++ {
		ePresence := PresenceOf(types.Entry(i))
		if IsPresent(presence, ePresence) {
			numFree += 1
		}
	}

	return numFree
}

func PrintError(str string, err error) {
	println(fmt.Sprintf("Error @ %v: %v", str, err.Error()))
}

func PrintPlacement(e types.Entry, r, c int) {
	println(fmt.Sprintf("Attempting to place %v at (%v, %v)", e, r, c))
}
