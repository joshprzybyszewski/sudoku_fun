package smart

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/joshprzybyszewski/sudoku_fun/twodee_bitwise/common"
	brute "github.com/joshprzybyszewski/sudoku_fun/utils"
	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
	utils "github.com/joshprzybyszewski/sudoku_fun/utils/speed"
	"github.com/joshprzybyszewski/sudoku_fun/utils/test_utils"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

func Test_SmartSolve(t *testing.T) {
	emptyPuzzle := common.SmartPuzzle{}
	emptyPuzzle.Solver = Solve
	emptyPuzzle.NumFreeInRow = [constants.SideLen]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}
	emptyPuzzle.NumFreeInCol = [constants.SideLen]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}
	emptyPuzzle.NumFreeInBox = [constants.SideLen]uint8{9, 9, 9, 9, 9, 9, 9, 9, 9}

	sparsePuzzle := common.SmartPuzzle{}
	sparsePuzzle.NumPlaced = 5
	sparsePuzzle.Solver = Solve
	sparsePuzzle.Tiles[0][4] = types.Tile(2)
	sparsePuzzle.Tiles[0][8] = types.Tile(9)
	sparsePuzzle.Tiles[1][1] = types.Tile(6)
	sparsePuzzle.Tiles[8][7] = types.Tile(5)
	sparsePuzzle.Tiles[8][8] = types.Tile(6)
	sparsePuzzle.Rows[0] = utils.PresenceOf(2) | utils.PresenceOf(9)
	sparsePuzzle.Rows[1] = utils.PresenceOf(6)
	sparsePuzzle.Rows[8] = utils.PresenceOf(5) | utils.PresenceOf(6)
	sparsePuzzle.Cols[1] = utils.PresenceOf(6)
	sparsePuzzle.Cols[4] = utils.PresenceOf(2)
	sparsePuzzle.Cols[7] = utils.PresenceOf(5)
	sparsePuzzle.Cols[8] = utils.PresenceOf(6) | utils.PresenceOf(9)
	sparsePuzzle.Boxs[0] = utils.PresenceOf(6)
	sparsePuzzle.Boxs[1] = utils.PresenceOf(2)
	sparsePuzzle.Boxs[2] = utils.PresenceOf(9)
	sparsePuzzle.Boxs[8] = utils.PresenceOf(5) | utils.PresenceOf(6)
	sparsePuzzle.NumFreeInRow = [constants.SideLen]uint8{7, 8, 9, 9, 9, 9, 9, 9, 7}
	sparsePuzzle.NumFreeInCol = [constants.SideLen]uint8{9, 8, 9, 9, 8, 9, 9, 8, 7}
	sparsePuzzle.NumFreeInBox = [constants.SideLen]uint8{8, 8, 8, 9, 9, 9, 9, 9, 7}

	solvedPuzzle := common.SmartPuzzle{
		81,
		Solve,
		[constants.SideLen][constants.SideLen]types.Tile{{3, 8, 4, 1, 7, 9, 6, 5, 2}, {2, 7, 6, 3, 5, 8, 4, 1, 9}, {1, 5, 9, 4, 6, 2, 8, 7, 3}, {9, 6, 1, 7, 8, 5, 2, 3, 4}, {4, 2, 7, 9, 1, 3, 5, 6, 8}, {5, 3, 8, 6, 2, 4, 7, 9, 1}, {6, 9, 2, 5, 4, 1, 3, 8, 7}, {7, 4, 3, 8, 9, 6, 1, 2, 5}, {8, 1, 5, 2, 3, 7, 9, 4, 6}},
		[constants.SideLen]types.Presence{constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence},
		[constants.SideLen]types.Presence{constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence},
		[constants.SideLen]types.Presence{constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence, constants.FullPresence},
		[constants.SideLen]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[constants.SideLen]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0},
		[constants.SideLen]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	completelyEmptyPuzzle := common.SmartPuzzle{}
	completelyEmptyPuzzle.Solver = Solve

	testCases := []struct {
		msg           string
		pzlStr        string
		isError       bool
		expFinalState *common.SmartPuzzle
	}{{
		msg:           `Empty Puzzle`,
		pzlStr:        constants.EmptyPuzzle,
		isError:       false,
		expFinalState: &emptyPuzzle,
	}, {
		msg:           `Sparse Puzzle`,
		pzlStr:        `....2...9.6....................................................................56`,
		isError:       false,
		expFinalState: &sparsePuzzle,
	}, {
		msg:           `Solved Puzzle`,
		pzlStr:        `384179652276358419159462873961785234427913568538624791692541387743896125815237946`,
		isError:       false,
		expFinalState: &solvedPuzzle,
	}, {
		msg:     `Errored Puzzle`,
		pzlStr:  `11...............................................................................`,
		isError: true,
	}}

	for _, tc := range testCases {
		failMsg := ``
		justReadPzl, err := common.GetSmartPuzzle(tc.pzlStr, Solve)
		if tc.isError {
			assert.Error(t, err, failMsg)
		} else {
			require.NoError(t, err, failMsg)
			assert.True(t, test_utils.SmartPuzzlesAreEqual(tc.expFinalState, &justReadPzl), failMsg)
		}
	}
}

