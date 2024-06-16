package main

var (
	HEIGHT       int
	WIDTH        int
	BOMB_COUNT   int
	CLOSED_COUNT int
	MOVE_COUNT   int
)

func main() {
	Greetings()

	MODE := ChooseMode()

	var matrix [][]int

	HEIGHT, WIDTH = ReadHeightAndWidth()

	if MODE == 1 { // Custom map
		BOMB_COUNT = CustomMap(&matrix)
	} else { // Random map
		BOMB_COUNT = GenerateRandomMap(&matrix)
	}
	CLOSED_COUNT = HEIGHT*WIDTH - BOMB_COUNT

	if BOMB_COUNT < 2 {
		PrintString("There must be at least two bombs.\n")
	}

	revealed := make([][]bool, HEIGHT+2)
	for i := range revealed {
		revealed[i] = make([]bool, WIDTH+2)
	}

	CountAdjacentBombs(&matrix)
	EnterMove(&matrix, &revealed)
}
