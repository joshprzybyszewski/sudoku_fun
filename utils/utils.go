package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

const (
	Root     = 3
	SideLen  = Root * Root
	NumTiles = SideLen * SideLen

	PresenceError = EmptyPresence
	EmptyPresence = Presence(0)
	FullPresence  = Presence(511) // calc'd cuz first 9 bits = 2^9 - 1

	EmptyTile     = Tile(0)
	EmptyTileChar = `.`

	verbose = false
	speed = true

	onePresence   = Presence(1 << uint(1-1))
	twoPresence   = Presence(1 << uint(2-1))
	threePresence = Presence(1 << uint(3-1))
	fourPresence  = Presence(1 << uint(4-1))
	fivePresence  = Presence(1 << uint(5-1))
	sixPresence   = Presence(1 << uint(6-1))
	sevenPresence = Presence(1 << uint(7-1))
	eightPresence = Presence(1 << uint(8-1))
	ninePresence  = Presence(1 << uint(9-1))

	bestKnownLocationStart = -1

	EmptyPuzzle = `.................................................................................`
)

var (
	AllEntries       = []Entry{1, 2, 3, 4, 5, 6, 7, 8, 9}

	allEntriesPresence = []Presence{
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
)

type Entry int
type Tile uint8
type Presence uint16

type Sudoku interface {
	Solve() (s Sudoku, err error)
	GetNumPlacements() int
	GetSimple() string
	PrintSimple()
	PrintPretty()
}

type BestKnown struct {
	loc      int
	val      uint8
	lower    uint8
}

func (bk *BestKnown) WillUpdate(loc int, val uint8) bool {
	return val < bk.val && val > bk.lower
}

func (bk *BestKnown) Update(loc int, val uint8) {
	if bk.WillUpdate(loc, val) {
		bk.UpdateUnsafe(loc, val)
	}
}

func (bk *BestKnown) UpdateUnsafe(loc int, val uint8) {
	bk.loc = loc
	bk.val = val
}

func (bk *BestKnown) IsError() (bool) {
	return bk.loc == bestKnownLocationStart
}

func (bk *BestKnown) IsStrictlyBetterThan(other *BestKnown) (bool) {
	return bk.val < other.val
}

func (bk *BestKnown) IsBetterThan(other *BestKnown) (bool) {
	return bk.val <= other.val
}

func (bk *BestKnown) Location() int {
	return bk.loc
}

func (bk *BestKnown) Reset() {
	bk.loc = bestKnownLocationStart
	bk.val = uint8(SideLen + 1)
	bk.lower = 0
}

func NewBK() (*BestKnown) {
	bk := &BestKnown{}
	bk.Reset()
	return bk
}



func GetPossibleEntries(entries []Entry, presence Presence) []Entry {
	posEntries := make([]Entry, 0, SideLen)
	for _, entry := range entries {
		ePresence := PresenceOf(entry)
		if !IsPresent(presence, ePresence) {
			posEntries = append(posEntries, entry)
		}
	}
	return posEntries
}

func GetPossibleEntriesQuickly(rowP, colP, boxP Presence) []Entry {
	posEntries := make([]Entry, 0, SideLen)
	presence := rowP | colP | boxP

	if !IsPresent(presence, onePresence) {
		posEntries = append(posEntries, Entry(1))
	}
	if !IsPresent(presence, twoPresence) {
		posEntries = append(posEntries, Entry(2))
	}
	if !IsPresent(presence, threePresence) {
		posEntries = append(posEntries, Entry(3))
	}
	if !IsPresent(presence, fourPresence) {
		posEntries = append(posEntries, Entry(4))
	}
	if !IsPresent(presence, fivePresence) {
		posEntries = append(posEntries, Entry(5))
	}
	if !IsPresent(presence, sixPresence) {
		posEntries = append(posEntries, Entry(6))
	}
	if !IsPresent(presence, sevenPresence) {
		posEntries = append(posEntries, Entry(7))
	}
	if !IsPresent(presence, eightPresence) {
		posEntries = append(posEntries, Entry(8))
	}
	if !IsPresent(presence, ninePresence) {
		posEntries = append(posEntries, Entry(9))
	}

	return posEntries
}

func PresenceOf(entry Entry) Presence {
	return presenceOfSpeed(entry)
	return presenceOfSlow(entry)
}
func presenceOfSpeed(entry Entry) Presence {
		return Presence(1 << uint(entry-1))
}
func presenceOfSlow(entry Entry) Presence {
	if entry <= 0 || entry > SideLen {
		return PresenceError
	}

	return Presence(1 << uint(entry-1))
}

func IsPresent(existing, entryPresence Presence) bool {
	return existing&entryPresence > 0
}

func GetBox(row, col int) (int, error) {
	return getBoxSpeed(row, col)
	return getBoxSlow(row, col)
}
func getBoxSpeed(row, col int) (int, error) {
	return ((row / Root) * Root) + (col / Root), nil
}
func getBoxSlow(row, col int) (int, error) {
	if row >= SideLen || row < 0 ||
		col >= SideLen || col < 0 {
		return -1, errors.New(`Row or col was out of bounds!`)
	}

	// 0 | 1 | 2
	// 3 | 4 | 5
	// 6 | 7 | 8

	box := ((row / Root) * Root) + (col / Root)

	return box, nil
}

func GetNumFreeSpots(presence Presence) int {
	return getNumFreeSpotsSpeed(presence)
	return getNumFreeSpotsSlow(presence)
}
func getNumFreeSpotsSpeed(presence Presence) int {
	numFree := 0

	for _, ePresence := range allEntriesPresence {
		if IsPresent(presence, ePresence) {
			numFree++
		}
	}

	return numFree
}
func getNumFreeSpotsSlow(presence Presence) int {
	numFree := 0

	for i := 1; i <= SideLen; i++ {
		ePresence := PresenceOf(Entry(i))
		if IsPresent(presence, ePresence) {
			numFree += 1
		}
	}

	return numFree
}

func PrintError(str string, err error) {
	if verbose {
		println(fmt.Sprintf("Error @ %v: %v", str, err.Error()))
	}
}

func PrintPlacement(e Entry, r, c int) {
	if verbose {
		println(fmt.Sprintf("Attempting to place %v at (%v, %v)", e, r, c))
	}
}

func BruteForceCheck(sudokuStr string) bool {
	var entries [SideLen][SideLen]int
	for i := 0; i < len(sudokuStr); i++ {
		entryChar := string(sudokuStr[i])
		entry, err := strconv.Atoi(entryChar)

		if entryChar == EmptyTileChar || err != nil {
			return false
		}

		row := i / SideLen
		col := i % SideLen

		entries[row][col] = entry
	}

	for r := range entries {
		for c := range entries[r] {
			value := entries[r][c]
			baseRow := (r / Root) * Root
			baseCol := (c / Root) * Root

			for i := 0; i < SideLen; i++ {
				if i != r && entries[i][c] == value {
					return false
				}

				if i != c && entries[r][i] == value {
					return false
				}

				boxEntryR := baseRow + (i / Root)
				boxEntryC := baseCol + (i % Root)

				if boxEntryR != r && boxEntryC != c && entries[boxEntryR][boxEntryC] == value {
					return false
				}
			}
		}
	}

	return true
}
