package smart

import (
	"errors"

	"github.com/joshprzybyszewski/sudoku_fun/twodee_bitwise/common"
	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	speedUtils "github.com/joshprzybyszewski/sudoku_fun/utils/speed"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

func getLocationAndEntries(p *common.SmartPuzzle) (row, col, box int, entries []types.Entry, err error) {
	row, col, box, _ = getEmptyTile(p)

	entries, err = getEntries(p, row, col, box)
	if err != nil {
		return -1, -1, -1, nil, err
	}

	return row, col, box, entries, nil
}

func getEmptyTile(p *common.SmartPuzzle) (row, col, box int, err error) {
	bkRow := types.NewBK()
	bkCol := types.NewBK()
	bkBox := types.NewBK()

	for i := 0; i < constants.SideLen; i++ {
		bkRow.Update(i, p.NumFreeInRow[i])
		bkCol.Update(i, p.NumFreeInCol[i])
	}

	if bkRow.IsError() || bkCol.IsError() {
		return -1, -1, -1, errors.New(`All the rows or cols are full!`)
	}

	if bkRow.IsBetterThan(bkCol) {
		getBestCol(p, bkRow.Location(), bkCol, bkBox)
	} else {
		getBestRow(p, bkCol.Location(), bkRow, bkBox)
	}

	return bkRow.Location(), bkCol.Location(), bkBox.Location(), nil
}
func getBestCol(p *common.SmartPuzzle, bestRowIndex int, bkCol, bkBox *types.BestKnown) {
	bkCol.Reset()

	for c := range p.NumFreeInCol {
		if p.Tiles[bestRowIndex][c] != constants.EmptyTile {
			continue
		}

		if !bkCol.WillUpdate(c, p.NumFreeInCol[c]) {
			continue
		}

		b, _ := speedUtils.GetBox(bestRowIndex, c)

		if p.NumFreeInBox[b] == 0 {
			continue
		}

		bkCol.UpdateUnsafe(c, p.NumFreeInCol[c])
		bkBox.UpdateUnsafe(b, p.NumFreeInBox[b])
	}
}
func getBestRow(p *common.SmartPuzzle, bestColIndex int, bkRow, bkBox *types.BestKnown) {
	bkRow.Reset()

	for r := range p.NumFreeInRow {
		if p.Tiles[r][bestColIndex] != constants.EmptyTile {
			continue
		}

		if !bkRow.WillUpdate(r, p.NumFreeInRow[r]) {
			continue
		}

		b, _ := speedUtils.GetBox(r, bestColIndex)

		if p.NumFreeInBox[b] == 0 {
			continue
		}

		bkRow.UpdateUnsafe(r, p.NumFreeInRow[r])
		bkBox.UpdateUnsafe(b, p.NumFreeInBox[b])
	}
}

func getEntries(p *common.SmartPuzzle, row, col, box int) ([]types.Entry, error) {
	entries := speedUtils.GetPossibleEntriesQuickly(p.Rows[row], p.Cols[col], p.Boxs[box])

	if len(entries) == 0 {
		return nil, errors.New(`no entries possible`)
	}

	return entries, nil
}

func Solve(p *common.SmartPuzzle) (solution *common.SmartPuzzle, err error) {
	row, col, box, entries, err := getLocationAndEntries(p)

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
