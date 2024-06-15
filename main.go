package main

import (
	"fmt"

	"github.com/alem-platform/ap"
)

var (
	HEIGHT int
	WIDTH  int
)

func main() {
	PrintString("Hello! This is Minesweeper.\n")
	PrintString("You can create your own map or play a generated one.\n\n")

	MODE := ChooseMode()

	var matrix [][]int

	if MODE == 1 {
		HEIGHT, WIDTH = ReadHeightAndWidth()
		CustomMap(HEIGHT, WIDTH, &matrix)
	} else {
	}
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
