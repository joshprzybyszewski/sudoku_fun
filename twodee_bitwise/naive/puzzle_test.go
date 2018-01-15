package naive

import (
	"../../utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_place(t *testing.T) {
	var pzl1 Puzzle
	var pzl2s, pzl2f Puzzle

	pzl1.tiles[0][0] = utils.Tile(4)
	pzl1.rows[3] = utils.PresenceOf(5)
	pzl1.cols[3] = utils.PresenceOf(5)
	pzl1.boxs[8] = utils.PresenceOf(5)

	pzl2f.tiles[6][6] = utils.Tile(3)
	pzl2f.rows[6] = utils.PresenceOf(3)
	pzl2f.cols[6] = utils.PresenceOf(3)
	pzl2f.boxs[8] = utils.PresenceOf(3)
	pzl2f.numPlaced = 1

	testCases := []struct {
		msg           string
		row           int
		col           int
		entry         utils.Entry
		pzzl          Puzzle
		isError       bool
		expVal        bool
		expFinalState Puzzle
	}{{
		msg:           `On a tile that has been placed`,
		row:           0,
		col:           0,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `On a row that contains that entry`,
		row:           3,
		col:           0,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `On a col that contains that entry`,
		row:           0,
		col:           3,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `In a box that contains that entry`,
		row:           7,
		col:           7,
		entry:         5,
		pzzl:          pzl1,
		isError:       true,
		expFinalState: pzl1,
	}, {
		msg:           `In a fine location`,
		row:           6,
		col:           6,
		entry:         3,
		pzzl:          pzl2s,
		isError:       false,
		expVal:        true,
		expFinalState: pzl2f,
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("test %v failed!", tc.msg)
		wasPlaced, err := tc.pzzl.place(tc.row, tc.col, tc.entry)
		if tc.isError {
			assert.Error(t, err, failMsg)
			assert.False(t, wasPlaced, failMsg)
		} else {
			require.NoError(t, err, failMsg)
			assert.Equal(t, tc.expVal, wasPlaced, failMsg)
		}
		assert.Equal(t, tc.expFinalState, tc.pzzl, failMsg)
	}
}
