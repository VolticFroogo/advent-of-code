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

	ids := []int{}

	for scanner.Scan() {
		idText := scanner.Text()
		id, err := strconv.Atoi(idText)
		if err != nil {
			log.Fatal(err)
		}

		ids = append(ids, id)
	}

	slices.Sort(ids)

	goodIngredients := 0
	currentRange := 0

out:
	for _, id := range ids {
		for id > ranges[currentRange].High {
			currentRange++
			if currentRange == len(ranges) {
				break out
			}
		}

		if id >= ranges[currentRange].Low {
			goodIngredients++
		}
	}

	log.Printf("Password: %v\n", goodIngredients)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
