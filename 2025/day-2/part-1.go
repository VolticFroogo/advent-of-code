package main

import (
	"bufio"
	"cmp"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Range struct {
	Low  int
	High int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputText := scanner.Text()
	rangesText := strings.Split(inputText, ",")
	ranges := make([]Range, len(rangesText))

	for i, v := range rangesText {
		splitRange := strings.Split(v, "-")
		low, err := strconv.Atoi(splitRange[0])
		if err != nil {
			log.Fatal(err)
		}
		high, err := strconv.Atoi(splitRange[1])
		if err != nil {
			log.Fatal(err)
		}

		ranges[i] = Range{
			Low:  low,
			High: high,
		}
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.Low, b.Low)
	})

	currentRange := 0
	invalidIdsSum := 0
	i := 1

out:
	for {
		// Repeat the number using just maths (e.g. 123 -> 123123)
		num := i*int(math.Pow10(int(math.Floor(math.Log10(float64(i))))+1)) + i

		for num > ranges[currentRange].High {
			currentRange++

			if currentRange == len(ranges) {
				break out
			}
		}

		if num >= ranges[currentRange].Low {
			invalidIdsSum += num
		}

		i++
	}

	log.Printf("Password: %v\n", invalidIdsSum)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
