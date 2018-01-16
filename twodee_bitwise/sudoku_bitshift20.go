package twodee_bitwise

import (
	"fmt"
	"time"

	//"./naive"
	"./smart"
	"../utils"
	"../utils/constants"
	"../utils/types"
)

func PuzzleSolver(sudokuPuzzle string) (*time.Duration, int, string) {
	puzzle, err := readSudoku(sudokuPuzzle)
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

func readSudoku(entries string) (s types.Sudoku, err error) {
	pzl, err := smart.ReadSudoku(entries)
	if err != nil {
		return nil, err
	}

	return types.Sudoku(pzl), nil
}
