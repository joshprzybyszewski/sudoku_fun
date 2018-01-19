package utils

import (
	"sort"
	"strconv"
	"strings"
)

type Board [9][9]int

func (b Board) Print() string {
	fullStr := []string{}
	for rowNum, row := range b {
		str := []string{}
		for colNum, value := range row {
			if colNum%3 == 0 {
				str = append(str, "*")
			}
			if value == 0 {
				str = append(str, `   `)
			} else {
				str = append(str, " "+strconv.Itoa(value)+" ")
			}
		}
		str = append(str, " ")
		if rowNum%3 == 0 {
			fullStr = append(fullStr, "---------------------------------------------")
		}
		fullStr = append(fullStr, "||"+strings.Join(str, "|")+"||")
	}
	fullStr = append(fullStr, "---------------------------------------------")
	return strings.Join(fullStr, "\n")
}

type TakenNumbersItem struct {
	row  int
	col  int
	list valueHitMap
}
type TakenNumbersItems []TakenNumbersItem

func (t TakenNumbersItems) Len() int { return len(t) }
func (t TakenNumbersItems) Less(i, j int) bool {
	iSet := t[i].list.numberSet()
	jSet := t[j].list.numberSet()
	if iSet > jSet {
		return true
	}
	if iSet == jSet {
		return t[i].row < t[j].row
	}
	return false
}
func (t TakenNumbersItems) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func (b Board) clone() Board {
	return b
}

func (t valueHitMap) clone() valueHitMap {
	return t
}

type Game struct {
	size           int
	gameboard      Board
	rowsSet        [9]valueHitMap
	colsSet        [9]valueHitMap
	subSquareSets  [3][3]valueHitMap
	individualSets [9][9]valueHitMap
}

func NewGame() *Game {
	return NewGameWithState(Board{})
}

func NewGameWithState(board Board) *Game {
	g := &Game{
		gameboard:      Board{},
		rowsSet:        [9]valueHitMap{},
		colsSet:        [9]valueHitMap{},
		subSquareSets:  [3][3]valueHitMap{},
		individualSets: [9][9]valueHitMap{},
	}

	for r, row := range board {
		for c, value := range row {
			if value <= 0 {
				continue
			}

			g.recordValue(r, c, value)
			g.gameboard[r][c] = value
			g.size++
		}
	}
	return g
}

func (g *Game) Copy() *Game {
	return g.clone()
}

func (g *Game) clone() *Game {
	return &Game{
		gameboard:      g.gameboard,
		size:           g.size,
		rowsSet:        g.rowsSet,
		colsSet:        g.colsSet,
		subSquareSets:  g.subSquareSets,
		individualSets: g.individualSets,
	}
}

// IsValid checks if the number is a valid number
func (g *Game) GetValue(row, col int) int {
	if row >= 9 || col >= 9 {
		return -1
	}

	val := g.gameboard[row][col]
	if val == 0 {
		val = -1
	}
	return val
}

// IsValid checks if the number is a valid number
func (g *Game) IsValid(row, col, value int) bool {
	if row > 9 || col > 9 || value <= 0 || value > 9 || g.Complete() {
		return false
	}

	invalidValue := g.individualSets[row][col].valueIsSet(value)
	// if this is an invalid value for this location
	// or if this location is already set
	// early out
	if invalidValue || g.GetValue(row, col) > 0 {
		return false
	}

	// integer math-ing ftw
	xSquare := (col / 3)
	ySquare := (row / 3)

	invalidRowValue := g.rowsSet[row].valueIsSet(value)
	invalidColValue := g.colsSet[col].valueIsSet(value)
	invalidSquareValue := g.subSquareSets[ySquare][xSquare].valueIsSet(value)
	return !invalidRowValue && !invalidColValue && !invalidSquareValue
}

func (g *Game) Complete() bool {
	return g.size == 81
}

func (g *Game) Set(row, col, value int) (*Game, bool) {
	if !g.IsValid(row, col, value) {
		return nil, false
	}

	g2 := g.clone()
	g2.size++
	g2.gameboard[row][col] = value
	g2.recordValue(row, col, value)
	return g2, true
}

func (g *Game) recordValue(row, col, value int) {
	colSquare := (col / 3) * 3
	rowSquare := (row / 3) * 3

	for r := 0; r < 9; r++ {
		g.individualSets[r][col] = g.individualSets[r][col].setValue(value)
		if r >= rowSquare && r < rowSquare+3 {
			for c := 0; c < 3; c++ {
				g.individualSets[r][c+colSquare] = g.individualSets[r][c+colSquare].setValue(value)
			}
		}
	}
	for c := 0; c < 9; c++ {
		g.individualSets[row][c] = g.individualSets[row][c].setValue(value)
		if c >= colSquare && c < colSquare+3 {
			for r := 0; r < 3; r++ {
				g.individualSets[r+rowSquare][c] = g.individualSets[r+rowSquare][c].setValue(value)
			}
		}
	}

	g.rowsSet[row] = g.rowsSet[row].setValue(value)
	g.colsSet[col] = g.colsSet[col].setValue(value)
	g.subSquareSets[row/3][col/3] = g.subSquareSets[row/3][col/3].setValue(value)
}

func (g *Game) Print() string {
	return g.gameboard.Print()
}

func (g *Game) ValidGame() bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if g.gameboard[i][j] == 0 && g.individualSets[i][j] == fullMapping {
				return false
			}
		}
	}
	return true
}

func (g *Game) NextGuessMove() (row, col int, values []int) {
	listOfPossibleGuesses := TakenNumbersItems{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if g.individualSets[i][j] != fullMapping && g.gameboard[i][j] == 0 {
				listOfPossibleGuesses = append(listOfPossibleGuesses, TakenNumbersItem{
					list: g.individualSets[i][j],
					row:  i,
					col:  j,
				})
			}
			if g.gameboard[i][j] == 0 && g.individualSets[i][j] == fullMapping {
				return 0, 0, nil
			}
		}
	}

	if len(listOfPossibleGuesses) > 0 {
		sort.Sort(listOfPossibleGuesses)

		set := listOfPossibleGuesses[0]
		return set.row, set.col, set.list.inverse().getValues()
	}
	return -1, -1, nil
}
