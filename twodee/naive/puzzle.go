package naive

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	utils "github.com/joshprzybyszewski/sudoku_fun/utils/slow"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

var numPlacements int

type Puzzle struct {
	numPlaced int

	tiles [constants.SideLen][constants.SideLen]types.Tile /* [row][col] */

	rows [constants.SideLen]types.Presence
	cols [constants.SideLen]types.Presence
	boxs [constants.SideLen]types.Presence
}

/// Reads in sudoku from a string representation of it
func ReadSudoku(entries string) (p Puzzle, err error) {
	numPlacements = 0

	pzl := Puzzle{}
	for i := 0; i < len(entries); i++ {
		entryChar := string(entries[i])
		entry, err := strconv.Atoi(entryChar)

		if entryChar == constants.EmptyTileChar || err != nil {
			continue
		}

		row := i / constants.SideLen
		col := i % constants.SideLen
		if wasPlaced, err := pzl.place(row, col, types.Entry(entry)); err != nil || !wasPlaced {
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

	clone.numPlaced = p.numPlaced

	for r := range p.tiles {
		for c := range p.tiles[r] {
			clone.tiles[r][c] = p.tiles[r][c]
		}
	}

	for i := 0; i < constants.SideLen; i++ {
		clone.rows[i] = p.rows[i]
		clone.cols[i] = p.cols[i]
		clone.boxs[i] = p.boxs[i]
	}

	return &clone
}

func (p *Puzzle) getLocationAndEntries() (row, col int, entries []types.Entry, err error) {
	row, col, err = p.getEmptyTile()
	if err != nil {
		return -1, -1, nil, err
	}

	box, err := utils.GetBox(row, col)
	if err != nil {
		return -1, -1, nil, err
	}

	entries = constants.AllEntries
	entries = utils.GetPossibleEntries(entries, p.rows[row])
	entries = utils.GetPossibleEntries(entries, p.cols[col])
	entries = utils.GetPossibleEntries(entries, p.boxs[box])

	if len(entries) == 0 {
		return -1, -1, nil, errors.New(fmt.Sprintf("There are no possible entries for location (%v, %v)", row, col))
	}

	return row, col, entries, nil
}

/// This just scans the puzzle row by row, col by col, until it finds
/// the first empty Tile
/// On the test cases produced:
/// Average Execution Time: ~ 77036622 ns
/// Average Number of Placements: 91718
func (p *Puzzle) getEmptyTile() (row, col int, err error) {
	for r := range p.tiles {
		if p.rows[r] == constants.FullPresence {
			continue
		}

		for c := range p.tiles[r] {
			if p.tiles[r][c] == constants.EmptyTile {
				return r, c, nil
			}
		}
	}
	return -1, -1, errors.New(`There are no empty tiles!`)
}

func (p *Puzzle) place(row, col int, entry types.Entry) (bool, error) {
	numPlacements++

	if p.tiles[row][col] != constants.EmptyTile {
		return false, errors.New(fmt.Sprintf("Tile already exists at (%v, %v)", row, col))
	}

	box, err := utils.GetBox(row, col)
	if err != nil {
		return false, err
	}

	ePresence := utils.PresenceOf(entry)
	if ePresence == constants.PresenceError {
		return false, errors.New(fmt.Sprintf("Invalid entry conversion to presence: %v", entry))
	}

	if utils.IsPresent(p.rows[row], ePresence) {
		return false, errors.New(fmt.Sprintf("Tile already exists in row %v", row))
	}

	if utils.IsPresent(p.cols[col], ePresence) {
		return false, errors.New(fmt.Sprintf("Tile already exists in col %v", col))
	}

	if utils.IsPresent(p.boxs[box], ePresence) {
		return false, errors.New(fmt.Sprintf("Tile already exists in box %v", box))
	}

	p.tiles[row][col] = types.Tile(entry)

	p.rows[row] |= ePresence
	p.cols[col] |= ePresence
	p.boxs[box] |= ePresence

	p.numPlaced += 1

	return true, nil
}

func (p *Puzzle) solve() (solution *Puzzle, err error) {
	if p.numPlaced == constants.NumTiles {
		return p, nil
	}

	row, col, entries, err := p.getLocationAndEntries()
	if err != nil {
		// Unsolvable!
		//utilPrint(fmt.Sprintf("Get Location and Entries Error: %v", err.Error()))
		return nil, err
	}

	for _, entry := range entries {
		pClone := p.clone()

		//utilPrint(fmt.Sprintf("Attempting to place %v at (%v, %v)", entry, row, col))
		wasPlaced, err := pClone.place(row, col, entry)
		if err != nil {
			// Decide how to handle a place error
			//utilPrint(fmt.Sprintf("Placement Error: %v", err.Error()))
		}

		if wasPlaced {
			cloneSolution, err := pClone.solve()
			if err != nil {
				// Decide how to handle a solution error
				//utilPrint(fmt.Sprintf("Solve Error: %v", err.Error()))
			}

			if cloneSolution != nil {
				return cloneSolution, nil
			}
		}
	}

	return nil, nil
}
