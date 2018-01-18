package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BruteForce(t *testing.T) {
	solved := BruteForceCheck(`348125769965347821712869543621784935573291684894536172156472398239658417487913256`)
	assert.True(t, solved)
}
