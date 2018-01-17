package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"time"

	"github.com/joshprzybyszewski/sudoku_fun/matt/sudoku/2d/utils"
)

var (
	filePath = flag.String(`input`, `files/input.sk`, `path of file containing puzzles`)
)

func peaceOut(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
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
