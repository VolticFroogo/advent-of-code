package main

import (
	"bufio"
	"log"
	"os"
)

func isPaperInSpotOobSafe(x, y int, matrix *[][]bool) int {
	// Return 0 if OOB
	if y < 0 || x < 0 || y >= len(*matrix) || x >= len((*matrix)[y]) {
		return 0
	}

	if (*matrix)[y][x] {
		return 1
	}

	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	matrix := make([][]bool, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowBytes := scanner.Bytes()
		row := make([]bool, len(rowBytes))

		for i, char := range rowBytes {
			if char == '@' {
				row[i] = true
			}
		}

		matrix = append(matrix, row)
	}

	reachable := 0

	for y := range matrix {
		for x := range matrix[y] {
			if !matrix[y][x] {
				continue
			}

			adjacent := 0
			adjacent += isPaperInSpotOobSafe(x-1, y-1, &matrix) // -1, -1 TL
			adjacent += isPaperInSpotOobSafe(x, y-1, &matrix)   //  0, -1 TM
			adjacent += isPaperInSpotOobSafe(x+1, y-1, &matrix) // +1, -1 TR
			adjacent += isPaperInSpotOobSafe(x-1, y, &matrix)   // -1,  0 ML
			adjacent += isPaperInSpotOobSafe(x+1, y, &matrix)   // +1,  0 MR
			adjacent += isPaperInSpotOobSafe(x-1, y+1, &matrix) // -1, +1 BL
			adjacent += isPaperInSpotOobSafe(x, y+1, &matrix)   //  0, +1 BM
			adjacent += isPaperInSpotOobSafe(x+1, y+1, &matrix) // +1, +1 BR

			if adjacent < 4 {
				reachable++
			}
		}
	}

	log.Printf("Password: %v\n", reachable)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
