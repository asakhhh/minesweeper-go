package main

import "github.com/alem-platform/ap"

func printMap(height, width int, value [][]int, revealed [][]bool) {
	length := 1
	p := 1
	for p*10 <= height {
		p *= 10
		length++
	}

	for h := 0; h < (3*height)+1; h++ {
		// Print vertical coordinates
		if (h+1)%3 == 0 {
			printVerticalCoord((h+1)/3, length)
		} else {
			for i := 0; i < length+2; i++ {
				ap.PutRune(' ')
			}
		}
		if h == 0 {
		}
		// prints row of the map
		printRow(width, length, matrix, h)
	}
}

func printRow(width int, length int, value int, revealed [][]bool, h int) {
	for w := 0; w < (8*width)+1; w++ {
		if w%8 == 0 && h > 0 {
			ap.PutRune('|')
		} else if h == 0 {
			if w == 0 {
				for i := 0; i < length+2; i++ {
					ap.PutRune(' ')
				}
			}
			if w > 0 && w != 8*width {
				ap.PutRune('_')
			} else {
				ap.PutRune(' ')
			}
		} else {
			printCell(matrix[(h-1)/3][w/8], revealed, h%3, w%8) // Print cell content
		}
	}
	ap.PutRune('\n')
}

func printCell(value int, revealed [][]bool, x, y int) {
	if !revealed[x][y] {
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
			ap.PutRune(value)
		} else if x == 0 {
			ap.PutRune('_')
		} else {
			ap.PutRune(' ')
		}
	}
}

func printVerticalCoord(num, length int) {
	p, curlength := 1, 1

	for p*10 <= num {
		p *= 10
		curlength++
	}

	// part that prints spaces so that matrix numbers arr printed correctly
	for i := 0; i < length-curlength; i++ {
		ap.PutRune(' ')
	}

	for ; p > 0; p /= 10 {
		ap.PutRune(rune('0' + (num/p)%10))
	}
	ap.PutRune(' ')
}
