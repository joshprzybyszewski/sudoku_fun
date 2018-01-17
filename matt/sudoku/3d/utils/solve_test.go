package utils

import (
	"log"
	"testing"
)

func TestSolve(t *testing.T) {
	c, i := Solve()
	log.Println(c, i)
}
