package smart

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"../../utils/constants"
	utils "../../utils/speed"
	"../../utils/types"
)

var numPlacements int

type Puzzle struct {
	numPlaced int

	tiles [constants.SideLen][constants.SideLen]types.Tile /* [row][col] */

	rows [constants.SideLen]types.Presence
	cols [constants.SideLen]types.Presence
	boxs [constants.SideLen]types.Presence

	numFreeInRow [constants.SideLen]uint8
	numFreeInCol [constants.SideLen]uint8
	numFreeInBox [constants.SideLen]uint8
}

func newPuzzle() Puzzle {
	pzl := Puzzle{}

	for i := 0; i < constants.SideLen; i++ {
		pzl.numFreeInRow[i] = constants.SideLen
		pzl.numFreeInCol[i] = constants.SideLen
		pzl.numFreeInBox[i] = constants.SideLen
	}

	return pzl
}

/// Reads in sudoku from a string representation of it
func ReadSudoku(str string) (p Puzzle, err error) {
	numPlacements = 0

	pzl := newPuzzle()
	for i := 0; i < len(str); i++ {
		entryChar := string(str[i])
		entry, err := strconv.Atoi(entryChar)

		if entryChar == constants.EmptyTileChar || err != nil {
			continue
		}

		row := i / constants.SideLen
		col := i % constants.SideLen
		box, err := utils.GetBox(row, col)
		if wasPlaced, err := pzl.place(row, col, box, types.Entry(entry)); err != nil || !wasPlaced {
			return Puzzle{}, errors.New(`BAD sudoku!!!`)
		}
	}

	return pzl, nil
}

func (p Puzzle) GetNumPlacements() int {
	return numPlacements
}
func (p Puzzle) GetSimple() string {
	line := ``

	for r := range p.tiles {
		for c := range p.tiles[r] {
			if p.tiles[r][c] == constants.EmptyTile {
				line += constants.EmptyTileChar
			} else {
				line += strconv.Itoa(int(p.tiles[r][c]))
			}
		}
	}

	return line
}
func (p Puzzle) PrintSimple() {
	println(p.GetSimple())
}
func (p Puzzle) PrintPretty() {
	for r := range p.tiles {
		line := ``

		for c := range p.tiles[r] {
			line += ` `

			if p.tiles[r][c] == constants.EmptyTile {
				line += constants.EmptyTileChar
			} else {
				line += strconv.Itoa(int(p.tiles[r][c]))
			}

			line += ` `

			if c == 2 || c == 5 {
				line += ` | `
			}
		}

		println(line)

		if r == 2 || r == 5 {
			println(`-+--+--+--|--+--+--+--|--+--+--+-`)
		}
	}
}
func (p Puzzle) Solve() (solution types.Sudoku, err error) {
	numPlacements = 0

	return p.solve()
}

func (p *Puzzle) clone() *Puzzle {
	clone := Puzzle{}

	clone = *p

	return &clone
}

func (p *Puzzle) getLocationAndEntries() (row, col, box int, entries []types.Entry, err error) {
	row, col, box, _ = p.getEmptyTile()
	//if err != nil {
	//	return -1, -1, -1, nil, errors.New(`no empty tile found`) //errors.New(fmt.Sprintf("There are no possible entries for location (%v, %v)", row, col))
	//}

	entries, err = p.getEntries(row, col, box)
	if err != nil {
		//p.PrintPretty()
		return -1, -1, -1, nil, errors.New(`no entries possible`) //errors.New(fmt.Sprintf("There are no possible entries for location (%v, %v)", row, col))
	}

	return row, col, box, entries, nil
}

func (p *Puzzle) getEntries(row, col, box int) ([]types.Entry, error) {
	//return p.stupidGetEntries(row, col, box)
	return p.quickGetEntries(row, col, box)
}
func (p *Puzzle) stupidGetEntries(row, col, box int) ([]types.Entry, error) {
	entries := constants.AllEntries
	entries = utils.GetPossibleEntries(entries, p.rows[row])
	entries = utils.GetPossibleEntries(entries, p.cols[col])
	entries = utils.GetPossibleEntries(entries, p.boxs[box])

	if len(entries) == 0 {
		//p.PrintPretty()
		return nil, errors.New(`no entries possible`) //errors.New(fmt.Sprintf("There are no possible entries for location (%v, %v)", row, col))
	}

	return entries, nil
}
func (p *Puzzle) quickGetEntries(row, col, box int) ([]types.Entry, error) {
	entries := utils.GetPossibleEntriesQuickly(p.rows[row], p.cols[col], p.boxs[box])

	if len(entries) == 0 {
		//p.PrintPretty()
		return nil, errors.New(`no entries possible`) //errors.New(fmt.Sprintf("There are no possible entries for location (%v, %v)", row, col))
	}

	return entries, nil
}

func (p *Puzzle) getEmptyTile() (row, col, box int, err error) {
	return p.getBestLocationSmartly()
}

