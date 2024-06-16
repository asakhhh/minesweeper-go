package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

var (
	HEIGHT       int
	WIDTH        int
	BOMB_COUNT   int
	CLOSED_COUNT int
	MOVE_COUNT   int
)

func main() {
	PrintString("Hello! This is Minesweeper.\n")
	PrintString("You can create your own map or play a generated one.\n")
	PrintString("Note that both height and width of the map must be at least 3.\n")
	PrintString("Also, there must be at least two bombs. Good luck!\n\n")

	MODE := ChooseMode()

	var matrix [][]int

	HEIGHT, WIDTH = ReadHeightAndWidth()
	if MODE == 1 {
		BOMB_COUNT = CustomMap(HEIGHT, WIDTH, &matrix)
		CLOSED_COUNT = HEIGHT*WIDTH - BOMB_COUNT
	} else { // Random map
		BOMB_COUNT = GenerateRandomMap(HEIGHT, WIDTH, &matrix)
		CLOSED_COUNT = HEIGHT*WIDTH - BOMB_COUNT
	}
	if BOMB_COUNT < 2 {
		return
	}

	revealed := make([][]bool, HEIGHT+2)
	for i := range revealed {
		revealed[i] = make([]bool, WIDTH+2)
	}

	EnterMove(&matrix, &revealed)
}

func ReadHeightAndWidth() (int, int) {
	print("Enter the height and width of the map separated by a space:\n")

	height, width := ReadTwoNumbers()

	if height < 3 || width < 3 {
		PrintString("Sorry, your input is invalid. Please enter two numbers separated by a whitespace.\n")
		return ReadHeightAndWidth()
	}

	return height, width
}

func ChooseMode() int {
	PrintString("Choose a mode:\n1. Enter a custom map\n2. Generate a random map\n")
	PrintString("Enter your choice (1/2): ")

	var inp string
	var r rune
	fmt.Scanf("%c", &r)
	if r == '\n' {
		PrintString("Sorry, you have entered an invalid option. Please choose 1 or 2.\n\n")

		return ChooseMode()
	}
	for r != '\n' {
		inp += string(r)
		fmt.Scanf("%c", &r)
	}

	if len(inp) != 1 || !(inp[0] == '1' || inp[0] == '2') {
		PrintString("Sorry, you have entered an invalid option. Please choose 1 or 2.\n\n")

		return ChooseMode()
	}

	return int(rune(inp[0]) - '0')
}

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

func ReadTwoNumbers() (int, int) {
	var inp string
	var r rune
	fmt.Scanf("%c", &r)
	if r == '\n' {
		return -1, -1
	}
	for r != '\n' {
		inp += string(r)
		fmt.Scanf("%c", &r)
	}

	has_space := false
	var num1, num2 string
	for i, c := range inp {
		if c == ' ' {
			if has_space || i == 0 {
				return -1, -1
			}
			has_space = true
		} else if c < '0' || c > '9' {
			return -1, -1
		} else {
			if has_space {
				num2 += string(c)
			} else {
				num1 += string(c)
			}
		}
	}

	return StringToInt(num1), StringToInt(num2)
}

func PrintString(str string) {
	for _, c := range str {
		ap.PutRune(c)
	}
}

func CountAdjacentBombs(matrix *[][]int) {
	bomb_count := 0

	for i := 1; i <= HEIGHT; i++ {
		for j := 1; j <= WIDTH; j++ {
			bomb_count = 0
			if (*matrix)[i][j] == 0 {

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

func PrintMatrix(matrix *[][]int) {
	for _, v := range *matrix {
		for _, vv := range v {
			fmt.Printf("%d ", vv)
		}
		ap.PutRune('\n')
	}
}
