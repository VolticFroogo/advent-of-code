package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	bytes := scanner.Bytes()
	beamTimelines := make([]int, len(bytes))

	for i, char := range bytes {
		if char == 'S' {
			beamTimelines[i] = 1
		}
	}

	for scanner.Scan() {
		bytes = scanner.Bytes()

		for i, char := range bytes {
			if char == '^' && beamTimelines[i] != 0 {
				beamTimelines[i-1] += beamTimelines[i]
				beamTimelines[i+1] += beamTimelines[i]
				beamTimelines[i] = 0
			}
		}
	}

	totalTimeslines := 0
	for _, timeline := range beamTimelines {
		totalTimeslines += timeline
	}

	log.Printf("Password: %v\n", totalTimeslines)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
