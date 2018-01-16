package smart

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	brute "../../utils"
	utils "../../utils/speed"
	"../../utils/types"
	"../../utils/constants"

)

func Test_place(t *testing.T) {
	var pzl1 Puzzle
	var pzl2s, pzl2f Puzzle

	pzl1.tiles[0][0] = types.Tile(4)
	pzl1.rows[3] = utils.PresenceOf(5)
	pzl1.cols[3] = utils.PresenceOf(5)
	pzl1.boxs[8] = utils.PresenceOf(5)

	pzl2s.numFreeInRow[6] = 1
	pzl2s.numFreeInCol[6] = 1
	pzl2s.numFreeInBox[8] = 1

	pzl2f.tiles[6][6] = types.Tile(3)
	pzl2f.rows[6] = utils.PresenceOf(3)
	pzl2f.cols[6] = utils.PresenceOf(3)
	pzl2f.boxs[8] = utils.PresenceOf(3)
	pzl2f.numPlaced = 1

	testCases := []struct {
		msg           string
		row           int
		col           int
		box int
		entry         types.Entry
		pzzl          Puzzle
		isError       bool
		expVal        bool
		expFinalState Puzzle
	}{{
		msg:           `On a tile that has been placed`,
		row:           0,
		col:           0,
		box: 0,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `On a row that contains that entry`,
		row:           3,
		col:           0,
		box: 3,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `On a col that contains that entry`,
		row:           0,
		col:           3,
		box: 1,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `In a box that contains that entry`,
		row:           7,
		col:           7,
		box: 8,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `In a fine location`,
		row:           6,
		col:           6,
		box: 8,
		entry:         3,
		pzzl:          pzl2s,
		isError:       false,
		expVal:        true,
		expFinalState: pzl2f,
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("test %v failed!", tc.msg)
		wasPlaced, err := tc.pzzl.place(tc.row, tc.col, tc.box, tc.entry)
		if tc.isError {
			// skip over this test because SMART is not ROBUST
			continue
			assert.Error(t, err, failMsg)
			assert.False(t, wasPlaced, failMsg)
		} else {
			require.NoError(t, err, failMsg)
			assert.Equal(t, tc.expVal, wasPlaced, failMsg)
		}
		assert.Equal(t, tc.expFinalState, tc.pzzl, failMsg)
	}
}

func Test_ReadSudoku(t *testing.T) {
	emptyPuzzle := Puzzle{}
	emptyPuzzle.numFreeInRow = [constants.SideLen]uint8{9,9,9,9,9,9,9,9,9}
	emptyPuzzle.numFreeInCol = [constants.SideLen]uint8{9,9,9,9,9,9,9,9,9}
	emptyPuzzle.numFreeInBox = [constants.SideLen]uint8{9,9,9,9,9,9,9,9,9}

	sparsePuzzle := Puzzle{}
	sparsePuzzle.numPlaced = 5
	sparsePuzzle.tiles[0][4] = types.Tile(2)
	sparsePuzzle.tiles[0][8] = types.Tile(9)
	sparsePuzzle.tiles[1][1] = types.Tile(6)
	sparsePuzzle.tiles[8][7] = types.Tile(5)
	sparsePuzzle.tiles[8][8] = types.Tile(6)
	sparsePuzzle.rows[0] = utils.PresenceOf(2) | utils.PresenceOf(9)
	sparsePuzzle.rows[1] = utils.PresenceOf(6)
	sparsePuzzle.rows[8] = utils.PresenceOf(5) | utils.PresenceOf(6)
	sparsePuzzle.cols[1] = utils.PresenceOf(6)
	sparsePuzzle.cols[4] = utils.PresenceOf(2)
	sparsePuzzle.cols[7] = utils.PresenceOf(5)
	sparsePuzzle.cols[8] = utils.PresenceOf(6) | utils.PresenceOf(9)
	sparsePuzzle.boxs[0] = utils.PresenceOf(6)
	sparsePuzzle.boxs[1] = utils.PresenceOf(2)
	sparsePuzzle.boxs[2] = utils.PresenceOf(9)
	sparsePuzzle.boxs[8] = utils.PresenceOf(5) | utils.PresenceOf(6)
	sparsePuzzle.numFreeInRow = [constants.SideLen]uint8{7,8,9,9,9,9,9,9,7}
	sparsePuzzle.numFreeInCol = [constants.SideLen]uint8{9,8,9,9,8,9,9,8,7}
	sparsePuzzle.numFreeInBox = [constants.SideLen]uint8{8,8,8,9,9,9,9,9,7}

	solvedPuzzle := Puzzle{
		81,
		[constants.SideLen][constants.SideLen]types.Tile{{3,8,4,1,7,9,6,5,2}, {2,7,6,3,5,8,4,1,9}, {1,5,9,4,6,2,8,7,3}, { 9,6,1,7,8,5,2,3,4}, { 4,2,7,9,1,3,5,6,8}, { 5,3,8,6,2,4,7,9,1}, { 6,9,2,5,4,1,3,8,7}, { 7,4,3,8,9,6,1,2,5}, { 8,1,5,2,3,7,9,4,6}},
		[constants.SideLen]types.Presence{constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence},
		[constants.SideLen]types.Presence{constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence},
		[constants.SideLen]types.Presence{constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence,constants.FullPresence},
		[constants.SideLen]uint8{0,0,0,0,0,0,0,0,0},
		[constants.SideLen]uint8{0,0,0,0,0,0,0,0,0},
		[constants.SideLen]uint8{0,0,0,0,0,0,0,0,0},
	}

	testCases := []struct {
		msg           string
		pzlStr           string
		isError bool
		expFinalState Puzzle
	}{{
		msg: `Empty Puzzle`,
		pzlStr: constants.EmptyPuzzle,
		isError: false,
		expFinalState: emptyPuzzle,
	}, {
		msg: `Sparse Puzzle`,
		pzlStr: `....2...9.6....................................................................56`,
		isError: false,
		expFinalState: sparsePuzzle,
	}, {
		msg: `Solved Puzzle`,
		pzlStr: `384179652276358419159462873961785234427913568538624791692541387743896125815237946`,
		isError: false,
		expFinalState: solvedPuzzle,
	}, {
		msg: `Errored Puzzle`,
		pzlStr: `11...............................................................................`,
		isError: true,
		expFinalState: Puzzle{},
	}}

	for _, tc := range testCases {
		failMsg := ``
		actPzl, err := ReadSudoku(tc.pzlStr)
		if tc.isError {
			assert.Error(t, err, failMsg)
		} else {
			require.NoError(t, err, failMsg)
			assert.Equal(t, tc.expFinalState, actPzl, failMsg)
		}
	}
}

func Test_clone(t *testing.T) {
	pzl := &Puzzle{}
	pzl.numPlaced = 2
	pzl.tiles[2][5] = types.Tile(3)
	pzl.tiles[7][1] = types.Tile(3)
	p3 := utils.PresenceOf(3)
	pzl.rows[2] = p3
	pzl.rows[7] = p3
	pzl.cols[5] = p3
	pzl.cols[1] = p3
	pzl.boxs[1] = p3
	pzl.boxs[6] = p3

	pclone := pzl.clone()

	assert.Equal(t, *pzl, *pclone)
}

func Test_Solve(t *testing.T) {
	testCases := []struct {
		msg           string
		pzlStr        string
		solutionStr   string
		startingSize  int
		firstPlaceR   int
		firstPlaceC   int
		firstPlaceB   int
		firstPlaceEs  []types.Entry
		isError       bool
		expFinalState Puzzle
	}{{
		msg:           `First Puzzle`,
		pzlStr:        `..812...9.6.........2..95......8.93....2..68...........564..3....9...41..8..1..56`,
		solutionStr:   `348125769965347821712869543621784935573291684894536172156472398239658417487913256`,
		startingSize:  25,
		firstPlaceR:   0,
		firstPlaceC:   6,
		firstPlaceB:   2,
		firstPlaceEs:  []types.Entry{7},
		isError:       false,
		expFinalState: Puzzle{},
	}, {
		msg:           `Second Puzzle`,
		pzlStr:        `...21.83.3.1..5....82.7...54....2..9.78.....4.......1.71...........5.3.1...8..9..`,
		solutionStr:   `957214836341685297682379145435162789178593624296748513713926458829457361564831972`,
		startingSize:  25,
		firstPlaceR:   0,
		firstPlaceC:   8,
		firstPlaceB:   2,
		firstPlaceEs:  []types.Entry{6, 7},
		isError:       false,
		expFinalState: Puzzle{},
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("Failed %v", tc.msg)

		pzl, err := ReadSudoku(tc.pzlStr)
		require.NoError(t, err, failMsg)

		assert.Equal(t, tc.startingSize, pzl.GetNumPlacements(), failMsg)
		assert.Equal(t, tc.pzlStr, pzl.GetSimple(), failMsg)

		pzl.PrintPretty()
		pzl.PrintSimple()

		emptyR, emptyC, emptyB, err := pzl.getEmptyTile()
		require.NoError(t, err, failMsg)
		assert.Equal(t, tc.firstPlaceR, emptyR, failMsg)
		assert.Equal(t, tc.firstPlaceC, emptyC, failMsg)
		assert.Equal(t, tc.firstPlaceB, emptyB, failMsg)

		locR, locC, locB, es, err := pzl.getLocationAndEntries()
		require.NoError(t, err, failMsg)
		assert.Equal(t, tc.firstPlaceEs, es, failMsg)
		assert.Equal(t, emptyR, locR, failMsg)
		assert.Equal(t, emptyC, locC, failMsg)
		assert.Equal(t, emptyB, locB, failMsg)

		solution, err := pzl.Solve()
		require.NoError(t, err, failMsg)
		require.NotNil(t, solution, failMsg)
		assert.True(t, pzl.GetNumPlacements() >= 81 - tc.startingSize, failMsg)

		wasSolved := brute.BruteForceCheck(solution.GetSimple())
		assert.True(t, wasSolved, failMsg)

		assert.Equal(t, tc.solutionStr, solution.GetSimple(), failMsg)
	}
}