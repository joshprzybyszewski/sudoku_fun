package robust

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	utils "github.com/joshprzybyszewski/sudoku_fun/utils/slow"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
	"github.com/joshprzybyszewski/sudoku_fun/utils/speed"
)

var numPlacements int
var maxPlaced = 0

type Puzzle struct {
	numPlaced int

	tiles [constants.SideLen][constants.SideLen][constants.SideLen]types.Tile /* [row][col] */

	rows [constants.SideLen]types.Presence
	cols [constants.SideLen]types.Presence
	pils [constants.SideLen]types.Presence


	numFreeInRow [constants.SideLen]uint8
	numFreeInCol [constants.SideLen]uint8
	numFreeInPil [constants.SideLen]uint8

	boxs [constants.SideLen]types.Presence
	numFreeInBox [constants.SideLen]uint8
}

func newPuzzle() Puzzle {
	pzl := Puzzle{}

	for i := 0; i < constants.SideLen; i++ {
		pzl.numFreeInRow[i] = constants.SideLen
		pzl.numFreeInCol[i] = constants.SideLen
		pzl.numFreeInPil[i] = constants.SideLen

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

		if entryChar == constants.EmptyTileChar {
			continue
		}

		entry, err := strconv.Atoi(entryChar)
		if err != nil {
			return Puzzle{}, err
		}

		row := i / constants.SideLen
		col := i % constants.SideLen
		pil := 0
		//box, err := utils.GetBox(row, col)
		wasPlaced, err := pzl.place(row, col, pil, types.Entry(entry))
		if err != nil {
			return Puzzle{}, err
		}
		if !wasPlaced {
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
			for pil := range p.tiles[r][c] {
			if p.tiles[r][c][pil] == constants.EmptyTile {
				line += constants.EmptyTileChar
			} else {
				line += strconv.Itoa(int(p.tiles[r][c][pil]))
			}
			}
		}
	}

	return line
}
func (p Puzzle) PrintSimple() {
	println(p.GetSimple())
}
func (p Puzzle) PrintPretty() {
	for i := 0; i < constants.SideLen; i++ {
		for j := 0; j < constants.SideLen; j++ {
			line := ``
			for k := 0; k < constants.SideLen; k++ {
				line += ` `

				if p.tiles[j][k][i] == constants.EmptyTile {
					line += constants.EmptyTileChar
				} else {
					line += strconv.Itoa(int(p.tiles[j][k][i]))
				}

				line += ` `

				if k == 2 || k == 5 {
					line += ` | `
				}
			}

			println(line)

			if j == 2 || j == 5 {
				println(`-+--+--+--|--+--+--+--|--+--+--+-`)
			}
		}
		println(`next level`)
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
	row, col, box, err = p.getEmptyTile()
	if err != nil {
		return -1, -1, -1, nil, errors.New(fmt.Sprintf("There are no empty tiles!"))
	}

	entries, err = p.getEntries(row, col, box)
	if err != nil {
		return -1, -1, -1, nil, errors.New(fmt.Sprintf("Errored on looking for entries for location (%v, %v): %v", row, col, box))
	}

	if len(entries) == 0 {
		return -1, -1, -1, nil, errors.New(fmt.Sprintf("There are no possible entries for location (%v, %v): %v", row, col, box))
	}

	return row, col, box, entries, nil
}
func (p *Puzzle) getEntries(row, col, box int) ([]types.Entry, error) {
	entries := utils.GetPossibleEntries(constants.AllEntries, p.rows[row])
	entries = utils.GetPossibleEntries(entries, p.cols[col])
	entries = utils.GetPossibleEntries(entries, p.boxs[box])

	if len(entries) == 0 {
		return nil, errors.New(fmt.Sprintf("There are no possible entries for location (%v, %v)", row, col))
	}

	return entries, nil
}
func (p *Puzzle) getEmptyTile() (row, col, box int, err error) {
	bkRow := types.NewBK()
	bkCol := types.NewBK()
	bkBox := types.NewBK()

	for i := 0; i < constants.SideLen; i++ {
		bkRow.Update(i, p.numFreeInRow[i])
		bkCol.Update(i, p.numFreeInCol[i])
	}

	if bkRow.IsError() {
		return -1, -1, -1, errors.New(`All the rows are full!`)
	}
	if bkCol.IsError() {
		return -1, -1, -1, errors.New(`All the cols are full!`)
	}

	if bkRow.IsBetterThan(bkCol) {
		p.getBestCol(bkRow.Location(), bkCol, bkBox)
	} else {
		p.getBestRow(bkCol.Location(), bkRow, bkBox)
	}

	if bkRow.IsError() {
		return -1, -1, -1, errors.New(`Couldn't find best known row'!`)
	}
	if bkCol.IsError() {
		return -1, -1, -1, errors.New(`Couldn't find best known col'!`)
	}
	if bkBox.IsError() {
		return -1, -1, -1, errors.New(`Couldn't find best known box'!`)
	}

	return bkRow.Location(), bkCol.Location(), bkBox.Location(), nil
}
func (p *Puzzle) getBestCol(bestRowIndex int, bkCol, bkBox *types.BestKnown) {
	bkCol.Reset()

	for c := range p.numFreeInCol {
		if p.tiles[bestRowIndex][c][0] != constants.EmptyTile {
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
		if p.tiles[r][bestColIndex][0] != constants.EmptyTile {
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

func (p *Puzzle) place(row, col, pil int, entry types.Entry) (bool, error) {
	numPlacements++

	if numPlacements >= 986409 {
	//if numPlacements % 1000 == 0 {
		println(`placed `, numPlacements, `times`)
		println(`max filled`, maxPlaced)
	}

	ePresence := utils.PresenceOf(entry)

	if err := p.entryIsPresent(row, col, pil, ePresence); err != nil {
		return false, err
	}

	//box := 0

	p.tiles[row][col][pil] = types.Tile(entry)

	p.rows[row] |= ePresence
	p.cols[col] |= ePresence
	p.pils[pil] |= ePresence

	//p.boxs[box] |= ePresence

	p.numFreeInRow[row] -= 1
	p.numFreeInCol[col] -= 1
	p.numFreeInPil[pil] -= 1

	//p.numFreeInBox[box] -= 1

	p.numPlaced += 1

	if maxPlaced < p.numPlaced {
		maxPlaced = p.numPlaced
		p.PrintPretty()
	}

	return true, nil
}

func (p *Puzzle) entryIsPresent(row, col, pil int, ePresence types.Presence) error {
	if p.tiles[row][col][pil] != constants.EmptyTile {
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

	if utils.IsPresent(p.pils[pil], ePresence) {
		return errors.New(fmt.Sprintf("Tile already exists in pillar %v", pil))
	}

	return nil
}

func (p *Puzzle) getNextLocationAndEntries() (row, col, pil int, entries []types.Entry, err error) {
	row, col, pil, err = p.getNextLocation()

	presence := p.rows[row] | p.cols[col] | p.pils[pil]
	numFree, entries := speed.GetCachedNumFreeAndPosEntries(presence)
	if numFree == 0 {
		println(`presence`, presence)
		println(`numFree`, numFree)
		println(`entries`, entries)
		println(`row`, row)
		println(`col`, col)
		println(`pil`, pil)
		println(`p.rows[row]`, p.rows[row])
		println(`p.cols[col]`, p.cols[col])
		println(`p.pils[pil]`, p.pils[pil])

		return -1, -1, -1, nil, errors.New(`idk nemore`)

	}

	return row, col, pil, entries, nil
}

func (p *Puzzle) getNextLocation() (row, col, pil int, err error) {
	for i := 0; i < constants.SideLen; i++ {
 		for j := 0; j < constants.SideLen; j++ {
			for k := 0; k < constants.SideLen; k++ {
				if p.tiles[i][j][k] == constants.EmptyTile {
					return i, j, k, nil
				}
			}
		}
	}


	println(`num placed`, p.numPlaced)
	return -1, -1, -1, errors.New(`idk nemore`)
}

func (p *Puzzle) solve() (solution *Puzzle, err error) {
	//row, col, pil, box, entries, err := p.getLocationAndEntries()
	row, col, pil, entries, err := p.getNextLocationAndEntries()

	if err != nil {
		// Unsolvable!
		utils.PrintError(`idk what`, err)
		return nil, err
	}

	for _, entry := range entries {
		pClone := p.clone()

		wasPlaced, err := pClone.place(row, col, pil, entry)
		if err != nil {
			utils.PrintError(`failed to place`, err)
			return nil, err
		}

		if wasPlaced {
			if pClone.numPlaced == constants.NumTiles3d {
				return pClone, nil
			}

			cloneSolution, err := pClone.solve()
			if err != nil {
				continue
			}

			if cloneSolution != nil {
				return cloneSolution, nil
			}
		}
	}

	return nil, errors.New(`Should never be here`)
}
