package twodee

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/joshprzybyszewski/sudoku_fun/twodee/common"
	"github.com/joshprzybyszewski/sudoku_fun/twodee/smart"
	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	"github.com/joshprzybyszewski/sudoku_fun/utils/speed"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

func errorPuzzleReader(_ string) (types.Sudoku, error) {
	return nil, errors.New(`lol broken`)
}

func errorOnBadCharPuzzleReader(str string) (types.Sudoku, error) {
	pzl := common.SmartPuzzle{}
	for i := 0; i < len(str); i++ {
		entryChar := string(str[i])
		entry, err := strconv.Atoi(entryChar)

		if err != nil {
			return types.Sudoku(common.SmartPuzzle{}), err
		} else if entryChar == constants.EmptyTileChar {
			continue
		}

		row := i / constants.SideLen
		col := i % constants.SideLen
		box, err := speed.GetBox(row, col)
		if wasPlaced, err := pzl.Place(row, col, box, types.Entry(entry)); err != nil || !wasPlaced {
			return types.Sudoku(common.SmartPuzzle{}), errors.New(`BAD sudoku!!!`)
		}
	}

	return pzl, nil
}

func badPuzzleReader(str string) (types.Sudoku, error) {
	pzl := common.SmartPuzzle{}
	pzl.Solver = badSolver
	for i := 0; i < len(str); i++ {
		row := i / constants.SideLen
		col := i % constants.SideLen
		pzl.Tiles[row][col] = 5
	}

	return pzl, nil
}

func smartRead(str string) (s types.Sudoku, err error) {
	pzl, err := common.GetSmartPuzzle(str, smart.Solve)
	if err != nil {
		return nil, err
	}

	return types.Sudoku(pzl), nil
}

func badSolver(_ *common.SmartPuzzle) (solution *common.SmartPuzzle, err error) {
	pzl := common.SmartPuzzle{}
	pzl.Tiles[4][4] = 4
	return &pzl, nil
}

func puzzleReaderWithABadSolver(entries string) (s types.Sudoku, err error) {
	pzl, err := common.GetSmartPuzzle(entries, badSolver)
	if err != nil {
		return nil, err
	}

	return types.Sudoku(pzl), nil
}

func Test_PuzzleSolver(t *testing.T) {
	justFineSudoku := `..812...9.6.........2..95......8.93....2..68...........564..3....9...41..8..1..56`
	justFineAnswer := `348125769965347821712869543621784935573291684894536172156472398239658417487913256`

	anotherFineSudoku := `5..96....6...4..5..498.................5.2..326...794.4..3...2...67..8..72....3..`
	anotherFineAnswer := `532961478687243159149875632975634281814592763263187945498316527356729814721458396`

	testCases := []struct {
		msg          string
		sudokuReader SudokuReader
		puzzleString string
		expNumTries  int
		expSolution  string
		expError     bool
	}{{
		msg:          `Having a Bad Sudoku Reader`,
		sudokuReader: errorPuzzleReader,
		puzzleString: `lol this doesn't matter because the puzzle reader is already borked`,
		expNumTries:  0,
		expSolution:  ``,
		expError:     true,
	}, {
		msg:          `Having a VERY bad Sudoku`,
		sudokuReader: errorOnBadCharPuzzleReader,
		puzzleString: `something that's not even a sudoku string`,
		expSolution:  ``,
		expError:     true,
	}, {
		msg:          `Having a Bad Sudoku`,
		sudokuReader: smartRead,
		puzzleString: `44..33.11..11..22..33..44..56.7.8.9..9.0.7.6.5.4...23.4..234.12.2.3.4.1...4..1.12`,
		expSolution:  ``,
		expError:     true,
	}, {
		msg:          `Having a Bad Sudoku Reader`,
		sudokuReader: badPuzzleReader,
		puzzleString: justFineSudoku,
		expSolution:  constants.EmptyPuzzle,
		expError:     true,
	}, {
		msg:          `Having a Bad Sudoku Solver`,
		sudokuReader: puzzleReaderWithABadSolver,
		puzzleString: justFineSudoku,
		expSolution:  constants.EmptyPuzzle,
		expError:     true,
	}, {
		msg:          `Having a Good Sudoku`,
		sudokuReader: smartRead,
		puzzleString: justFineSudoku,
		expNumTries:  3737,
		expSolution:  justFineAnswer,
		expError:     false,
	}, {
		msg:          `Having another Good Sudoku`,
		sudokuReader: smartRead,
		puzzleString: anotherFineSudoku,
		expNumTries:  309,
		expSolution:  anotherFineAnswer,
		expError:     false,
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("Failed %v test", tc.msg)
		actDur, actNumTries, actSolution, actError := PuzzleSolver(tc.sudokuReader, tc.puzzleString)
		if tc.expError {
			require.Error(t, actError, failMsg)
			assert.Nil(t, actDur, failMsg)
			assert.Equal(t, -1, actNumTries, failMsg)
		} else {
			require.NoError(t, actError, failMsg)
			lower := int64(0)
			upper := int64(tc.expNumTries * 5000)
			durFailMsg := fmt.Sprintf("%v: %v <= %v <= %v", failMsg, lower, actDur.Nanoseconds(), upper)
			assert.True(t, actDur.Nanoseconds() >= lower, durFailMsg)
			assert.True(t, actDur.Nanoseconds() <= upper, durFailMsg)
			assert.Equal(t, tc.expNumTries, actNumTries, failMsg)
		}
		assert.Equal(t, tc.expSolution, actSolution, failMsg)
	}
}
