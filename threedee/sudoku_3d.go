package threedee

import (
	"github.com/joshprzybyszewski/sudoku_fun/threedee/robust"
	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
)

func FindPuzzle() (bool, error) {
	puzzle, err  := robust.ReadSudoku(constants.Empty3dPuzzle)
	if err != nil {
		println(`puzzle not read!`)
		return false, err
	}
	println(`puzzle read!`)

	t, err := puzzle.Solve()
	if err != nil {
		println(`puzzle not solved`)
		return false, err
	}
	t.PrintPretty()

	return true, nil
}