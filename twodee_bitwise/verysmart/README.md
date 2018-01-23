# Josh's "very smart" algorithm

## How It Works
I made this as a branch from the Smart algorithm. So go read [that README](../smart/README.md), and I'll explain the differences here

### Solving Algorithm Differences

#### Find the best (row, col) location
Previously, we scanned every row and every col to find the one with the fewest possible entries.
Now, we scan every Tile to find the single location with the fewest possible entries.

How do we do that efficiently?
 1. Initialize a table of all 512 Presence options (remember, Presence is just an int, so this is literally just an array).
    - The table contains the number of possible entries for that presence, and a list of what those entries are.
    - We do this once before we start solving puzzles.
 2. Pass in the OR of the presences of each Tile to that cached table: (row|col|box)
 3. If we encounter a Tile with only one possible entry, choose that Tile immediately and break out of the r X c loop.
