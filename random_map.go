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

func GenerateRandomSize() (int, int) {
	return int(rand.Int31n(8) + 3), int(rand.Int31n(8) + 3)
}

func GenerateRandomMap(h, w int, matrix *[][]int) int {
	*matrix = (*matrix)[:0]
	for i := 0; i < h+2; i++ {
		var t []int
		for j := 0; j < w+2; j++ {
			t = append(t, 0)
		}
		*matrix = append(*matrix, t)
	}

	bomb_count := (h*w*13 + 99) / 100 // 13% of the map, rounded up

	for len(bomb_coords) < bomb_count {
		new_bomb := rand.Int31n(int32(h * w))
		if !AlreadyBomb(int(new_bomb)) {
			bomb_coords = append(bomb_coords, int(new_bomb))
		}
	}

	for _, v := range bomb_coords {
		bomb_h := 1 + v/w
		bomb_w := 1 + v%w
		(*matrix)[bomb_h][bomb_w] = -1
	}

	return bomb_count
}
