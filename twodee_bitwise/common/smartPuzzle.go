package common

import (
	"errors"
	"strconv"

	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	speedUtils "github.com/joshprzybyszewski/sudoku_fun/utils/speed"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

var numAttempts int

type SolverFn func(*SmartPuzzle) (solution *SmartPuzzle, err error)

type SmartPuzzle struct {
	NumPlaced int
	solver    SolverFn

	Tiles [constants.SideLen][constants.SideLen]types.Tile /* [row][col] */

	Rows [constants.SideLen]types.Presence
	Cols [constants.SideLen]types.Presence
	Boxs [constants.SideLen]types.Presence

	NumFreeInRow [constants.SideLen]uint8
	NumFreeInCol [constants.SideLen]uint8
	NumFreeInBox [constants.SideLen]uint8
}

func newPuzzle(slvr SolverFn) SmartPuzzle {
	pzl := SmartPuzzle{}

	for i := 0; i < constants.SideLen; i++ {
		pzl.NumFreeInRow[i] = constants.SideLen
		pzl.NumFreeInCol[i] = constants.SideLen
		pzl.NumFreeInBox[i] = constants.SideLen
	}

	pzl.solver = slvr

	return pzl
}

/// Reads in sudoku from a string representation of it
func GetSmartPuzzle(str string, slvr SolverFn) (p SmartPuzzle, err error) {
	pzl := newPuzzle(slvr)
	for i := 0; i < len(str); i++ {
		entryChar := string(str[i])
		entry, err := strconv.Atoi(entryChar)

		if entryChar == constants.EmptyTileChar || err != nil {
			continue
		}

		row := i / constants.SideLen
		col := i % constants.SideLen
		box, err := speedUtils.GetBox(row, col)
		if wasPlaced, err := pzl.Place(row, col, box, types.Entry(entry)); err != nil || !wasPlaced {
			return SmartPuzzle{}, errors.New(`BAD sudoku!!!`)
		}
	}

	return pzl, nil
}

func (p SmartPuzzle) GetNumPlacements() int {
	return numAttempts
}
func (p SmartPuzzle) GetSimple() string {
	line := ``

	for r := range p.Tiles {
		for c := range p.Tiles[r] {
			if p.Tiles[r][c] == constants.EmptyTile {
				line += constants.EmptyTileChar
			} else {
				line += strconv.Itoa(int(p.Tiles[r][c]))
			}
		}
	}

	return line
}
func (p SmartPuzzle) PrintSimple() {
	println(p.GetSimple())
}
func (p SmartPuzzle) PrintPretty() {
	for r := range p.Tiles {
		line := ``

		for c := range p.Tiles[r] {
			line += ` `

			if p.Tiles[r][c] == constants.EmptyTile {
				line += constants.EmptyTileChar
			} else {
				line += strconv.Itoa(int(p.Tiles[r][c]))
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

func (p SmartPuzzle) Solve() (solution types.Sudoku, err error) {
	numAttempts = 0
	return p.solver(&p)
}

func (p *SmartPuzzle) entryIsPresent(row, col, box int, ePresence types.Presence) error {
	if speedUtils.IsPresent(p.Rows[row]|p.Cols[col]|p.Boxs[box], ePresence) {
		return errors.New(`its present`)
	}

	return nil
}

func (p *SmartPuzzle) Place(row, col, box int, entry types.Entry) (bool, error) {
	numAttempts++

	ePresence := speedUtils.PresenceOf(entry)

	if err := p.entryIsPresent(row, col, box, ePresence); err != nil {
		return false, err
	}

	p.Tiles[row][col] = types.Tile(entry)

	p.Rows[row] |= ePresence
	p.Cols[col] |= ePresence
	p.Boxs[box] |= ePresence

	p.NumFreeInRow[row] -= 1
	p.NumFreeInCol[col] -= 1
	p.NumFreeInBox[box] -= 1

	p.NumPlaced += 1

	return true, nil
}

func (p *SmartPuzzle) Clone() *SmartPuzzle {
	clone := SmartPuzzle{}

	clone = *p

	return &clone
}
