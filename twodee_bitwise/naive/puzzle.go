package naive

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	"../../utils"
)

var numPlacements int

type Puzzle struct {
	numPlaced int

	tiles [utils.SideLen][utils.SideLen]utils.Tile /* [row][col] */

	rows [utils.SideLen]utils.Presence
	cols [utils.SideLen]utils.Presence
	boxs [utils.SideLen]utils.Presence
}

/// Reads in sudoku from a string representation of it
func ReadSudoku(entries string) (p Puzzle, err error) {
	pzl := Puzzle{}
	for i := 0; i < len(entries); i++ {
		entryChar := string(entries[i])
		entry, err := strconv.Atoi(entryChar)

		if entryChar == utils.EmptyTileChar || err != nil {
			continue
		}

		row := i / utils.SideLen
		col := i % utils.SideLen
		if wasPlaced, err := pzl.place(row, col, utils.Entry(entry)); err != nil || !wasPlaced {
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
			if p.tiles[r][c] == utils.EmptyTile {
				line += utils.EmptyTileChar
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

			if p.tiles[r][c] == utils.EmptyTile {
				line += utils.EmptyTileChar
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
func (p Puzzle) Solve() (solution utils.Sudoku, err error) {
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

	for i := 0; i < utils.SideLen; i++ {
		clone.rows[i] = p.rows[i]
		clone.cols[i] = p.cols[i]
		clone.boxs[i] = p.boxs[i]
	}

	return &clone
}

func (p *Puzzle) getLocationAndEntries() (row, col int, entries []utils.Entry, err error) {
	row, col, err = p.getEmptyTile()
	if err != nil {
		return -1, -1, nil, err
	}

	box, err := utils.GetBox(row, col)
	if err != nil {
		return -1, -1, nil, err
	}

	entries = utils.AllEntries
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
		if p.rows[r] == utils.FullPresence {
			continue
		}

		for c := range p.tiles[r] {
			if p.tiles[r][c] == utils.EmptyTile {
				return r, c, nil
			}
		}
	}
	return -1, -1, errors.New(`There are no empty tiles!`)
}

func (p *Puzzle) place(row, col int, entry utils.Entry) (bool, error) {
	numPlacements++

	if p.tiles[row][col] != utils.EmptyTile {
		return false, errors.New(fmt.Sprintf("Tile already exists at (%v, %v)", row, col))
	}

	box, err := utils.GetBox(row, col)
	if err != nil {
		return false, err
	}

	ePresence := utils.PresenceOf(entry)
	if ePresence == utils.PresenceError {
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

	p.tiles[row][col] = utils.Tile(entry)

	p.rows[row] |= ePresence
	p.cols[col] |= ePresence
	p.boxs[box] |= ePresence

	p.numPlaced += 1

	return true, nil
}

func (p *Puzzle) solve() (solution *Puzzle, err error) {
	if p.numPlaced == utils.NumTiles {
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
