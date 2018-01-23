# Common files

We have two main common files, the smart puzzle and the naive puzzle. These mainly just differ on what data they keep.

## SmartPuzzle

### What functions we provide for a SmartPuzzle
 - GetSmartPuzzle(sudoku_string, puzzle_solver)
   - sudoku_string is a string form of a sudoku
   - puzzle_sovler is the function to call to solve the SmartPuzzle
   - returns a SmartPuzzle struct
 - GetNumPlacements
   - the number of placements is reset on every `GetSmartPuzzle` and `Solve` call
   - returns the number of placement attempts
 - GetSimple
   - returns a string representation of a sudoku
 - PrintSimple
   - prints the simple string to stdout
   - returns void
 - PrintPretty
   - prints a pretty version of the sudoku to stdout
   - returns void
 - Solve
   - resets number of placement attempts, and calls the puzzle solver function on the SmartPuzzle
   - returns a solution and nil, if there is one, else returns nil and the error
 - Place
   - increments number of placement attempts, and checks if the entry is valid.
   - if so, updates all affected data fields, and returns true, nil
   - if not, returns false and the error
 - Clone
   - Creates a new SmartPuzzle exactly like the last one and returns a pointer to it

### What we store on each SmartPuzzle
 - the total number of entries placed in the puzzle.
   - Stored as an `int`
 - a function called to solve the puzzle
 - the `Entry` in each `Tile`.
   - An `Entry` is an `int`, since I don't care about input.
   - A `Tile` is a `uint8`, since we only need numbers 0 through 9 (where 0 is an empty `Tile`).
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

## NaivePuzzle
 Currently, this is still in naive/, but I'm working on bringing it in here.