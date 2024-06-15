package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func EnterMove(matrix *[][]int, revealed *[][]bool) {
	PrintMap(&matrix, &revealed)

	PrintString("Enter coordinates: ")

	var x, y int
	_, err := fmt.Scanf("%d %d", &x, &y)
	if err != nil || !(x > 0 && x <= WIDTH) || !(y > 0 && y <= HEIGHT) {
		PrintString("Invalid coordinates. Please enter in format <x y>\n")
		EnterMove(matrix, revealed)
	}

	if (*matrix)[y][x] == -1 { // BOMB
		PrintMap()
		PrintString("Game Over!\n")
		MOVE_COUNT++
		PrintStatistics()
	} else if (*revealed)[y][x] { // INVALID MOVE
		PrintString("Invalid coordinates: already revealed cell.\n")
		EnterMove(matrix, revealed)
	} else { // NORMAL MOVE
		Reveal(matrix, revealed, x, y)
		MOVE_COUNT++

		if CLOSED_COUNT == 0 { // WIN
			PrintMap()
			PrintString("You Win!\n")
			PrintStatistics()
		} else {
			EnterMove(matrix, revealed)
		}
	}
}

func PrintStatistics() {
	PrintString("Your statistics:\n")

	PrintString("- Field size: ")
	PrintNum(HEIGHT)
	ap.PutRune('x')
	PrintNum(WIDTH)
	ap.PutRune('\n')

	PrintString("- Number of bombs: ")
	PrintNum(BOMB_COUNT)
	ap.PutRune('\n')

	PrintString("- Number of moves: ")
	PrintNum(MOVE_COUNT)
	ap.PutRune('\n')
}

func Reveal(matrix *[][]int, revealed *[][]bool, x, y int) {
	(*revealed)[y][x] = true
	CLOSED_COUNT--

	if (*matrix)[y][x] > 0 {
		return
	}

	for y_next := y - 1; y_next <= y+1; y_next++ {
		for x_next := x - 1; x_next <= x+1; x_next++ {
			if y_next == 0 || y_next == HEIGHT+1 || x_next == 0 || x_next == WIDTH+1 {
				continue
			}
			if !(*revealed)[y_next][x_next] {
				Reveal(matrix, revealed, x_next, y_next)
			}
		}
	}
}
