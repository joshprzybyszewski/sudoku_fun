package utils

import (
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

func (t valueHitMap) clone() valueHitMap {
	return t
}

type Game struct {
	size           int
	gameboard      Board
	rowsSet        [9]valueHitMap
	colsSet        [9]valueHitMap
	subSquareSets  [3][3]valueHitMap
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
	}

	for r, row := range board {
		for c, value := range row {
			if value <= 0 {
				continue
			}

			g.recordValue(r, c, value)
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

	// if this is an invalid value for this location
	// or if this location is already set
	// early out
	if  g.GetValue(row, col) > 0 {
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
	g2.recordValue(row, col, value)
	return g2, true
}

func (g *Game) recordValue(row, col, value int) {
	g.size++
	g.gameboard[row][col] = value
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
			if g.gameboard[i][j] == 0 && g.valueSetAt(i, j) == fullMapping {
				return false
			}
		}
	}
	return true
}

func (g *Game) valueSetAt(row, col int) valueHitMap {
	return g.rowsSet[row] | g.colsSet[col]| g.subSquareSets[row/3][col/3]
}

func (g *Game) NextGuessMove() (row, col int, values []int) {
	var nextGuess *TakenNumbersItem

	for c := 0;c<9;c++{
		for r :=0;r<9;r++{
			if  g.gameboard[r][c] != 0 {
				continue
			}
			locationHitMap := g.valueSetAt(r, c)
			if locationHitMap == fullMapping {
				return 0, 0, nil
			}

			totalSet := locationHitMap.numberSet()
			if nextGuess == nil || totalSet > nextGuess.list.numberSet() {
				nextGuess = &TakenNumbersItem{
					list: locationHitMap,
					row:  r,
					col:  c,
				}
				if totalSet == 8 {
					return r, c, locationHitMap.inverse().getValues()
				}
			}
		}
	}

	if nextGuess == nil {
		return -1, -1, nil
	}
	return nextGuess.row, nextGuess.col, nextGuess.list.inverse().getValues()
}
