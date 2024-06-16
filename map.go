package main

// import "github.com/alem-platform/ap"

// // colors for characters
// var (
// 	RESET  = []rune{'\033', '[', '0', 'm'}
// 	RED    = []rune{'\033', '[', '3', '1', 'm'}
// 	WHITE  = []rune{'\033', '[', '9', '7', 'm'}
// 	BLUE   = []rune{'\033', '[', '3', '4', 'm'}
// 	YELLOW = []rune{'\033', '[', '3', '3', 'm'}

// 	colorCodes = [][]rune{RED, RESET, BLUE, YELLOW}
// )

// // colors for background
// var (
// 	BgRed    = []rune("\033[41m")
// 	BgGreen  = []rune("\033[42m")
// 	BgYellow = []rune("\033[1;33;43m")
// 	BgBlue   = []rune("\033[44m")
// 	BgPurple = []rune("\033[45m")
// 	BgCyan   = []rune("\033[46m")
// 	BgWhite  = []rune("\033[47m")

// 	colorCodesBG = [][]rune{BgRed, RESET, BgBlue, BgYellow, BgPurple, BgCyan}
// )

// // function that colorizes Foreground
// func colorizeFG(val int) {
// 	if val >= 0 && val < len(colorCodes) {
// 		for _, r := range colorCodes[val] {
// 			ap.PutRune(r)
// 		}
// 	}
// }

// // function that colorizes background
// func colorizeBG(val int) {
// 	if val >= 0 && val < len(colorCodesBG) {
// 		for _, r := range colorCodesBG[val] {
// 			ap.PutRune(r)
// 		}
// 	}
// }

// func printMap(height, width int, value [][]int) {
// 	horiz_coord := make([][7]rune, width)
// 	horiz_coord_count(width, &horiz_coord)

// 	length := 1
// 	p := 1
// 	for p*10 <= height {
// 		p *= 10
// 		length++
// 	}

// 	for h := 0; h < (3*height)+1; h++ {
// 		// Print vertical coordinates
// 		if (h+1)%3 == 0 {
// 			printVerticalCoord((h+1)/3, length)
// 		} else {
// 			for i := 0; i < length+2; i++ {
// 				ap.PutRune(' ')
// 			}
// 		}
// 		if h == 0 {
// 			printHorizontalCoord(width, &horiz_coord)
// 		}
// 		// prints row of the map
// 		printRow(width, length, value, h)
// 	}
// }

// func printRow(width int, length int, value [][]int, h int) {
// 	for w := 0; w < (8*width)+1; w++ {
// 		if w%8 == 0 && h > 0 {
// 			ap.PutRune('|')
// 		} else if h == 0 {
// 			if w == 0 {
// 				for i := 0; i < length+2; i++ {
// 					ap.PutRune(' ')
// 				}
// 			}
// 			if w > 0 && w != 8*width {
// 				ap.PutRune('_')
// 			} else {
// 				ap.PutRune(' ')
// 			}
// 		} else {
// 			printCell(value[(h-1)/3][w/8], h%3, w%8) // Print cell content
// 		}
// 	}
// 	ap.PutRune('\n')
// }

// // printCell function prints a single cell of a row based on its value
// func printCell(value, y, x int) {
// 	// colorizeBG(value)
// 	symbols := []rune{wall, ' ', player, award}

// 	if value == 0 { // print value of change symbol for wall
// 		ap.PutRune(symbols[value])
// 	} else if value == 1 { // print empty cell
// 		if y == 0 {
// 			ap.PutRune('_')
// 		} else {
// 			ap.PutRune(symbols[value])
// 		}
// 	} else { //  print value of change symbol
// 		if y == 2 && x == 4 {
// 			ap.PutRune(symbols[value])
// 		} else if y == 0 {
// 			ap.PutRune('_')
// 		} else {
// 			ap.PutRune(symbols[1])
// 		}
// 	}
// 	// colorizeBG(1)
// }

// func printHorizontalCoord(width int, horiz_coord *[][7]rune) {
// 	for w := 0; w < (8*width)+1; w++ {
// 		if w%8 != 0 {
// 			ap.PutRune((*horiz_coord)[w/8][w%8-1])
// 		} else {
// 			ap.PutRune(' ')
// 		}
// 	}
// 	ap.PutRune('\n')
// }

// func printVerticalCoord(num, length int) {
// 	p, curlength := 1, 1

// 	for p*10 <= num {
// 		p *= 10
// 		curlength++
// 	}

// 	// part that prints spaces so that matrix numbers arr printed correctly
// 	for i := 0; i < length-curlength; i++ {
// 		ap.PutRune(' ')
// 	}

// 	for ; p > 0; p /= 10 {
// 		ap.PutRune(rune('0' + (num/p)%10))
// 	}
// 	ap.PutRune(' ')
// }
