package main

import (
	"math/rand"

	"github.com/alem-platform/ap"
)

func CountAdjacentBombs(matrix *[][]int) {
	bomb_count := 0

	for i := 1; i <= HEIGHT; i++ {
		for j := 1; j <= WIDTH; j++ {
			if (*matrix)[i][j] != -1 {

				bomb_count = 0
				for x := i - 1; x <= i+1; x++ {
					for y := j - 1; y <= j+1; y++ {
						if (*matrix)[x][y] == -1 {
							bomb_count += 1
						}
					}
				}

				(*matrix)[i][j] = bomb_count
			}
		}
	}
}

func EnterMove(matrix *[][]int, revealed *[][]bool) {
	PrintMap(matrix, revealed)

	PrintString("Enter coordinates: ")

	x, y := ReadTwoNumbers()
	if x <= 0 || x > WIDTH || y <= 0 || y > HEIGHT {
		PrintString("Invalid coordinates. Please enter in format <x y>\n")
		EnterMove(matrix, revealed)
		return
	}

	if (*matrix)[y][x] == -1 { // BOMB
		if MOVE_COUNT == 0 && BOMB_COUNT < HEIGHT*WIDTH {
			new_cell := int(rand.Int31n(int32(HEIGHT * WIDTH)))
			cell_h := 1 + new_cell/WIDTH
			cell_w := 1 + new_cell%WIDTH

			for (*matrix)[cell_h][cell_w] == -1 {
				new_cell = int(rand.Int31n(int32(HEIGHT * WIDTH)))
				cell_h = 1 + new_cell/WIDTH
				cell_w = 1 + new_cell%WIDTH
			}

			(*matrix)[y][x] = 0
			(*matrix)[cell_h][cell_w] = -1
			CountAdjacentBombs(matrix)

			Reveal(matrix, revealed, x, y)
			MOVE_COUNT++

			if CLOSED_COUNT == 0 { // WIN
				PrintMap(matrix, revealed)
				PrintString("You Win!\n")
				PrintStatistics()
				return
			} else {
				EnterMove(matrix, revealed)
				return
			}
		}

		RevealBombs(matrix, revealed)
		PrintMap(matrix, revealed)
		PrintString("Game Over!\n")
		MOVE_COUNT++
		PrintStatistics()
		return
	} else if (*revealed)[y][x] { // INVALID MOVE
		PrintString("Invalid coordinates: already revealed cell.\n")
		EnterMove(matrix, revealed)
		return
	} else { // NORMAL MOVE
		Reveal(matrix, revealed, x, y)
		MOVE_COUNT++

		if CLOSED_COUNT == 0 { // WIN
			PrintMap(matrix, revealed)
			PrintString("You Win!\n")
			PrintStatistics()
			return
		} else {
			EnterMove(matrix, revealed)
			return
		}
	}
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

func RevealBombs(matrix *[][]int, revealed *[][]bool) {
	for i := 1; i <= HEIGHT; i++ {
		for j := 1; j <= WIDTH; j++ {
			if (*matrix)[i][j] == -1 {
				(*revealed)[i][j] = true
			}
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
