package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"time"

	twodee "./twodee_bitwise"
	"./utils"
)

var (
	filePath = flag.String(`sudokuInput`, `example-puzzles/puzzles.sk`, `path of file containing puzzles`)

	byTime          = map[int64]puzzleInfo{}
	byNumPlacements = map[int]puzzleInfo{}
	byPuzzleNumber  = map[int]puzzleInfo{}
)

type puzzleInfo struct  {
	duration *time.Duration
	numPlacements int
	solutionStr string
	puzzleNumber int
}

type puzzleSolver func (singleLinePzl string) (executionTime *time.Duration, numPlacements int, solutionString string)

func aaaaaahhhhhhh(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	runTestForAllPuzzles(twodee.PuzzleSolver)
}

func runTestForAllPuzzles(slvr puzzleSolver) {
	flag.Parse()

	fp, err := os.Open(*filePath)
	aaaaaahhhhhhh(err)

	bfp := bufio.NewReader(fp)

	for i := 1;; i++ {
		line, _, err := bfp.ReadLine()
		if err == io.EOF || i > 2 {
			break
		}
		aaaaaahhhhhhh(err)
		println(string(line))

		dur, num, str := slvr(string(line))
		pi := puzzleInfo{dur, num, str, i}
		byTime[dur.Nanoseconds()] = pi
		byNumPlacements[num] = pi
		byPuzzleNumber[i] = pi
	}

	printEveryPuzzle()
	printAverages()
	printWorstByTime()
	printWorstByTrials()
}

func printEveryPuzzle() {
	for i, pzl := range byPuzzleNumber {
		actuallySolved := utils.BruteForceCheck(pzl.solutionStr)
		println(fmt.Sprintf("Solved (%v) Puzzle #%v in \t%9.4fms with \t%6v tries", actuallySolved, i, float64(pzl.duration.Nanoseconds()) / 1000000.0, pzl.numPlacements))
		println(pzl.solutionStr)
	}

}
func printAverages() {
	total := 0
	totalDur := int64(0)
	totalNumPlaces := 0
	for _, pi := range byPuzzleNumber {
		total++
		totalDur += pi.duration.Nanoseconds()
		totalNumPlaces += pi.numPlacements
	}
	averageDurationNs := totalDur / int64(total)
	averageNumPlacesNs := totalNumPlaces / total

	println(`======AVERAGES======`)
	println(fmt.Sprintf("Total # tests:        %v", total))
	println(fmt.Sprintf("Average duration:  %9.4fms", float64(averageDurationNs) / 1000000.0))
	println(fmt.Sprintf("Average # Placements: %v", averageNumPlacesNs))
	println(`====================`)
}
func printWorstByTime() {
	println(`==WORST 10 (by duration of solve)==`)

	numWorst := 10
	worstTimes := make([]int64, 0, numWorst + 1)
	for testTime := range byTime {
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

		pi := byTime[key]

		println(fmt.Sprintf("Puzzle #%3v took %9.4fms", pi.puzzleNumber, float64(pi.duration.Nanoseconds()) / 1000000.0))
		i++
	}
}
func printWorstByTrials() {
	println(`==WORST 10 (by number of placements)==`)
	i := 0
	numWorst := 10
	placements := make([]int, 0, len(byNumPlacements))
	for p := range byNumPlacements {
		placements = append(placements, p)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(placements)))
	for _, key := range placements {
		if i > numWorst {
			break
		}

		pi := byNumPlacements[key]

		println(fmt.Sprintf("Puzzle #%3v took %6v placements", pi.puzzleNumber, pi.numPlacements))
		i++
	}
}