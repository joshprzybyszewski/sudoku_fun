package utils

import (
	"sort"
	"strings"

	utils2D "github.com/Workiva/sudoku/2d/utils"
)

type Cube struct {
	zboards [9]*utils2D.Game
	yboards [9]*utils2D.Game
	xboards [9]*utils2D.Game
}

type nextGuess struct {
	row, col, z int
	guesses     []int
}

type nextGuessList []nextGuess

func (t nextGuessList) Len() int { return len(t) }
func (t nextGuessList) Less(i, j int) bool {
	iSet := len(t[i].guesses)
	jSet := len(t[j].guesses)
	if iSet > jSet {
		return true
	}
	if iSet == jSet {
		return t[i].z < t[j].z
	}
	return false
}
func (t nextGuessList) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func NewGame() *Cube {
	c := &Cube{
		zboards: [9]*utils2D.Game{},
		yboards: [9]*utils2D.Game{},
		xboards: [9]*utils2D.Game{},
	}
	for i, z := range c.zboards {
		if z == nil {
			c.zboards[i] = utils2D.NewGame()
		}
	}
	for i, y := range c.yboards {
		if y == nil {
			c.yboards[i] = utils2D.NewGame()
		}
	}
	for i, x := range c.xboards {
		if x == nil {
			c.xboards[i] = utils2D.NewGame()
		}
	}
	return c
}

func NewGameWithState(initialState [9]*utils2D.Game) *Cube {
	c := &Cube{
		zboards: [9]*utils2D.Game{},
		yboards: [9]*utils2D.Game{},
		xboards: [9]*utils2D.Game{},
	}
	for i, z := range initialState {
		if z == nil {
			c.zboards[i] = z.Copy()
		}
	}
	for i, y := range c.yboards {
		if y == nil {
			c.yboards[i] = utils2D.NewGame()
		}
	}
	for i, x := range c.xboards {
		if x == nil {
			c.xboards[i] = utils2D.NewGame()
		}
	}
	return c
}

func (c *Cube) clone() *Cube {
	c2 := NewGame()
	for i, z := range c.zboards {
		c2.zboards[i] = z.Copy()
	}
	for i, y := range c.yboards {
		c2.yboards[i] = y.Copy()
	}
	for i, x := range c.xboards {
		c2.xboards[i] = x.Copy()
	}
	return c2
}

// IsValid checks if the number is a valid number
func (c *Cube) ValidCube() bool {
	for _, z := range c.zboards {
		if !z.ValidGame() {
			return false
		}
	}
	for _, y := range c.yboards {
		if !y.ValidGame() {
			return false
		}
	}
	for _, x := range c.xboards {
		if !x.ValidGame() {
			return false
		}
	}
	return true
}

// IsValid checks if the number is a valid number
func (c *Cube) IsValid(row, col, z, value int) bool {
	if !c.zboards[z].IsValid(row, col, value) {
		return false
	}
	if !c.yboards[row].IsValid(z, col, value) {
		return false
	}
	return c.xboards[col].IsValid(row, z, value)
}

func (c *Cube) Set(row, col, z, value int) (*Cube, bool) {
	if !c.IsValid(row, col, z, value) {
		return nil, false
	}

	c2 := c.clone()
	z2, ok := c.zboards[z].Set(row, col, value)
	if !ok {
		return nil, false
	}

	y2, ok := c.yboards[row].Set(z, col, value)
	if !ok {
		return nil, false
	}

	x2, ok := c.xboards[col].Set(row, z, value)
	if !ok {
		return nil, false
	}
	c2.zboards[z] = z2
	c2.yboards[row] = y2
	c2.xboards[row] = x2
	return c2, true
}

func (c *Cube) Complete() bool {
	for _, z := range c.zboards {
		if !z.Complete() {
			return false
		}
	}
	return true
}

func (c *Cube) NextGuessMove() (row, col, z int, guesses []int) {
	ngl := nextGuessList{}
	for i, z := range c.zboards {
		row, col, guesses := z.NextGuessMove()
		if len(guesses) > 0 {
			ngl = append(ngl, nextGuess{
				row:     row,
				col:     col,
				z:       i,
				guesses: guesses,
			})
		} else if len(guesses) == 0 && !z.Complete() {
			return 0, 0, 0, nil
		}
	}
	for i, y := range c.yboards {
		row, col, guesses := y.NextGuessMove()
		if len(guesses) > 0 {
			ngl = append(ngl, nextGuess{
				row:     i,
				col:     col,
				z:       row,
				guesses: guesses,
			})
		} else if len(guesses) == 0 && !y.Complete() {
			return 0, 0, 0, nil
		}
	}
	for i, x := range c.xboards {
		row, col, guesses := x.NextGuessMove()
		if len(guesses) > 0 {
			ngl = append(ngl, nextGuess{
				row:     row,
				col:     i,
				z:       col,
				guesses: guesses,
			})
		} else if len(guesses) == 0 && !x.Complete() {
			return 0, 0, 0, nil
		}
	}

	sort.Sort(ngl)
	if len(ngl) == 0 {
		return 0, 0, 0, nil
	}

	return ngl[0].row, ngl[0].col, ngl[0].z, ngl[0].guesses
}

func (c *Cube) Print() string {
	str := []string{}
	for _, z := range c.zboards {
		str = append(str, z.Print())
	}
	return strings.Join(str, "\n")
}
