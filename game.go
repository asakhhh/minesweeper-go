package main

import "fmt"

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
		PrintStatistics()
	} else if (*revealed)[y][x] { // INVALID MOVE
		PrintString("Invalid coordinates: already revealed cell.\n")
		EnterMove(matrix, revealed)
	} else { // NORMAL MOVE
		if (*matrix)[y][x] == 0 { // EMPTY CELL
			Reveal(matrix, revealed, x, y)
		} else {
			(*revealed)[y][x] = true
			CLOSED_COUNT--
		}

		if CLOSED_COUNT == 0 { // WIN
			PrintMap()
			PrintString("You Win!\n")
			PrintStatistics()
		} else {
			EnterMove(matrix, revealed)
		}
	}
}
