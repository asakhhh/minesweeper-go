package main

import (
	"fmt"
)

func ReadLine() string {
	var inp string
	var r rune
	fmt.Scanf("%c", &r)
	for r != '\n' {
		inp += string(r)
		fmt.Scanf("%c", &r)
	}
	return inp
}

func ChooseMode() int {
	PrintString("Choose a mode:\n1. Enter a custom map\n2. Generate a random map\n")
	PrintString("Enter your choice (1/2): ")

	inp := ReadLine()

	if len(inp) != 1 || !(inp[0] == '1' || inp[0] == '2') {
		PrintString("Sorry, you have entered an invalid option. Please choose 1 or 2.\n\n")

		return ChooseMode()
	}

	return int(rune(inp[0]) - '0')
}

func ReadTwoNumbers() (int, int) {
	inp := ReadLine()
	has_space := false
	var num1, num2 string

	if len(inp) == 0 {
		return -1, -1
	}

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

	if len(num2) == 0 {
		return -1, -1
	}

	return StringToInt(num1), StringToInt(num2)
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
