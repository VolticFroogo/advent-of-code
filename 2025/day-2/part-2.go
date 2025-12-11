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

	invalidIdsSum := 0
	seenInvalidIds := map[int]bool{}
	repeats := 2

	for {
		currentRange := 0
		i := 1

	out:
		for {
			num, _ := strconv.Atoi(strings.Repeat(strconv.Itoa(i), repeats))

			for num > ranges[currentRange].High {
				currentRange++

				if currentRange == len(ranges) {
					break out
				}
			}

			if num >= ranges[currentRange].Low && !seenInvalidIds[num] {
				invalidIdsSum += num
				seenInvalidIds[num] = true
			}

			i++
		}

		repeats++
		if int(math.Pow10(repeats)) > ranges[len(ranges)-1].High {
			break
		}
	}

	log.Printf("Password: %v\n", invalidIdsSum)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
