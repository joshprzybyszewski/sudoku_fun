package utils

import (
	"log"
	"testing"
)

func TestSolve(t *testing.T) {
	t.Skip(`skipping test until we know it will solve`)
	c, i := Solve()
	log.Println(c, i)
}
