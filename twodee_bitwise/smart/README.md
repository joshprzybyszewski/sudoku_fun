# Josh's "smart" algorithm

## How It Works
Simple! Forget about storage space and care only about speed!
Also, I broke out all the common logic of this and "very smart" to the [common directory](../common/README.md) using the SmartPuzzle

### What data structure we use
 This uses a `SmartPuzzle`. Go read about them [here](../common/README.md).

### Solving Algorithm
 Base Case: The puzzle has the number of entries equal to the number of tiles
  - Assume the puzzle is solved and return the `Puzzle`

 Else: Try solving the Puzzle!
  - Find the best (row, col) location `[br][bc]` for attempting placement
    - Scan the rows and cols and find the row and the col with the least number of entries
    - If the best row `br` has fewer or the same free spots as the best col,
      - Find the column `c` which has an open tile in `[br][c]` and the fewest open tiles in column `c`
    - Else the best column `bc` has fewer free spots than the best row,
      - Find the row `r` which has an open tile in `[r][bc]` and the fewest open tiles in row `r`
  - Get the Entrys, `entries`, which are available at the best location
    - Basically use the bitstrings of the present Entrys in the row, col, box to identify what can be placed at `[br][bc]`
  - foreach `Entry e` in `entries`
    - Copy the current puzzle into `pClone`
    - In `pClone`, place `e` at `[br][bc]` and update all the other puzzle state vars
    - Call `solve` on `pClone`
    - If that returns a valid solution, pass that on up
    - Else, continue with the foreach loop because we are guaranteed one of the possible entries goes in this location