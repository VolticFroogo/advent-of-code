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

	totalJoltage := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		battery := scanner.Bytes()

		maxLeftNum := 0
		maxRightNum := 0

		for i, char := range battery {
			num := int(char) - '0'

			if num > maxLeftNum && i != len(battery)-1 {
				maxLeftNum = num
				maxRightNum = 0
			} else if num > maxRightNum {
				maxRightNum = num
			}
		}

		totalJoltage += maxLeftNum*10 + maxRightNum
	}

	log.Printf("Password: %v\n", totalJoltage)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
