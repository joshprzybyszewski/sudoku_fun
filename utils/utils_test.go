package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_presenceOf(t *testing.T) {
	testCases := []struct {
		msg     string
		entry   Entry
		expPres Presence
	}{{
		msg:     `for 1`,
		entry:   1,
		expPres: Presence(uint16(1 << 0)),
	}, {
		msg:     `for 2`,
		entry:   2,
		expPres: Presence(uint16(1 << 1)),
	}, {
		msg:     `for 3`,
		entry:   3,
		expPres: Presence(uint16(1 << 2)),
	}, {
		msg:     `for 4`,
		entry:   4,
		expPres: Presence(uint16(1 << 3)),
	}, {
		msg:     `for 5`,
		entry:   5,
		expPres: Presence(uint16(1 << 4)),
	}, {
		msg:     `for 6`,
		entry:   6,
		expPres: Presence(uint16(1 << 5)),
	}, {
		msg:     `for 7`,
		entry:   7,
		expPres: Presence(uint16(1 << 6)),
	}, {
		msg:     `for 8`,
		entry:   8,
		expPres: Presence(uint16(1 << 7)),
	}, {
		msg:     `for 9`,
		entry:   9,
		expPres: Presence(uint16(1 << 8)),
	}, {
		msg:     `for negative`,
		entry:   -1,
		expPres: Presence(0),
	}, {
		msg:     `for zero`,
		entry:   0,
		expPres: Presence(0),
	}, {
		msg:     `for greater than sideLen`,
		entry:   10,
		expPres: Presence(0),
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("test %v failed!", tc.msg)
		assert.Equal(t, tc.expPres, PresenceOf(tc.entry), failMsg)
	}
}

func Test_getBox(t *testing.T) {
	testCases := []struct {
		row, col int
		expBox   int
		isError  bool
	}{{
		row:    0,
		col:    0,
		expBox: 0,
	}, {
		row:    0,
		col:    1,
		expBox: 0,
	}, {
		row:    0,
		col:    2,
		expBox: 0,
	}, {
		row:    1,
		col:    0,
		expBox: 0,
	}, {
		row:    1,
		col:    1,
		expBox: 0,
	}, {
		row:    1,
		col:    2,
		expBox: 0,
	}, {
		row:    2,
		col:    0,
		expBox: 0,
	}, {
		row:    2,
		col:    1,
		expBox: 0,
	}, {
		row:    2,
		col:    2,
		expBox: 0,
	}, {
		row:    0,
		col:    3,
		expBox: 1,
	}, {
		row:    0,
		col:    4,
		expBox: 1,
	}, {
		row:    0,
		col:    5,
		expBox: 1,
	}, {
		row:    1,
		col:    3,
		expBox: 1,
	}, {
		row:    1,
		col:    4,
		expBox: 1,
	}, {
		row:    1,
		col:    5,
		expBox: 1,
	}, {
		row:    2,
		col:    3,
		expBox: 1,
	}, {
		row:    2,
		col:    4,
		expBox: 1,
	}, {
		row:    2,
		col:    5,
		expBox: 1,
	}, {
		row:    0,
		col:    6,
		expBox: 2,
	}, {
		row:    0,
		col:    7,
		expBox: 2,
	}, {
		row:    0,
		col:    8,
		expBox: 2,
	}, {
		row:    1,
		col:    6,
		expBox: 2,
	}, {
		row:    1,
		col:    7,
		expBox: 2,
	}, {
		row:    1,
		col:    8,
		expBox: 2,
	}, {
		row:    2,
		col:    6,
		expBox: 2,
	}, {
		row:    2,
		col:    7,
		expBox: 2,
	}, {
		row:    2,
		col:    8,
		expBox: 2,
	}, {
		row:    3,
		col:    0,
		expBox: 3,
	}, {
		row:    3,
		col:    1,
		expBox: 3,
	}, {
		row:    3,
		col:    2,
		expBox: 3,
	}, {
		row:    4,
		col:    0,
		expBox: 3,
	}, {
		row:    4,
		col:    1,
		expBox: 3,
	}, {
		row:    4,
		col:    2,
		expBox: 3,
	}, {
		row:    5,
		col:    0,
		expBox: 3,
	}, {
		row:    5,
		col:    1,
		expBox: 3,
	}, {
		row:    5,
		col:    2,
		expBox: 3,
	}, {
		row:    3,
		col:    3,
		expBox: 4,
	}, {
		row:    3,
		col:    4,
		expBox: 4,
	}, {
		row:    3,
		col:    5,
		expBox: 4,
	}, {
		row:    4,
		col:    3,
		expBox: 4,
	}, {
		row:    4,
		col:    4,
		expBox: 4,
	}, {
		row:    4,
		col:    5,
		expBox: 4,
	}, {
		row:    5,
		col:    3,
		expBox: 4,
	}, {
		row:    5,
		col:    4,
		expBox: 4,
	}, {
		row:    5,
		col:    5,
		expBox: 4,
	}, {
		row:    3,
		col:    6,
		expBox: 5,
	}, {
		row:    3,
		col:    7,
		expBox: 5,
	}, {
		row:    3,
		col:    8,
		expBox: 5,
	}, {
		row:    4,
		col:    6,
		expBox: 5,
	}, {
		row:    4,
		col:    7,
		expBox: 5,
	}, {
		row:    4,
		col:    8,
		expBox: 5,
	}, {
		row:    5,
		col:    6,
		expBox: 5,
	}, {
		row:    5,
		col:    7,
		expBox: 5,
	}, {
		row:    5,
		col:    8,
		expBox: 5,
	}, {
		row:    6,
		col:    0,
		expBox: 6,
	}, {
		row:    6,
		col:    1,
		expBox: 6,
	}, {
		row:    6,
		col:    2,
		expBox: 6,
	}, {
		row:    7,
		col:    0,
		expBox: 6,
	}, {
		row:    7,
		col:    1,
		expBox: 6,
	}, {
		row:    7,
		col:    2,
		expBox: 6,
	}, {
		row:    8,
		col:    0,
		expBox: 6,
	}, {
		row:    8,
		col:    1,
		expBox: 6,
	}, {
		row:    8,
		col:    2,
		expBox: 6,
	}, {
		row:    6,
		col:    3,
		expBox: 7,
	}, {
		row:    6,
		col:    4,
		expBox: 7,
	}, {
		row:    6,
		col:    5,
		expBox: 7,
	}, {
		row:    7,
		col:    3,
		expBox: 7,
	}, {
		row:    7,
		col:    4,
		expBox: 7,
	}, {
		row:    7,
		col:    5,
		expBox: 7,
	}, {
		row:    8,
		col:    3,
		expBox: 7,
	}, {
		row:    8,
		col:    4,
		expBox: 7,
	}, {
		row:    8,
		col:    5,
		expBox: 7,
	}, {
		row:    6,
		col:    6,
		expBox: 8,
	}, {
		row:    6,
		col:    7,
		expBox: 8,
	}, {
		row:    6,
		col:    8,
		expBox: 8,
	}, {
		row:    7,
		col:    6,
		expBox: 8,
	}, {
		row:    7,
		col:    7,
		expBox: 8,
	}, {
		row:    7,
		col:    8,
		expBox: 8,
	}, {
		row:    8,
		col:    6,
		expBox: 8,
	}, {
		row:    8,
		col:    7,
		expBox: 8,
	}, {
		row:    8,
		col:    8,
		expBox: 8,
	}, {
		row:     -1,
		col:     8,
		isError: true,
	}, {
		row:     4,
		col:     -1,
		isError: true,
	}, {
		row:     9,
		col:     8,
		isError: true,
	}, {
		row:     6,
		col:     9,
		isError: true,
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("test (%v, %v) = box #%v failed!", tc.row, tc.col, tc.expBox)
		box, err := GetBox(tc.row, tc.col)
		if tc.isError {
			require.Error(t, err, failMsg)
		} else {
			require.NoError(t, err, failMsg)
			assert.Equal(t, tc.expBox, box, failMsg)
		}
	}
}

func Test_isPresent(t *testing.T) {
	testCases := []struct {
		ePresence Presence
		presence  Presence
		expVal    bool
	}{{
		ePresence: PresenceOf(4),
		presence:  Presence(0),
		expVal:    false,
	}, {
		ePresence: PresenceOf(4),
		presence:  Presence(1 << 3),
		expVal:    true,
	}, {
		ePresence: PresenceOf(3),
		presence:  Presence(1<<3) | Presence(1<<4),
		expVal:    false,
	}, {
		ePresence: PresenceOf(2),
		presence:  Presence(1<<0) | Presence(1<<1) | Presence(1<<3) | Presence(1<<4) | Presence(1<<7),
		expVal:    true,
	}, {
		ePresence: PresenceOf(3),
		presence:  Presence(1<<0) | Presence(1<<1) | Presence(1<<3) | Presence(1<<4) | Presence(1<<7),
		expVal:    false,
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("test entry, %v (%09b), in %09b failed!", tc.ePresence, tc.ePresence, tc.presence)
		assert.Equal(t, tc.expVal, IsPresent(tc.presence, tc.ePresence), failMsg)
	}
}

func Test_getPossibleEntries(t *testing.T) {
	testCases := []struct {
		msg           string
		entries       []Entry
		presence      Presence
		expPosEntries []Entry
	}{{
		msg:           `Where 5 is present`,
		entries:       AllEntries,
		presence:      PresenceOf(5),
		expPosEntries: []Entry{1, 2, 3, 4, 6, 7, 8, 9},
	}, {
		msg:           `Where 5, 6, and 7 are present`,
		entries:       AllEntries,
		presence:      PresenceOf(5) | PresenceOf(6) | PresenceOf(7),
		expPosEntries: []Entry{1, 2, 3, 4, 8, 9},
	}, {
		msg:           `Where 1, 2, 8, and 9 are present`,
		entries:       AllEntries,
		presence:      PresenceOf(1) | PresenceOf(2) | PresenceOf(8) | PresenceOf(9),
		expPosEntries: []Entry{3, 4, 5, 6, 7},
	}, {
		msg:           `Where all entries are present`,
		entries:       AllEntries,
		presence:      FullPresence, //presenceOf(1) | presenceOf(2) | presenceOf(3) | presenceOf(4) | presenceOf(5) | presenceOf(6) | presenceOf(7) | presenceOf(8) | presenceOf(9),
		expPosEntries: []Entry{},
	}, {
		msg:           `Where no entries are present`,
		entries:       AllEntries,
		presence:      EmptyPresence,
		expPosEntries: AllEntries,
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("test %v failed!", tc.msg)
		assert.Equal(t, tc.expPosEntries, GetPossibleEntries(tc.entries, tc.presence), failMsg)
	}
}

func Test_getPossibleEntriesQuickly(t *testing.T) {
	testCases := []struct {
		msg           string
		entries       []Entry
		rp            Presence
		cp            Presence
		bp            Presence
		expPosEntries []Entry
	}{{
		msg:           `Where 5 is present`,
		entries:       AllEntries,
		rp:            PresenceOf(5),
		cp:            PresenceOf(5),
		bp:            PresenceOf(5),
		expPosEntries: []Entry{1, 2, 3, 4, 6, 7, 8, 9},
	}, {
		msg:           `Where 5, 6, and 7 are present`,
		entries:       AllEntries,
		rp:            PresenceOf(5) | PresenceOf(6) | PresenceOf(7),
		cp:            PresenceOf(5),
		bp:            PresenceOf(5),
		expPosEntries: []Entry{1, 2, 3, 4, 8, 9},
	}, {
		msg:           `Where 1, 2, 8, and 9 are present`,
		entries:       AllEntries,
		rp:            PresenceOf(1) | PresenceOf(8),
		cp:            PresenceOf(2),
		bp:            PresenceOf(9),
		expPosEntries: []Entry{3, 4, 5, 6, 7},
	}, {
		msg:           `Where all entries are present`,
		entries:       AllEntries,
		rp:            PresenceOf(1) | PresenceOf(2) | PresenceOf(3) | PresenceOf(6) | PresenceOf(7) | PresenceOf(9),
		cp:            PresenceOf(4),
		bp:            PresenceOf(5) | PresenceOf(8),
		expPosEntries: []Entry{},
	}, {
		msg:           `Where no entries are present`,
		entries:       AllEntries,
		rp:            EmptyPresence,
		cp:            EmptyPresence,
		bp:            EmptyPresence,
		expPosEntries: AllEntries,
	}}

	for _, tc := range testCases {
		failMsg := fmt.Sprintf("test %v failed!", tc.msg)
		assert.Equal(t, tc.expPosEntries, GetPossibleEntriesQuickly(tc.rp, tc.cp, tc.bp), failMsg)
	}
}
