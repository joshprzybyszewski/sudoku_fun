package twodee_bitwise

import (
	"fmt"
	"time"

	"github.com/joshprzybyszewski/sudoku_fun/utils"
	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

type SudokuReader func(str string) (p types.Sudoku, err error)

func PuzzleSolver(readPuzzle SudokuReader, sudokuPuzzle string) (*time.Duration, int, string) {
	puzzle, err := readPuzzle(sudokuPuzzle)
	if err != nil {
		println(`BAD SUDOKU!!`)
		return nil, 0, ``
	}

	t0 := time.Now()
	solution, err := puzzle.Solve()
	duration := time.Since(t0)

	if err != nil {
		println(fmt.Sprintf("COULDN'T SOLVE: %v", err.Error()))
		return nil, -1, constants.EmptyPuzzle
	}

	if !utils.BruteForceCheck(solution.GetSimple()) {
		solution.PrintPretty()
		println(`WARNING IT DIDN'T ACTUALLY SOLVE'`)
		return nil, -1, constants.EmptyPuzzle
	}

	return &duration, solution.GetNumPlacements(), solution.GetSimple()
}
