package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"time"

	twodee "github.com/joshprzybyszewski/sudoku_fun/twodee_bitwise"
	"github.com/joshprzybyszewski/sudoku_fun/twodee_bitwise/naive"
	"github.com/joshprzybyszewski/sudoku_fun/twodee_bitwise/robust"
	"github.com/joshprzybyszewski/sudoku_fun/twodee_bitwise/smart"
	matt "github.com/joshprzybyszewski/sudoku_fun/matt/sudoku/2d/solver"
	"github.com/joshprzybyszewski/sudoku_fun/utils"
	"github.com/joshprzybyszewski/sudoku_fun/utils/types"
)

var (
	filePath = flag.String(`sudokuInput`, `example-puzzles/puzzles.sk`, `path of file containing puzzles`)

	naivePerfomance  = &algoPerformance{naiveRead, map[int64]puzzleInfo{}, map[int]puzzleInfo{}, map[int]puzzleInfo{}}
	smartPerfomance  = &algoPerformance{smartRead, map[int64]puzzleInfo{}, map[int]puzzleInfo{}, map[int]puzzleInfo{}}
	robustPerfomance = &algoPerformance{robustRead, map[int64]puzzleInfo{}, map[int]puzzleInfo{}, map[int]puzzleInfo{}}
	mattsPerfomance  = &algoPerformance{mattRead, map[int64]puzzleInfo{}, map[int]puzzleInfo{}, map[int]puzzleInfo{}}
)

func naiveRead(entries string) (s types.Sudoku, err error) {
	pzl, err := naive.ReadSudoku(entries)
	if err != nil {
		return nil, err
	}

	return types.Sudoku(pzl), nil
}
func smartRead(entries string) (s types.Sudoku, err error) {
	pzl, err := smart.ReadSudoku(entries)
	if err != nil {
		return nil, err
	}

	return types.Sudoku(pzl), nil
}
func robustRead(entries string) (s types.Sudoku, err error) {
	pzl, err := robust.ReadSudoku(entries)
	if err != nil {
		return nil, err
	}

	return types.Sudoku(pzl), nil
}
func mattRead(entries string) (s types.Sudoku, err error) {
	pzl, err := matt.ReadSudoku(entries)
	if err != nil {
		return nil, err
	}

	return types.Sudoku(pzl), nil
}

type algoPerformance struct {
	readPuzzle      twodee.SudokuReader
	byTime          map[int64]puzzleInfo
	byNumPlacements map[int]puzzleInfo
	byPuzzleNumber  map[int]puzzleInfo
}

type puzzleInfo struct {
	duration      *time.Duration
	numPlacements int
	solutionStr   string
	puzzleNumber  int
}

type puzzleSolver func(readPuzzle twodee.SudokuReader, singleLinePzl string) (executionTime *time.Duration, numPlacements int, solutionString string)

func aaaaaahhhhhhh(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	runTestForAllPuzzles(naivePerfomance, twodee.PuzzleSolver)
	println(`finished naive!`)
	runTestForAllPuzzles(smartPerfomance, twodee.PuzzleSolver)
	println(`finished smart!`)
	runTestForAllPuzzles(robustPerfomance, twodee.PuzzleSolver)
	println(`finished robust!`)
	runTestForAllPuzzles(mattsPerfomance, matt.PuzzleSolver)
	println(`finished matts!`)

	println(`NAIVE STATS`)
	naivePerfomance.printPerformanceStats()
	println(`SMART STATS`)
	smartPerfomance.printPerformanceStats()
	println(`ROBUST STATS`)
	robustPerfomance.printPerformanceStats()
	println(`MATT STATS`)
	mattsPerfomance.printPerformanceStats()
}

func runTestForAllPuzzles(ap *algoPerformance, slvr puzzleSolver) {
	flag.Parse()

	fp, err := os.Open(*filePath)
	aaaaaahhhhhhh(err)

	bfp := bufio.NewReader(fp)

	for i := 1; ; i++ {
		line, _, err := bfp.ReadLine()
		if err == io.EOF {
			break
		}
		aaaaaahhhhhhh(err)

		dur, num, str := slvr(ap.readPuzzle, string(line))
		pi := puzzleInfo{dur, num, str, i}
		ap.byTime[dur.Nanoseconds()] = pi
		ap.byNumPlacements[num] = pi
		ap.byPuzzleNumber[i] = pi
	}
}

func (ap *algoPerformance) printPerformanceStats() {
	//ap.printEveryPuzzle()
	ap.printAverages()
	ap.printWorstByTime()
	ap.printWorstByTrials()
}
func (ap *algoPerformance) printEveryPuzzle() {
	for i, pzl := range ap.byPuzzleNumber {
		actuallySolved := utils.BruteForceCheck(pzl.solutionStr)
		println(fmt.Sprintf("Solved (%v) Puzzle #%v in \t%9.4fms with \t%6v tries", actuallySolved, i, float64(pzl.duration.Nanoseconds())/1000000.0, pzl.numPlacements))
	}

}
func (ap *algoPerformance) printAverages() {
	total := 0
	totalDur := int64(0)
	totalNumPlaces := 0
	for _, pi := range ap.byPuzzleNumber {
		total++
		totalDur += pi.duration.Nanoseconds()
		totalNumPlaces += pi.numPlacements
	}
	averageDurationNs := totalDur / int64(total)
	averageNumPlacesNs := totalNumPlaces / total

	println(`======AVERAGES======`)
	println(fmt.Sprintf("Total # tests:        %v", total))
	println(fmt.Sprintf("Average duration:  %9.4fms", float64(averageDurationNs)/1000000.0))
	println(fmt.Sprintf("Average # Placements: %v", averageNumPlacesNs))
	println(`====================`)
}
func (ap *algoPerformance) printWorstByTime() {
	println(`==WORST 10 (by duration of solve)==`)

	numWorst := 10
	worstTimes := make([]int64, 0, numWorst+1)
	for testTime := range ap.byTime {
		shouldAdd := false
		minI := 0
		var minVal int64
		if len(worstTimes) > 0 {
			minVal = worstTimes[0]
		} else {
			shouldAdd = true
		}

		for worstIndex, key := range worstTimes {
			if testTime > key {
				shouldAdd = true
			}
			if key < minVal {
				minI = worstIndex
				minVal = key
			}
		}

		if shouldAdd {
			if len(worstTimes) >= numWorst {
				worstTimes[minI] = testTime
			} else {
				worstTimes = append(worstTimes, testTime)
			}
		}
	}

	i := 0
	for _, key := range worstTimes {
		if i > numWorst {
			break
		}

		pi := ap.byTime[key]

		println(fmt.Sprintf("Puzzle #%3v took %9.4fms", pi.puzzleNumber, float64(pi.duration.Nanoseconds())/1000000.0))
		i++
	}
}
func (ap *algoPerformance) printWorstByTrials() {
	println(`==WORST 10 (by number of placements)==`)
	i := 0
	numWorst := 10
	placements := make([]int, 0, len(ap.byNumPlacements))
	for p := range ap.byNumPlacements {
		placements = append(placements, p)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(placements)))
	for _, key := range placements {
		if i > numWorst {
			break
		}

		pi := ap.byNumPlacements[key]

		println(fmt.Sprintf("Puzzle #%3v took %6v placements", pi.puzzleNumber, pi.numPlacements))
		i++
	}
}
