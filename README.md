# sudoku_fun
[![codecov](https://codecov.io/gh/joshprzybyszewski/sudoku_fun/branch/master/graph/badge.svg)](https://codecov.io/gh/joshprzybyszewski/sudoku_fun)

A sudoku solver using [Go](https://golang.org/).

## How It Works
Simple! Forget about storage space and care only about speed!

### What we store on each `Puzzle`:
 - the `Entry` in each `Tile`.
   - An `Entry` is an `int`, since I don't care about input.
   - A `Tile` is a `uint8`, since we only need numbers 0 through 9 (where 0 is an empty `Tile`).
 - the total number of entries placed in the puzzle.
   - Stored as an `int`
 - the entries present in each row.
   - Represented as a `Presence`, which is `uint16`
   - The bits on `Presence` represent which entries are present in that row, where `1 << 0` means 1 is present, `1 << 1` means 2 is present, etc.
   - We make it a `uint16` because we need at least 9 bits to represent numbers 1 through 9 and use 0 as "No entrys present in this row".
 - the entries present in each column.
 - the entries present in each box.
 - the number of open spots in each row.
   - Represented as a `uint8`
   - Initialized to 9 on creation of the `Puzzle`
   - We cache this so that we don't have to calculate it from the `Presence` for the row every time we need it.
 - the number of open spots in each col.
 - the number of open spots in each box.
 
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

## How To Run
 1. Make sure this is in your machine's Go path.
 2. Running `go run main_runner.go` will solve all the puzzles in the example file.
 
## Contributing
 This is primarily used for personal purposes. Please feel free to look through the code and use it for yourself, but don't expect me to respond to issues or pull requests.
 
## Why Did I Make This?
 A coworker and I want to try solving three dimensional Sudokus, if that's possible. So we're in a little friendly competition about who can write the better two dimensional solver. His solves in 120ms, but mine's at 9ms on average.
