package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

func StringToInt(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int(rune(c)-'0')
	}
	return n
}

func PrintString(str string) {
	for _, c := range str {
		ap.PutRune(c)
	}
}

func LengthOfNum(num int) int {
	p, res := 1, 1
	for p*10 <= num {
		p *= 10
		res++
	}
	return res
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

func PrintMatrix(matrix *[][]int, revealed *[][]bool) {
	for i := 1; i <= HEIGHT; i++ {
		for j := 1; j <= WIDTH; j++ {
			if !(*revealed)[i][j] {
				PrintString("X  ")
			} else {
				fmt.Printf("%d  ", (*matrix)[i][j])
			}
		}
		ap.PutRune('\n')
	}
}

func Greetings() {
	PrintString("Hello! This is Minesweeper.\n")
	PrintString("You can create your own map or play a generated one.\n\n")
	PrintString("!!! Note that both height and width of the map must be at least 3. !!!\n")
	PrintString("!!! Also, there must be at least two bombs. !!!\nGood luck!\n\n")
}
