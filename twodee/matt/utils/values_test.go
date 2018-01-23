package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValueMap_getValues(t *testing.T) {
	assert.Equal(t, []int{1}, oneMapping.getValues())
	assert.Equal(t, []int{2}, twoMapping.getValues())
	assert.Equal(t, []int{3}, threeMapping.getValues())
	assert.Equal(t, []int{4}, fourMapping.getValues())
	assert.Equal(t, []int{5}, fiveMapping.getValues())
	assert.Equal(t, []int{6}, sixMapping.getValues())
	assert.Equal(t, []int{7}, sevenMapping.getValues())
	assert.Equal(t, []int{8}, eightMapping.getValues())
	assert.Equal(t, []int{9}, nineMapping.getValues())
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, fullMapping.getValues())
}

func TestValueMap_setValue(t *testing.T) {
	assert.True(t, oneMapping.valueIsSet(1))
	assert.True(t, twoMapping.valueIsSet(2))
	assert.True(t, threeMapping.valueIsSet(3))
	assert.True(t, fourMapping.valueIsSet(4))
	assert.True(t, fiveMapping.valueIsSet(5))
	assert.True(t, sixMapping.valueIsSet(6))
	assert.True(t, sevenMapping.valueIsSet(7))
	assert.True(t, eightMapping.valueIsSet(8))
	assert.True(t, nineMapping.valueIsSet(9))

	assert.False(t, oneMapping.valueIsSet(2))
	assert.False(t, twoMapping.valueIsSet(3))
	assert.False(t, threeMapping.valueIsSet(4))
	assert.False(t, fourMapping.valueIsSet(5))
	assert.False(t, fiveMapping.valueIsSet(6))
	assert.False(t, sixMapping.valueIsSet(7))
	assert.False(t, sevenMapping.valueIsSet(8))
	assert.False(t, eightMapping.valueIsSet(9))
	assert.False(t, nineMapping.valueIsSet(10))
	assert.False(t, nineMapping.valueIsSet(0))

	assert.Equal(t, oneMapping, oneMapping.setValue(1))
	assert.Equal(t, oneMapping|threeMapping, oneMapping.setValue(3))
	assert.Equal(t, 2, oneMapping.setValue(3).numberSet())
}
