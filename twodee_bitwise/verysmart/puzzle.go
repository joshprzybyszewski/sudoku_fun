package verysmart

import (
	"errors"

	"github.com/joshprzybyszewski/sudoku_fun/twodee_bitwise/common"
	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	speedUtils "github.com/joshprzybyszewski/sudoku_fun/utils/speed"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

func getLocationAndEntriesVerySmart(p *common.SmartPuzzle) (row, col, box int, entries []types.Entry, err error) {
	bestRow, bestCol, bestBox := -1, -1, -1
	bestFreeSpots := constants.SideLen + 1
	var bestSpotEntries []types.Entry

	for r := range p.Tiles {
		for c := range p.Tiles[r] {
			if p.Tiles[r][c] != constants.EmptyTile {
				continue
			}

			b, _ := speedUtils.GetBox(r, c)
			freeSpots, entries := speedUtils.GetCachedNumFreeAndPosEntries(p.Rows[r] | p.Cols[c] | p.Boxs[b])

			if freeSpots == 0 {
				return -1, -1, -1, nil, errors.New(`there are no possible entries here!`)
			} else if freeSpots == 1 {
				return r, c, b, entries, nil
			} else if freeSpots < bestFreeSpots {
				bestFreeSpots = freeSpots
				bestSpotEntries = entries

				bestRow = r
				bestCol = c
				bestBox = b
			}
		}
	}

	return bestRow, bestCol, bestBox, bestSpotEntries, nil
}

func Solve(p *common.SmartPuzzle) (solution *common.SmartPuzzle, err error) {
	row, col, box, entries, err := getLocationAndEntriesVerySmart(p)

	if err != nil {
		// Unsolvable!
		return nil, err
	}

	for _, entry := range entries {
		pClone := p.Clone()

		wasPlaced, _ := pClone.Place(row, col, box, entry)

		if wasPlaced {
			if pClone.NumPlaced == constants.NumTiles {
				return pClone, nil
			}

			cloneSolution, _ := Solve(pClone)

			if cloneSolution != nil {
				return cloneSolution, nil
			}
		}
	}

	return nil, errors.New(`Should never be here`)
}
