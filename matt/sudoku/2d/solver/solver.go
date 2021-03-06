package solver

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joshprzybyszewski/sudoku_fun/matt/sudoku/2d/utils"

	"github.com/joshprzybyszewski/sudoku_fun/twodee_bitwise"
	joshUtils "github.com/joshprzybyszewski/sudoku_fun/utils"
	"github.com/joshprzybyszewski/sudoku_fun/utils/constants"
)

var (
	filePath = flag.String(`input`, `files/input.sk`, `path of file containing puzzles`)
)

func peaceOut(err error) {
	if err != nil {
		panic(err)
	}
}

func whatusedtobemain() {
	flag.Parse()

	fp, err := os.Open(*filePath)
	peaceOut(err)

	bfp := bufio.NewReader(fp)
	for {
		line, _, err := bfp.ReadLine()
		if err == io.EOF {
			break
		}
		peaceOut(err)

		board := utils.Board{}
		for i, b := range []rune(string(line)) {
			switch b {
			case '1':
				board[i/9][i%9] = 1
			case '2':
				board[i/9][i%9] = 2
			case '3':
				board[i/9][i%9] = 3
			case '4':
				board[i/9][i%9] = 4
			case '5':
				board[i/9][i%9] = 5
			case '6':
				board[i/9][i%9] = 6
			case '7':
				board[i/9][i%9] = 7
			case '8':
				board[i/9][i%9] = 8
			case '9':
				board[i/9][i%9] = 9
			}
		}
		t0 := time.Now()
		_, attempts := utils.Solve(board)
		log.Printf("Solved in %v after %d attempts", time.Since(t0), attempts)
	}
}

func readPuzzle(line string) (*utils.Board, error) {
	board := utils.Board{}
	for i, b := range []rune(string(line)) {
		switch b {
		case '1':
			board[i/9][i%9] = 1
		case '2':
			board[i/9][i%9] = 2
		case '3':
			board[i/9][i%9] = 3
		case '4':
			board[i/9][i%9] = 4
		case '5':
			board[i/9][i%9] = 5
		case '6':
			board[i/9][i%9] = 6
		case '7':
			board[i/9][i%9] = 7
		case '8':
			board[i/9][i%9] = 8
		case '9':
			board[i/9][i%9] = 9
		case '.':
			board[i/9][i%9] = 0
		default:
			return nil, errors.New(`why did this happen`)
		}
	}

	return &board, nil
}

type BoardReader func(str string) (p *utils.Board, err error)

func PuzzleSolver(_ twodee_bitwise.SudokuReader, sudokuPuzzle string) (*time.Duration, int, string) {
	board, err := readPuzzle(sudokuPuzzle)
	if err != nil {
		println(`BAD SUDOKU!!`)
		return nil, 0, ``
	}

	t0 := time.Now()
	solvedBoard, attempts := utils.Solve(*board)
	duration := time.Since(t0)

	if err != nil {
		println(fmt.Sprintf("COULDN'T SOLVE: %v", err.Error()))
		return nil, -1, constants.EmptyPuzzle
	}

	solutionStr := toString(&solvedBoard)

	if !joshUtils.BruteForceCheck(solutionStr) {
		println(solutionStr)
		println(`WARNING IT DIDN'T ACTUALLY SOLVE'`)
		return nil, -1, constants.EmptyPuzzle
	}

	return &duration, int(attempts), solutionStr
}

func toString(board *utils.Board) string {
	str := ``
	for r := range board {
		for c := range board[r] {
			thechar := strconv.Itoa(board[r][c])
			str += thechar
		}

	}

	return str
}