func Test_SolveAndHelpers(t *testing.T) {
	testCases := []struct {
		msg          string
		pzlStr       string
		solutionStr  string
		startingSize int
		firstPlaceR  int
		firstPlaceC  int
		firstPlaceB  int
		firstPlaceEs []types.Entry
	}{{
		msg:          `First Puzzle`,
		pzlStr:       `..812...9.6.........2..95......8.93....2..68...........564..3....9...41..8..1..56`,
		solutionStr:  `348125769965347821712869543621784935573291684894536172156472398239658417487913256`,
		startingSize: 25,
		firstPlaceR:  0,
		firstPlaceC:  6,
		firstPlaceB:  2,
		firstPlaceEs: []types.Entry{7},
	}, {
		msg:          `Second Puzzle`,
		pzlStr:       `...21.83.3.1..5....82.7...54....2..9.78.....4.......1.71...........5.3.1...8..9..`,
		solutionStr:  `957214836341685297682379145435162789178593624296748513713926458829457361564831972`,
		startingSize: 25,
		firstPlaceR:  0,
		firstPlaceC:  8,
		firstPlaceB:  2,
		firstPlaceEs: []types.Entry{6, 7},
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("Failed %v", tc.msg)

		pzl, err := common.GetSmartPuzzle(tc.pzlStr, Solve)
		require.NoError(t, err, failMsg)

		assert.Equal(t, tc.startingSize, pzl.GetNumPlacements(), failMsg)
		assert.Equal(t, tc.pzlStr, pzl.GetSimple(), failMsg)

		emptyR, emptyC, emptyB, err := getEmptyTile(&pzl)
		require.NoError(t, err, failMsg)
		assert.Equal(t, tc.firstPlaceR, emptyR, failMsg)
		assert.Equal(t, tc.firstPlaceC, emptyC, failMsg)
		assert.Equal(t, tc.firstPlaceB, emptyB, failMsg)

		locR, locC, locB, es, err := getLocationAndEntries(&pzl)
		require.NoError(t, err, failMsg)
		assert.Equal(t, tc.firstPlaceEs, es, failMsg)
		assert.Equal(t, emptyR, locR, failMsg)
		assert.Equal(t, emptyC, locC, failMsg)
		assert.Equal(t, emptyB, locB, failMsg)

		solution, err := pzl.Solve()
		require.NoError(t, err, failMsg)
		require.NotNil(t, solution, failMsg)
		assert.True(t, pzl.GetNumPlacements() >= 81-tc.startingSize, failMsg)

		wasSolved := brute.BruteForceCheck(solution.GetSimple())
		assert.True(t, wasSolved, failMsg)

		assert.Equal(t, tc.solutionStr, solution.GetSimple(), failMsg)
	}
}