func (p *Puzzle) getBestLocationSmartly() (row, col, box int, err error) {
	bkRow := types.NewBK()
	bkCol := types.NewBK()
	bkBox := types.NewBK()

	for i := 0; i < constants.SideLen; i++ {
		bkRow.Update(i, p.numFreeInRow[i])
		bkCol.Update(i, p.numFreeInCol[i])
	}

	if bkRow.IsError() || bkCol.IsError() {
		return -1, -1, -1, errors.New(`All the rows or cols are full!`)
	}

	if bkRow.IsBetterThan(bkCol) {
		p.getBestCol(bkRow.Location(), bkCol, bkBox)
	} else {
		p.getBestRow(bkCol.Location(), bkRow, bkBox)
	}

	return bkRow.Location(), bkCol.Location(), bkBox.Location(), nil
}
func (p *Puzzle) getBestCol(bestRowIndex int, bkCol, bkBox *types.BestKnown) {
	bkCol.Reset()

	for c := range p.numFreeInCol {
		if p.tiles[bestRowIndex][c] != constants.EmptyTile {
			continue
		}

		if !bkCol.WillUpdate(c, p.numFreeInCol[c]) {
			continue
		}

		b, _ := utils.GetBox(bestRowIndex, c)

		if p.numFreeInBox[b] == 0 {
			continue
		}

		bkCol.UpdateUnsafe(c, p.numFreeInCol[c])
		bkBox.UpdateUnsafe(b, p.numFreeInBox[b])
	}
}
func (p *Puzzle) getBestRow(bestColIndex int, bkRow, bkBox *types.BestKnown) {
	bkRow.Reset()

	for r := range p.numFreeInRow {
		if p.tiles[r][bestColIndex] != constants.EmptyTile {
			continue
		}

		if !bkRow.WillUpdate(r, p.numFreeInRow[r]) {
			continue
		}

		b, _ := utils.GetBox(r, bestColIndex)

		if p.numFreeInBox[b] == 0 {
			continue
		}

		bkRow.UpdateUnsafe(r, p.numFreeInRow[r])
		bkBox.UpdateUnsafe(b, p.numFreeInBox[b])
	}
}

func (p *Puzzle) place(row, col, box int, entry types.Entry) (bool, error) {
	numPlacements++

	ePresence := utils.PresenceOf(entry)

	if err := p.entryIsPresent(row, col, box, ePresence); err != nil {
		return false, err
	}

	p.tiles[row][col] = types.Tile(entry)

	p.rows[row] |= ePresence
	p.cols[col] |= ePresence
	p.boxs[box] |= ePresence

	p.numFreeInRow[row] -= 1
	p.numFreeInCol[col] -= 1
	p.numFreeInBox[box] -= 1

	p.numPlaced += 1

	return true, nil
}

func (p *Puzzle) entryIsPresent(row, col, box int, ePresence types.Presence) error {
	return p.entryIsPresentQuickly(row, col, box, ePresence)
	//return p.entryIsPresentSlowly(row, col, box, ePresence)
}
func (p *Puzzle) entryIsPresentSlowly(row, col, box int, ePresence types.Presence) error {
	if p.tiles[row][col] != constants.EmptyTile {
		return errors.New(fmt.Sprintf("Tile already exists at (%v, %v)", row, col))
	}

	if ePresence == constants.PresenceError {
		return errors.New(fmt.Sprintf(`Bad Entry Presence: %09b`, ePresence))
	}

	if utils.IsPresent(p.rows[row], ePresence) {
		return errors.New(fmt.Sprintf("Tile already exists in row %v", row))
	}

	if utils.IsPresent(p.cols[col], ePresence) {
		return errors.New(fmt.Sprintf("Tile already exists in col %v", col))
	}

	if utils.IsPresent(p.boxs[box], ePresence) {
		return errors.New(fmt.Sprintf("Tile already exists in box %v", box))
	}

	return nil
}
func (p *Puzzle) entryIsPresentQuickly(row, col, box int, ePresence types.Presence) error {
	if utils.IsPresent(p.rows[row]|p.cols[col]|p.boxs[box], ePresence) {
		return errors.New(`its present`)
	}

	return nil
}

/// Solves in 9.4802ms on average, with 4,585 average tries
func (p *Puzzle) solve() (solution *Puzzle, err error) {
	row, col, box, entries, err := p.getLocationAndEntries()
	if err != nil {
		// Unsolvable!
		return nil, err
	}

	for _, entry := range entries {
		pClone := p.clone()

		wasPlaced, _ := pClone.place(row, col, box, entry)

		if wasPlaced {
			if pClone.numPlaced == constants.NumTiles {
				return pClone, nil
			}

			cloneSolution, _ := pClone.solve()

			if cloneSolution != nil {
				return cloneSolution, nil
			}
		}
	}

	return nil, errors.New(`Should never be here`)
}
