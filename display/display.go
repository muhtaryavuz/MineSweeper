package display

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const numOfNeighbors = 8
const haveMine = 100

type intCell interface {
	init()
}

type cell struct {
	isVisible       bool
	row, col, value int
	neighbors       [numOfNeighbors]*cell
}

func (t *cell) init(r, c, v int) {
	t.row = r
	t.col = c
	t.value = v
	t.isVisible = false
	for i := 0; i < numOfNeighbors; i++ {
		t.neighbors[i] = nil
	}
}

var gameBoard [][]cell
var numOfRow, numOfColumn, numOfMines int

func inRange(row, col int) bool {
	if row > 0 && row <= numOfRow && col > 0 && col <= numOfColumn {
		return true
	}
	return false
}

func minePlacer() {
	var randNumber, index int = -1, 0
	var row, col int

	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)

	for index < numOfMines {
		randNumber = r2.Intn(numOfRow * numOfColumn)
		if randNumber >= 0 && randNumber < numOfRow*numOfColumn {
			row = randNumber / numOfColumn
			col = randNumber % numOfColumn
			if gameBoard[row][col].value != haveMine {
				gameBoard[row][col].value = haveMine
				index++
			}
		}
	}
}

func calculateValues() {

	for i := 0; i < numOfRow; i++ {
		for j := 0; j < numOfColumn; j++ {
			if gameBoard[i][j].value != haveMine {
				for k := 0; k < numOfNeighbors; k++ {
					if gameBoard[i][j].neighbors[k] != nil && gameBoard[i][j].neighbors[k].value == haveMine {
						gameBoard[i][j].value++
					}
				}
			}
		}
	}
}

func openZeroValues(box *cell) {

	if box.value == 0 {
		box.isVisible = true

		for i := 0; i < numOfNeighbors; i++ {
			if box.neighbors[i] != nil {
				if box.neighbors[i].value == 0 && box.neighbors[i].isVisible == false {
					openZeroValues(box.neighbors[i])
				} else {
					box.neighbors[i].isVisible = true
				}
			}
		}
	}
}

//Clear clean the terminal
func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

//Init inits game board
func Init(row, column, mineNumber int) {
	numOfRow = row
	numOfColumn = column
	numOfMines = mineNumber

	gameBoard = make([][]cell, numOfRow)

	for i := range gameBoard {
		gameBoard[i] = make([]cell, numOfColumn)
	}
	for i := 0; i < numOfRow; i++ {
		for j := 0; j < numOfColumn; j++ {
			gameBoard[i][j].init(i, j, 0)
		}
	}

	for i := 0; i < numOfRow; i++ {
		for j := 0; j < numOfColumn; j++ {
			if i-1 >= 0 {
				gameBoard[i][j].neighbors[0] = &gameBoard[i-1][j]

				if j+1 <= numOfColumn-1 {
					gameBoard[i][j].neighbors[1] = &gameBoard[i-1][j+1]
				}

				if j-1 >= 0 {
					gameBoard[i][j].neighbors[7] = &gameBoard[i-1][j-1]
				}
			}
			if i+1 <= numOfRow-1 {
				gameBoard[i][j].neighbors[4] = &gameBoard[i+1][j]

				if j+1 <= numOfColumn-1 {
					gameBoard[i][j].neighbors[3] = &gameBoard[i+1][j+1]
				}

				if j-1 >= 0 {
					gameBoard[i][j].neighbors[5] = &gameBoard[i+1][j-1]
				}
			}

			if j+1 <= numOfColumn-1 {
				gameBoard[i][j].neighbors[2] = &gameBoard[i][j+1]
			}

			if j-1 >= 0 {
				gameBoard[i][j].neighbors[6] = &gameBoard[i][j-1]
			}
		}
	}
	minePlacer()
	calculateValues()
}

// Display displays minefield with current configuration
func Display(row, col int) int {

	var ret, openCellNumber int = 0, 0

	if inRange(row, col) {
		if gameBoard[row-1][col-1].value == haveMine {
			ret = 1
			for i := 0; i < numOfRow; i++ {
				for j := 0; j < numOfColumn; j++ {
					gameBoard[i][j].isVisible = true
				}
			}
		} else {
			gameBoard[row-1][col-1].isVisible = true
			openZeroValues(&gameBoard[row-1][col-1])
		}

	} else {
		if row != 0 || col != 0 {
			ret = -1
		}
	}

	if ret == 1 {
		Clear()
	}

	for i := 0; i < numOfRow; i++ {
		if i == 0 {
			fmt.Print("   ")
			for a := 0; a < numOfColumn; a++ {
				fmt.Printf(" %2d ", a+1)
			}
			fmt.Println()
		}
		fmt.Print("   ")
		for a := 0; a < numOfColumn; a++ {
			fmt.Print("....")
		}
		fmt.Println(".")

		for j := 0; j < numOfColumn; j++ {
			if j == 0 {
				fmt.Printf("%2d ", i+1)
			}

			fmt.Print(":")

			if gameBoard[i][j].isVisible {
				openCellNumber++
				if gameBoard[i][j].value == haveMine {
					fmt.Print(" * ")
				} else {
					fmt.Printf(" %d ", gameBoard[i][j].value)
				}
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println(":")
	}
	fmt.Print("   ")
	for a := 0; a < numOfColumn; a++ {
		fmt.Print("....")
	}
	fmt.Println(".")

	if ret != 1 && openCellNumber >= numOfRow*numOfColumn-numOfMines {
		ret = 2
	}

	return ret
}
