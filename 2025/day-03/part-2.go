package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalJoltage := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		battery := scanner.Bytes()

		left := 0

		for digit := 0; digit < 12; digit++ {
			maxNum := 0
			maxNumIndex := left

			for i := left; i < len(battery)-11+digit; i++ {
				num := int(battery[i]) - '0'

				if num > maxNum {
					maxNum = num
					maxNumIndex = i
				}
			}

			left = maxNumIndex + 1
			totalJoltage += maxNum * int(math.Pow10(11-digit))
		}
	}

	log.Printf("Password: %v\n", totalJoltage)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
