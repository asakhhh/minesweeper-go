package main

import (
	"fmt"
)

func CustomMap(matrix *[][]int) int {
	NewMatrix := make([][]int, HEIGHT+2)
	for i := range NewMatrix {
		NewMatrix[i] = make([]int, WIDTH+2)
	}

	PrintString("Enter ")
	PrintNum(HEIGHT)
	PrintString(" lines, ")
	PrintString("each with ")
	PrintNum(WIDTH)
	PrintString(" characters representing the grid: '.' represents an empty cell, '*' represents a bomb (e.g., '..*' for width=3):\n")

	bomb_count := 0

	for i := 1; i <= HEIGHT; i++ {
		var row string
		_, err := fmt.Scanf("%s", &row)
		if err != nil || len(row) != WIDTH {
			print("Invalid format for row. Please enter exactly ")
			PrintNum(WIDTH)
			print(" characters.\n")
			return -1
		}
		for j := 1; j <= WIDTH; j++ {
			if row[j-1] != '.' && row[j-1] != '*' {
				print("Invalid character in row. Please enter only '.' or '*'.\n")
				return -1
			}
			if row[j-1] == '*' {
				NewMatrix[i][j] = -1
				bomb_count++
			} else if row[j-1] == '.' {
				NewMatrix[i][j] = 0
			}
		}
	}
	*matrix = append(*matrix, NewMatrix...)

	return bomb_count
}
