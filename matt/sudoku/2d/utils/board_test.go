package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidMove(t *testing.T) {
	b := NewGame()
	b2, ok := b.Set(1, 1, 1)
	assert.True(t, ok)
	assert.NotEqual(t, b2, b)

	_, ok = b.Set(1, 1, 1)
	assert.True(t, ok)

	b3, ok := b2.Set(1, 1, 1)
	assert.False(t, ok)
	assert.Nil(t, b3)

	assert.False(t, b.Complete())
	assert.False(t, b2.Complete())
}

func TestInvalidMove(t *testing.T) {
	b := NewGame()

	_, ok := b.Set(1, 1, 10)
	assert.False(t, ok)
	_, ok = b.Set(1, 1, 0)
	assert.False(t, ok)

	_, ok = b.Set(1, 10, 0)
	assert.False(t, ok)
	_, ok = b.Set(1, -1, 0)
	assert.False(t, ok)

	_, ok = b.Set(10, 1, 0)
	assert.False(t, ok)
	_, ok = b.Set(-1, 1, 0)
	assert.False(t, ok)

	// make a valid move
	b, ok = b.Set(1, 1, 1)
	assert.True(t, ok)
	assert.NotNil(t, b)

	_, ok = b.Set(1, 1, 1)
	assert.False(t, ok)
	_, ok = b.Set(1, 1, 0)
	assert.False(t, ok)

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			_, ok = b.Set(r, c, 1)
			assert.False(t, ok)
		}
	}

	for c := 0; c < 9; c++ {
		_, ok = b.Set(1, c, 1)
		assert.False(t, ok)
	}
	for r := 0; r < 9; r++ {
		_, ok = b.Set(r, 1, 1)
		assert.False(t, ok)
	}
}

func TestNewWithState(t *testing.T) {
	b := NewGameWithState(Board{{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	}, {
		4, 5, 6, 7, 8, 9, 1, 2, 3,
	}, {
		7, 8, 9, 1, 2, 3, 4, 5, 6,
	}, {
		2, 3, 4, 5, 6, 7, 8, 9, 1,
	}, {
		5, 6, 7, 8, 9, 1, 2, 3, 4,
	}, {
		8, 9, 1, 2, 3, 4, 5, 6, 7,
	}, {
		3, 4, 5, 6, 7, 8, 9, 1, 2,
	}, {
		6, 7, 8, 9, 1, 2, 3, 4, 5,
	}, {
		9, 1, 2, 3, 4, 5, 6, 7, 8,
	}})
	assert.True(t, b.Complete())

	b = NewGameWithState(Board{{
		1, 2, 3, 4, 5, 6, 7, 8, 0,
	}, {
		4, 5, 6, 7, 8, 9, 1, 2, 3,
	}, {
		7, 8, 9, 1, 2, 3, 4, 5, 6,
	}, {
		2, 3, 4, 5, 6, 7, 8, 9, 1,
	}, {
		5, 6, 7, 8, 9, 1, 2, 3, 4,
	}, {
		8, 9, 1, 2, 3, 4, 5, 6, 7,
	}, {
		3, 4, 5, 6, 7, 8, 9, 1, 2,
	}, {
		6, 7, 8, 9, 1, 2, 3, 4, 5,
	}, {
		9, 1, 2, 3, 4, 5, 6, 7, 8,
	}})
	assert.False(t, b.Complete())

	for i := 1; i < 9; i++ {
		_, ok := b.Set(0, 8, i)
		assert.False(t, ok)
	}

	rowGuess, colGuess, valuesGuess := b.NextGuessMove()
	assert.Equal(t, 0, rowGuess)
	assert.Equal(t, 8, colGuess)
	assert.Equal(t, []int{9}, valuesGuess)

	b2, ok := b.Set(rowGuess, colGuess, valuesGuess[0])
	assert.True(t, ok)
	assert.NotNil(t, b2)
	assert.True(t, b2.Complete())
}

func TestNewWithStateHitMap(t *testing.T) {
	b := NewGameWithState(Board{{1, 2, 3}})
	assert.False(t, b.Complete())
	assert.Equal(t, oneMapping|twoMapping|threeMapping, b.subSquareSets[0][0])
	assert.Equal(t, oneMapping|twoMapping|threeMapping, b.rowsSet[0])
	assert.Equal(t, oneMapping, b.colsSet[0])
	assert.Equal(t, twoMapping, b.colsSet[1])
	assert.Equal(t, threeMapping, b.colsSet[2])

	for i := 1; i < 9; i++ {
		assert.Equal(t, emptyMapping, b.rowsSet[i])
	}
	for i := 3; i < 9; i++ {
		assert.Equal(t, emptyMapping, b.colsSet[i])
	}

	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			assert.Equal(t, emptyMapping, b.subSquareSets[i][j])
		}
	}

	for j := 0; j < 9; j++ {
		assert.Equal(t, oneMapping|twoMapping|threeMapping, b.individualSets[0][j])
	}

	for i := 1; i < 9; i++ {
		if i < 3 {
			for j := 0; j < 3; j++ {
				assert.Equal(t, oneMapping|twoMapping|threeMapping, b.individualSets[i][j])
			}
			for j := 3; j < 9; j++ {
				assert.Equal(t, emptyMapping, b.individualSets[i][j])
			}
		} else {
			assert.Equal(t, oneMapping, b.individualSets[i][0])
			assert.Equal(t, twoMapping, b.individualSets[i][1])
			assert.Equal(t, threeMapping, b.individualSets[i][2])
			for j := 3; j < 9; j++ {
				assert.Equal(t, emptyMapping, b.individualSets[i][j])
			}
		}
	}
}

func TestNextMoves(t *testing.T) {
	b := NewGameWithState(Board{{
		1, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		4, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		7, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		2, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		5, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		8, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		3, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		6, 0, 0, 0, 0, 0, 0, 0, 0,
	}, {
		9, 0, 0, 0, 0, 0, 0, 0, 0,
	}})
	assert.False(t, b.Complete())

	rowGuess, colGuess, valuesGuess := b.NextGuessMove()
	assert.Equal(t, 0, rowGuess)
	assert.Equal(t, 2, colGuess)
	assert.Len(t, valuesGuess, 6)
	for _, v := range []int{2, 3, 5, 6, 8, 9} {
		assert.Contains(t, valuesGuess, v)
	}

	oldGuess := valuesGuess[0]
	b, ok := b.Set(rowGuess, colGuess, valuesGuess[0])
	assert.True(t, ok)

	rowGuess, colGuess, valuesGuess = b.NextGuessMove()
	assert.Equal(t, 0, rowGuess)
	assert.Equal(t, 1, colGuess)
	assert.Len(t, valuesGuess, 5)
	assert.NotContains(t, valuesGuess, oldGuess)
}

func TestNextMoveUnsolvable(t *testing.T) {
	b := NewGameWithState(Board{{
		1, 2, 4, 5, 6, 8, 7, 9, 3,
	}, {
		3, 5, 6, 0, 7, 9, 1, 2, 4,
	}, {
		7, 8, 9, 3, 1, 2, 5, 6, 0,
	}})

	rowGuess, colGuess, valuesGuess := b.NextGuessMove()
	assert.Equal(t, 0, rowGuess)
	assert.Equal(t, 0, colGuess)
	assert.Empty(t, valuesGuess)
}
