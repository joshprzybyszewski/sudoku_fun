package types

const (
	bestKnownLocationStart = -1
	sideLen = 9
)

type Entry int
type Tile uint8
type Presence uint16

type Sudoku interface {
	Solve() (s Sudoku, err error)
	GetNumPlacements() int
	GetSimple() string
	PrintSimple()
	PrintPretty()
}

type BestKnown struct {
	loc      int
	val      uint8
	lower    uint8
}

func (bk *BestKnown) WillUpdate(loc int, val uint8) bool {
	return val < bk.val && val > bk.lower
}

func (bk *BestKnown) Update(loc int, val uint8) {
	if bk.WillUpdate(loc, val) {
		bk.UpdateUnsafe(loc, val)
	}
}

func (bk *BestKnown) UpdateUnsafe(loc int, val uint8) {
	bk.loc = loc
	bk.val = val
}

func (bk *BestKnown) IsError() (bool) {
	return bk.loc == bestKnownLocationStart
}

func (bk *BestKnown) IsStrictlyBetterThan(other *BestKnown) (bool) {
	return bk.val < other.val
}

func (bk *BestKnown) IsBetterThan(other *BestKnown) (bool) {
	return bk.val <= other.val
}

func (bk *BestKnown) Location() int {
	return bk.loc
}

func (bk *BestKnown) Reset() {
	bk.loc = bestKnownLocationStart
	bk.val = uint8(sideLen + 1)
	bk.lower = 0
}

func NewBK() (*BestKnown) {
	bk := &BestKnown{}
	bk.Reset()
	return bk
}
