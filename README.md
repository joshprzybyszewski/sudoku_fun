# sudoku_fun
[![codecov](https://codecov.io/gh/joshprzybyszewski/sudoku_fun/branch/master/graph/badge.svg)](https://codecov.io/gh/joshprzybyszewski/sudoku_fun)

A sudoku solver using [Go](https://golang.org/).

## What we're doing
We've established different algorithms for solving sudokus, both 2D and 3D.
At the moment, we only have functioning 2d solutions, but we're thinking about how to solve 3d and what the standards of completion will be.

### Josh's 2D solutions
[Naive](./twodee_bitwise/naive/README.md)
  _This just chooses the next location by finding the next open spot. Very stoopid._
[Smart](./twodee_bitwise/smart/README.md)
  _Scans the puzzle for the a good open spot to attempt next placement._
["robust" version of Smart](./twodee_bitwise/robust/README.md)
  _Completes every error check along the way. This solution checks every err != nil and adds ~25% of execution time._
[Very Smart](./twodee_bitwise/verysmart/README.md)
  _Checks every Tile to find the best location to attempt placement and uses a cache to find the possible entries._

### Matt's 2D Solution
[Sometime there might be a readme](./matt/sudoku/2d/solver/solver.go)
  _He should edit this README so that he can explain himself a little better._

## How To Run
 1. Make sure this is in your machine's Go path.
   - I think you could do this with `go get github.com/joshprzybyszewski/sudoku_fun`
   - Then you might need to do a `go install` or a `godep save`. Honestly, I have no idea right now, and I need a friend to try it out to be able to document exactly what you'll need.
 2. Running `go run main_runner.go` will solve all the puzzles in the example file.
   - The purpose of this is to pit the algorithms against each other. Doing anything else rn will take some tlc.
 
## Contributing
 This is primarily used for personal purposes. Please feel free to look through the code and use it for yourself, but don't expect me to respond to issues or pull requests.

## Why Did I Make This?
 A coworker and I want to try solving three dimensional Sudokus, if that's possible. So we're in a little friendly competition about who can write the better two dimensional solver.
