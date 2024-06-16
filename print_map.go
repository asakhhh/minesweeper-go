package main

import "github.com/alem-platform/ap"

// func PrintMap(matrix *[][]int, revealed *[][]bool) {
// 	PrintMatrix(matrix, revealed)
// }

func PrintMap(matrix *[][]int, revealed *[][]bool) {
	length := LengthOfNum(HEIGHT)

	for h := 0; h < (3*HEIGHT)+1; h++ {
		// Print vertical coordinates
		if h%3 == 2 {
			PrintYCoord((h+1)/3, length)
		} else {
			for i := 0; i < length+2; i++ {
				ap.PutRune(' ')
			}
		}
		if h == 0 {
			PrintXCoord()
			ap.PutRune('\n')
		}
		// prints row of the map
		printRow(length, matrix, revealed, h)
	}
}

func printRow(length int, matrix *[][]int, revealed *[][]bool, h int) {
	for w := 0; w < (8*WIDTH)+1; w++ {
		if w%8 == 0 && h > 0 {
			ap.PutRune('|')
		} else if h == 0 {
			if w == 0 {
				for i := 0; i < length+2; i++ {
					ap.PutRune(' ')
				}
			}
			if w > 0 && w < 8*WIDTH {
				ap.PutRune('_')
			} else {
				ap.PutRune(' ')
			}
		} else {
			printCell((*matrix)[(h+2)/3][w/8+1], (*revealed)[(h+2)/3][w/8+1], h%3, w%8) // Print cell content
		}
	}
	ap.PutRune('\n')
}

func printCell(value int, revealed bool, x, y int) {
	if !revealed {
		ap.PutRune('X')
	} else if value == -1 {
		if x == 2 && y == 4 {
			ap.PutRune('*')
		} else if x == 0 {
			ap.PutRune('_')
		} else {
			ap.PutRune(' ')
		}
	} else {
		if x == 2 && y == 4 {
			ap.PutRune(rune('0' + value))
		} else if x == 0 {
			ap.PutRune('_')
		} else {
			ap.PutRune(' ')
		}
	}
}

func PrintYCoord(num, length int) {
	p, curlength := 1, LengthOfNum(num)

	ap.PutRune(' ')
	// part that prints spaces so that matrix numbers arr printed correctly
	for i := 0; i < length-curlength; i++ {
		ap.PutRune(' ')
	}

	for p > 0 {
		ap.PutRune(rune('0' + (num/p)%10))
		p /= 10
	}
	ap.PutRune(' ')
}

func LengthOfNum(num int) int {
	p, res := 1, 1
	for p*10 <= num {
		p *= 10
		res++
	}
	return res
}

func PrintXCoord() {
	PrintString("    ")
	for i := 1; i <= WIDTH; i++ {
		PrintNum(i)
		for z := 0; z < 8-LengthOfNum(i); z++ {
			ap.PutRune(' ')
		}
	}
}
