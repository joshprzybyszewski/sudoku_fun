package utils

import (
	"strconv"

	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
)

func BruteForceCheck(sudokuStr string) bool {
	var entries [constants.SideLen][constants.SideLen]int
	for i := 0; i < len(sudokuStr); i++ {
		entryChar := string(sudokuStr[i])
		entry, err := strconv.Atoi(entryChar)

		if entryChar == constants.EmptyTileChar || err != nil {
			return false
		}

		row := i / constants.SideLen
		col := i % constants.SideLen

		entries[row][col] = entry
	}

	for r := range entries {
		for c := range entries[r] {
			value := entries[r][c]
			baseRow := (r / constants.Root) * constants.Root
			baseCol := (c / constants.Root) * constants.Root

			for i := 0; i < constants.SideLen; i++ {
				if i != r && entries[i][c] == value {
					return false
				}

				if i != c && entries[r][i] == value {
					return false
				}

				boxEntryR := baseRow + (i / constants.Root)
				boxEntryC := baseCol + (i % constants.Root)

				if boxEntryR != r && boxEntryC != c && entries[boxEntryR][boxEntryC] == value {
					return false
				}
			}
		}
	}

	return true
}
