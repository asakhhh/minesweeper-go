package main

import (
	"math/rand"
)

var bomb_coords []int

func AlreadyBomb(n int) bool {
	for _, v := range bomb_coords {
		if v == n {
			return true
		}
	}
	return false
}

func GenerateRandomMap(matrix *[][]int) int {
	*matrix = (*matrix)[:0]
	for i := 0; i < HEIGHT+2; i++ {
		var t []int
		for j := 0; j < WIDTH+2; j++ {
			t = append(t, 0)
		}
		*matrix = append(*matrix, t)
	}

	bomb_percentage := 12
	bomb_count := (HEIGHT*WIDTH*bomb_percentage + 99) / 100 // 13% of the map, rounded up

	for len(bomb_coords) < bomb_count {
		new_bomb := rand.Int31n(int32(HEIGHT * WIDTH))
		if !AlreadyBomb(int(new_bomb)) {
			bomb_coords = append(bomb_coords, int(new_bomb))
		}
	}

	for _, v := range bomb_coords {
		bomb_h := 1 + v/WIDTH
		bomb_w := 1 + v%WIDTH
		(*matrix)[bomb_h][bomb_w] = -1
	}

	return bomb_count
}
