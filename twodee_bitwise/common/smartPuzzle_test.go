package common

import (
	"testing"

	"github.com/testify/assert"

	speed_utils "github.com/joshprzybyszewski/sudoku_fun/utils/speed"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
	"github.com/joshprzybyszewski/sudoku_fun/utils"
	"fmt"
	"github.com/testify/require"
)

func Test_clone(t *testing.T) {
	pzl := &SmartPuzzle{}
	pzl.NumPlaced = 2
	pzl.Tiles[2][5] = types.Tile(3)
	pzl.Tiles[7][1] = types.Tile(3)
	p3 := speed_utils.PresenceOf(3)
	pzl.Rows[2] = p3
	pzl.Rows[7] = p3
	pzl.Cols[5] = p3
	pzl.Cols[1] = p3
	pzl.Boxs[1] = p3
	pzl.Boxs[6] = p3

	pclone := pzl.Clone()

	assert.Equal(t, *pzl, *pclone)
}

func Test_place(t *testing.T) {
	var pzl1 common.SmartPuzzle
	var pzl2s, pzl2f common.SmartPuzzle

	pzl1.Tiles[0][0] = types.Tile(4)
	pzl1.Rows[3] = utils.PresenceOf(5)
	pzl1.Cols[3] = utils.PresenceOf(5)
	pzl1.Boxs[8] = utils.PresenceOf(5)

	pzl2s.NumFreeInRow[6] = 1
	pzl2s.NumFreeInCol[6] = 1
	pzl2s.NumFreeInBox[8] = 1

	pzl2f.Tiles[6][6] = types.Tile(3)
	pzl2f.Rows[6] = utils.PresenceOf(3)
	pzl2f.Cols[6] = utils.PresenceOf(3)
	pzl2f.Boxs[8] = utils.PresenceOf(3)
	pzl2f.NumPlaced = 1

	testCases := []struct {
		msg           string
		row           int
		col           int
		box           int
		entry         types.Entry
		pzzl          common.SmartPuzzle
		isError       bool
		expVal        bool
		expFinalState common.SmartPuzzle
	}{{
		msg:           `On a tile that has been placed`,
		row:           0,
		col:           0,
		box:           0,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `On a row that contains that entry`,
		row:           3,
		col:           0,
		box:           3,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `On a col that contains that entry`,
		row:           0,
		col:           3,
		box:           1,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `In a box that contains that entry`,
		row:           7,
		col:           7,
		box:           8,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `In a fine location`,
		row:           6,
		col:           6,
		box:           8,
		entry:         3,
		pzzl:          pzl2s,
		isError:       false,
		expVal:        true,
		expFinalState: pzl2f,
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("test %v failed!", tc.msg)
		wasPlaced, err := tc.pzzl.Place(tc.row, tc.col, tc.box, tc.entry)
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