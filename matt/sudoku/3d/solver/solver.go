package main

import (
	"log"

	"github.com/joshprzybyszewski/sudoku_fun/matt/sudoku/3d/utils"
)

func main() {
	c, num := utils.Solve()
	log.Println(c.Print(), num)
}
