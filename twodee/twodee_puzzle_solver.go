package twodee

import (
	"errors"
	"fmt"
	"time"

	"github.com/joshprzybyszewski/sudoku_fun/utils"
	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

type SudokuReader func(str string) (p types.Sudoku, err error)

func PuzzleSolver(readPuzzle SudokuReader, sudokuPuzzle string) (*time.Duration, int, string, error) {
	puzzle, err := readPuzzle(sudokuPuzzle)
	if err != nil {
		return nil, -1, ``, errors.New(fmt.Sprintf("Bad Sudoku: %v", err.Error()))
	}

	t0 := time.Now()
	solution, err := puzzle.Solve()
	duration := time.Since(t0)

	if err != nil {
		return nil, -1, constants.EmptyPuzzle, errors.New(fmt.Sprintf("Couldn't Solve: %v", err.Error()))
	}

	if !utils.BruteForceCheck(solution.GetSimple()) {
		return nil, -1, constants.EmptyPuzzle, errors.New(fmt.Sprintf("It thinks it solved, but it didn't!\n%v", solution.GetSimple()))
	}

	return &duration, solution.GetNumPlacements(), solution.GetSimple(), nil
}
