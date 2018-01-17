package utils

import (
	"fmt"
	"sync/atomic"
)

func Solve() (*Cube, int64) {
	attempts := int64(0)
	solveList, _ := solve(NewGame(), 0, &attempts)
	if len(solveList) > 0 && solveList[0] != nil {
		return solveList[0], attempts
	}
	return nil, attempts
}

func solve(g *Cube, depth int64, attempts *int64) ([]*Cube, bool) {
	if g.Complete() {
		return []*Cube{g}, true
	}

	rowGuess, colGuess, zGuess, guesses := g.NextGuessMove()
	for _, guess := range guesses {
		if g2, ok := g.Set(rowGuess, colGuess, zGuess, guess); ok {
			//fmt.Println("START:", g2.Print(), "END")

			//time.Sleep(time.Second)

			fmt.Printf("%v: %v. [%d, %d, %d]@%d\n", depth, atomic.AddInt64(attempts, 1), rowGuess, colGuess, zGuess, guess)
			if solveList, solved := solve(g2, depth+1, attempts); solved {
				return append(solveList, g), solved
			}
		}
	}
	return nil, false
}
