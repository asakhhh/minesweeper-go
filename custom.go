package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func CustomMap(h, w int, matrix [][]int) {
	matrix = make([][]int, h+2)
	for i := range matrix {
		matrix[i] = make([]int, w+2)
	}
	PrintString("Enter ")
	PrintNum(h)
	PrintString(" lines follow, ")
	PrintString("each with ")
	PrintNum(w)
	PrintString(" characters representing the grid: '.' represents an empty cell, '*' represents a bomb(e.g., '..*' for 3 width):\n")
	for i := 1; i <= h; i++ {
		var row string
		_, err := fmt.Scanf("%s", &row)
		if err != nil || len(row) != w {
			print("Invalid format for row. Please enter exactly ")
			PrintNum(w)
			print(" characters.\n")
			return
		}
		for j := 1; j <= w; j++ {
			if row[j] != '.' || row[j] != '*' {
				print("Invalid character in row. Please enter only '.' or '*'.\n")
				return
			}
			value := 0
			if row[j] == '*' {
				value = -1
			} else if row[j] == '.' {
				value = 0
			}
			matrix[i][j] = value
		}
	}
}

func PrintNum(n int) { // to print numbers
	decimal := 1
	for decimal <= n {
		decimal *= 10
	}
	decimal = decimal / 10 // to find decimal of n
	for decimal > 0 {
		ap.PutRune(rune(n/decimal) + '0')
		n = n % decimal
		decimal /= 10
	}
}
