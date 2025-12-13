package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	numsMatrix := [][]int{}
	operations := []rune{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if text[0] == '+' || text[0] == '*' {
			for _, char := range text {
				if char != ' ' {
					operations = append(operations, char)
				}
			}

			break
		}

		nums := []int{}

		split := strings.SplitSeq(text, " ")
		for numString := range split {
			if numString != "" {
				num, err := strconv.Atoi(numString)
				if err != nil {
					log.Fatal(err)
				}

				nums = append(nums, num)
			}
		}

		numsMatrix = append(numsMatrix, nums)
	}

	sum := 0

	for i := range operations {
		if operations[i] == '*' {
			value := 1
			for j := range numsMatrix {
				value *= numsMatrix[j][i]
			}
			sum += value
		} else {
			for j := range numsMatrix {
				sum += numsMatrix[j][i]
			}
		}
	}

	log.Printf("Password: %v\n", sum)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
