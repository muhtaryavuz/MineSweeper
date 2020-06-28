package main

import (
	"fmt"
	"minesweeper/display"
)

func main() {

	var runState int = 1
	var boardRowSize, boardColSize, rowNumber, colNumber, mineNumber, displayRet int
	state := 0
	for runState > 0 {

		switch state {
		case 0:
			display.Clear()
			fmt.Println("Hello	miners")
			fmt.Print("Enter board size 1...99 [row col] (e.g 10 10) > ")
			fmt.Scanf("%d %d", &boardRowSize, &boardColSize)
			fmt.Printf("%d %d", boardRowSize, boardColSize)
			if boardRowSize > 0 && boardRowSize < 100 && boardColSize > 0 && boardColSize < 100 {
				state = 1
			} else {
				state = 0
			}

		case 1:
			display.Clear()
			fmt.Printf("Enter mine size 1...%d > ", boardColSize*boardColSize*8/10)
			fmt.Scanf("%d", &mineNumber)

			if mineNumber > 0 && mineNumber < boardColSize*boardColSize*8/10 {
				display.Init(boardRowSize, boardColSize, mineNumber)
				state = 2
			} else {
				state = 1
			}

		case 2:
			display.Clear()
			display.Display(0, 0)
			fmt.Print("Enter cell coordinate [row col] (e.g 3 2) > ")
			fmt.Scanf("%d %d", &rowNumber, &colNumber)
			displayRet = display.Display(rowNumber, colNumber)
			if displayRet == -1 || displayRet == 0 {
				state = 2
			} else if displayRet == 1 {
				state = 3
			} else {
				state = 4
			}

		case 3:

			fmt.Println("*** GAME OVER ***")
			fmt.Println("Enter 9 for New Game or any key to exit")
			fmt.Scanf("%d", &runState)
			if runState == 9 {
				state = 0
			} else {
				runState = 0
			}

		case 4:
			fmt.Println("*** CONGRATULATIONS!!! YOU WON ***")
			fmt.Println("Enter 9 for New Game or any key to exit")
			fmt.Scanf("%d", &runState)
			if runState == 9 {
				state = 0
			} else {
				runState = 0
			}
		}
	}
}
