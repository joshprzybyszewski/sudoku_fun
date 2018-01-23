package utils

import "sync/atomic"

func Solve(b Board) (Board, int64) {
	attempts := int64(0)
	solveList, _ := solve(NewGameWithState(b), &attempts)
	if len(solveList) > 0 && solveList[0] != nil {
		return solveList[0].gameboard, attempts
	}
	return Board{}, attempts
}

func solve(g *Game, attempts *int64) ([]*Game, bool) {
	if g.Complete() {
		return []*Game{g}, true
	}

	rowGuess, colGuess, guesses := g.NextGuessMove()
	for _, guess := range guesses {
		if g2, ok := g.Set(rowGuess, colGuess, guess); ok {
			atomic.AddInt64(attempts, 1)
			if solveList, solved := solve(g2, attempts); solved {
				return append(solveList, g), solved
			}
		}
	}
	return []*Game{}, false
}
