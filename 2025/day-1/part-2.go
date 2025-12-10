package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	pos := 50
	zeroHits := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rotation := scanner.Text()
		rotationLetter := rotation[0]
		rotationNumber, err := strconv.Atoi(rotation[1:])
		if err != nil {
			log.Fatal(err)
		}

		rotationsOldPos := math.Floor(float64(pos) / 100)

		if rotationLetter == 'R' {
			pos += rotationNumber
		} else {
			if pos%100 == 0 {
				zeroHits--
			}

			pos -= rotationNumber

			if pos%100 == 0 {
				zeroHits++
			}
		}

		rotationsPos := math.Floor(float64(pos) / 100)
		zeroHits += int(math.Abs(rotationsPos - rotationsOldPos))
	}

	log.Printf("Password: %v\n", zeroHits)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
