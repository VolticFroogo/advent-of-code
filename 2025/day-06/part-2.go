package main

import (
	"bufio"
	"log"
	"os"
)

func valueForOperation(operation byte) int {
	if operation == '+' {
		return 0
	}

	return 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	matrix := [][]byte{}
	operations := []byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()

		if bytes[0] == '+' || bytes[0] == '*' {
			for _, char := range bytes {
				if char != ' ' {
					operations = append(operations, char)
				}
			}
		}

		bytesCopy := make([]byte, len(bytes))
		copy(bytesCopy, bytes)
		matrix = append(matrix, bytesCopy)
	}

	sum := 0
	operation := len(operations) - 1
	value := valueForOperation(operations[operation])

	for x := len(matrix[0]) - 1; x >= 0; x-- {
		num := 0
		for y := 0; y < len(matrix)-1; y++ {
			if matrix[y][x] == ' ' {
				continue
			}

			num = num*10 + int(matrix[y][x]) - '0'
		}

		if operations[operation] == '*' {
			value *= num
		} else {
			value += num
		}

		if matrix[len(matrix)-1][x] != ' ' {
			sum += value

			x--
			operation--
			if operation < 0 {
				break
			}
			value = valueForOperation(operations[operation])
		}
	}

	log.Printf("Password: %v\n", sum)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
