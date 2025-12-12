package main

import (
	"bufio"
	"cmp"
	"log"
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

	ranges := []Range{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rangeText := scanner.Text()
		if rangeText == "" {
			break
		}

		splitRange := strings.Split(rangeText, "-")
		low, err := strconv.Atoi(splitRange[0])
		if err != nil {
			log.Fatal(err)
		}
		high, err := strconv.Atoi(splitRange[1])
		if err != nil {
			log.Fatal(err)
		}

		ranges = append(ranges, Range{
			Low:  low,
			High: high,
		})
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return cmp.Compare(a.Low, b.Low)
	})

	freshIngredients := 0
	highestAdded := 0

	for _, idRange := range ranges {
		freshIngredients += max(idRange.High-max(highestAdded, idRange.Low-1), 0)
		highestAdded = max(highestAdded, idRange.High)
	}

	log.Printf("Password: %v\n", freshIngredients)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
