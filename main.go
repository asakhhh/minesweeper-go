package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

var (
	HEIGHT       int
	WIDTH        int
	CLOSED_COUNT int
)

func main() {
	PrintString("Hello! This is Minesweeper.\n")
	PrintString("You can create your own map or play a generated one.\n\n")

	MODE := ChooseMode()

	var matrix [][]int

	if MODE == 1 {
		HEIGHT, WIDTH = ReadHeightAndWidth()
		CLOSED_COUNT = HEIGHT*WIDTH - CustomMap(HEIGHT, WIDTH, &matrix)
	} else { // Random map
		HEIGHT, WIDTH = GenerateRandomSize()
		CLOSED_COUNT = HEIGHT*WIDTH - GenerateRandomMap(HEIGHT, WIDTH, &matrix)
	}
	if CLOSED_COUNT > HEIGHT*WIDTH {
		return
	}

	revealed := make([][]bool, HEIGHT+2)
	for i := range revealed {
		revealed[i] = make([]bool, WIDTH+2)
	}

	EnterMove(&matrix, &revealed)
}

func ReadHeightAndWidth() (int, int) {
	var height, width int

	print("Enter the number of height and width of the map (e.g., '2 3'):\n")

	_, err := fmt.Scanf("%d %d", &height, &width)
	if err != nil || height < 3 || width < 3 {
		PrintString("Invalid format for map dimensions. Please enter two positive integers more that 3 separated by a space.")
		return -1, -1
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
