package main

import (
	"log"

	"github.com/Workiva/sudoku/3d/utils"
)

func main() {
	c, num := utils.Solve()
	log.Println(c.Print(), num)
}
