package twodee_bitwise

import (
//"testing"
//
//"github.com/stretchr/testify/assert"
//"github.com/stretchr/testify/require"
//
//"fmt"
)

//func Test_getLeastLocations(t *testing.T) {
//	lr1lc3lb1 := `...123...12345678....78.......2........5........8........3........6..............`
//	lr8lc8lb8 := `........8........7........6........5........4........3......432......56112345678.`
//	lr3lc3lb3 := `...8........2........3.....12.4567894567.....7891........5........6..............`
//	lr0lc0lb1 := `1234.....456......789............................................................`
//	lr0lc1lb6 := `1234.....456......789........1........4........7........2.......35........8......`
//
//	//.........
//	//.........
//	//.........
//	//.........
//	//.........
//	//.........
//	//.........
//	//.........
//	//.........
//
//
//	testCases := []struct {
//		msg    string
//		pzlstr string
//		ebr    int
//		ebc    int
//		ebb    int
//	}{{
//		msg:    `lr: 1; lc: 3; lb: 1;`,
//		pzlstr: lr1lc3lb1,
//		ebr:    1,
//		ebc:    3,
//		ebb:    1,
//	}, {
//		msg:    `lr: 8; lc: 8; lb: 8;`,
//		pzlstr: lr8lc8lb8,
//		ebr:    8,
//		ebc:    8,
//		ebb:    8,
//	}, {
//		msg:    `lr: 3; lc: 3; lb: 3;`,
//		pzlstr: lr3lc3lb3,
//		ebr:    3,
//		ebc:    3,
//		ebb:    3,
//	}, {
//		msg:    `lr: 0; lc: 0; lb: 1; has full box 0`,
//		pzlstr: lr0lc0lb1,
//		ebr:    0,
//		ebc:    0,
//		ebb:    1,
//	}, {
//		msg:    `lr: 0; lc: 1; lb: 6; has full box 0 and full col 2`,
//		pzlstr: lr0lc1lb6,
//		ebr:    0,
//		ebc:    1,
//		ebb:    6,
//	}}
//
//	for _, tc := range testCases {
//		puzzle, err := readSudoku(tc.pzlstr)
//
//		if err != nil {
//			println(`josh youre an idiot `, tc.msg)
//			continue
//		}
//
//		abr, abc, abb := puzzle.getLeastLocations()
//		assert.Equal(t, tc.ebr, abr, "%v\tROW expected %d but recieved %d", tc.msg, tc.ebr, abr)
//		assert.Equal(t, tc.ebc, abc, "%v\tCOL expected %d but recieved %d", tc.msg, tc.ebc, abc)
//		assert.Equal(t, tc.ebb, abb, "%v\tBOX expected %d but recieved %d", tc.msg, tc.ebb, abb)
//	}
//}
