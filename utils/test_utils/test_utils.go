package test_utils

import (
	"github.com/joshprzybyszewski/sudoku_fun/twodee_bitwise/common"
	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	"fmt"
)

func SmartPuzzlesAreEqual(exp, act *common.SmartPuzzle) bool {
	if exp.NumPlaced != act.NumPlaced {
		println(fmt.Sprintf("NumPlaced expected: %v, actual: %v", exp.NumPlaced, act.NumPlaced))
		return false
	}

	//if exp.Solver == act.Solver {
	//	println(fmt.Sprintf("Solver expected: %v, actual: %v", &exp.Solver, &act.Solver))
	//	return false
	//}

	for i := 0; i < constants.SideLen; i++ {
		if exp.Rows[i] != act.Rows[i] {
			println(fmt.Sprintf("Rows[%v] expected: %v, actual: %v", i, exp.Rows[i], act.Rows[i]))
			return false
		}
		if exp.Cols[i] != act.Cols[i] {
			println(fmt.Sprintf("Cols[%v] expected: %v, actual: %v", i, exp.Cols[i], act.Cols[i]))
			return false
		}
		if exp.Boxs[i] != act.Boxs[i] {
			println(fmt.Sprintf("Boxs[%v] expected: %v, actual: %v", i, exp.Boxs[i], act.Boxs[i]))
			return false
		}
		if exp.NumFreeInRow[i] != act.NumFreeInRow[i] {
			println(fmt.Sprintf("NumFreeInRow[%v] expected: %v, actual: %v", i, exp.NumFreeInRow[i], act.NumFreeInRow[i]))
			return false
		}
		if exp.NumFreeInCol[i] != act.NumFreeInCol[i] {
			println(fmt.Sprintf("NumFreeInCol[%v] expected: %v, actual: %v", i, exp.NumFreeInCol[i], act.NumFreeInCol[i]))
			return false
		}
		if exp.NumFreeInBox[i] != act.NumFreeInBox[i] {
			println(fmt.Sprintf("NumFreeInBox[%v] expected: %v, actual: %v", i, exp.NumFreeInBox[i], act.NumFreeInBox[i]))
			return false
		}
		for j := 0; j < constants.SideLen; j++ {
			if exp.Tiles[i][j] != act.Tiles[i][j] {
				println(fmt.Sprintf("Tiles[%v][%v] expected: %v, actual: %v", i, j, exp.Tiles[i][j], act.Tiles[i][j]))
				return false
			}
		}
	}

	return true
}