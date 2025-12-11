package main

import (
	"bufio"
	"log"
	"os"
)

func updateAdjacent(x, y int, delta int8, adjacentMatrix *[][]int8, matrix *[][]bool, toRemove *map[uint16]bool) {
	updateSpot(x-1, y-1, delta, adjacentMatrix, matrix, toRemove) // -1, -1 TL
	updateSpot(x, y-1, delta, adjacentMatrix, matrix, toRemove)   //  0, -1 TM
	updateSpot(x+1, y-1, delta, adjacentMatrix, matrix, toRemove) // +1, -1 TR
	updateSpot(x-1, y, delta, adjacentMatrix, matrix, toRemove)   // -1,  0 ML
	updateSpot(x+1, y, delta, adjacentMatrix, matrix, toRemove)   // +1,  0 MR
	updateSpot(x-1, y+1, delta, adjacentMatrix, matrix, toRemove) // -1, +1 BL
	updateSpot(x, y+1, delta, adjacentMatrix, matrix, toRemove)   //  0, +1 BM
	updateSpot(x+1, y+1, delta, adjacentMatrix, matrix, toRemove) // +1, +1 BR
}

func updateSpot(x, y int, delta int8, adjacentMatrix *[][]int8, matrix *[][]bool, toRemove *map[uint16]bool) {
	// Return 0 if OOB
	if y < 0 || x < 0 || y >= len(*adjacentMatrix) || x >= len((*adjacentMatrix)[y]) {
		return
	}

	(*adjacentMatrix)[y][x] += delta

	if matrix == nil {
		return
	}

	if !(*matrix)[y][x] {
		return
	}

	if (*adjacentMatrix)[y][x] < 4 {
		(*toRemove)[uint16(y<<8|x)] = true
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	matrix := make([][]bool, 0)
	adjacentMatrix := make([][]int8, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowBytes := scanner.Bytes()
		row := make([]bool, len(rowBytes))
		adjacentRow := make([]int8, len(rowBytes))

		for i, char := range rowBytes {
			if char == '@' {
				row[i] = true
			}
		}

		matrix = append(matrix, row)
		adjacentMatrix = append(adjacentMatrix, adjacentRow)
	}

	// Build adjacent matrix
	for y := range matrix {
		for x := range matrix[y] {
			if !matrix[y][x] {
				continue
			}

			updateAdjacent(x, y, 1, &adjacentMatrix, nil, nil)
		}
	}

	reachable := 0
	toRemove := map[uint16]bool{}

	for y := range matrix {
		for x := range matrix[y] {
			if !matrix[y][x] {
				continue
			}

			if adjacentMatrix[y][x] >= 4 {
				continue
			}

			toRemove[uint16(y<<8|x)] = true
		}
	}

	for len(toRemove) != 0 {
		for pos := range toRemove {
			delete(toRemove, pos)

			y := int(pos >> 8)
			x := int(pos & 0xFF)

			matrix[y][x] = false
			updateAdjacent(x, y, -1, &adjacentMatrix, &matrix, &toRemove)

			reachable++
		}
	}

	log.Printf("Password: %v\n", reachable)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
