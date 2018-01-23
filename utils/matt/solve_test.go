package matt

import (
	"log"
	"testing"
)

func TestSolve(t *testing.T) {
	b := Board{}
	b, _ = Solve(b)
	log.Println(NewGameWithState(b).Print())
}
