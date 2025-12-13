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

type Vector struct {
	X, Y, Z int
}

func distance(a, b Vector) float64 {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z

	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

type Distance struct {
	A, B     int
	Distance float64
}

type Circuit struct {
	Merged *Circuit
	Count  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	vectors := []Vector{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, ",")

		x, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		z, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatal(err)
		}

		vectors = append(vectors, Vector{
			X: x,
			Y: y,
			Z: z,
		})
	}

	distances := []Distance{}

	for i := range vectors {
		for j := i + 1; j < len(vectors); j++ {
			distances = append(distances, Distance{
				A:        i,
				B:        j,
				Distance: distance(vectors[i], vectors[j]),
			})
		}
	}

	slices.SortFunc(distances, func(a, b Distance) int {
		return cmp.Compare(a.Distance, b.Distance)
	})

	circuits := map[int]*Circuit{}

	for _, distance := range distances[:1000] {
		circuitA, containsA := circuits[distance.A]
		circuitB, containsB := circuits[distance.B]

		if containsA && containsB {
			for circuitA.Merged != nil {
				circuitA = circuitA.Merged
			}
			for circuitB.Merged != nil {
				circuitB = circuitB.Merged
			}

			if circuitA != circuitB {
				circuit := Circuit{
					Count: circuitA.Count + circuitB.Count,
				}
				circuitA.Merged = &circuit
				circuitB.Merged = &circuit
			}
		} else if containsA {
			for circuitA.Merged != nil {
				circuitA = circuitA.Merged
			}
			circuitA.Count++
			circuits[distance.B] = circuitA
		} else if containsB {
			for circuitB.Merged != nil {
				circuitB = circuitB.Merged
			}
			circuitB.Count++
			circuits[distance.A] = circuitB
		} else {
			circuit := Circuit{
				Count: 2,
			}
			circuits[distance.A] = &circuit
			circuits[distance.B] = &circuit
		}
	}

	seenCircuits := map[*Circuit]bool{}
	circuitCounts := []int{}

	for _, circuit := range circuits {
		for circuit.Merged != nil {
			circuit = circuit.Merged
		}

		if seenCircuits[circuit] {
			continue
		}
		seenCircuits[circuit] = true

		circuitCounts = append(circuitCounts, circuit.Count)
	}

	slices.Sort(circuitCounts)

	circuitMultiple := 1
	for i := len(circuitCounts) - 1; i > len(circuitCounts)-4; i-- {
		circuitMultiple *= circuitCounts[i]
	}

	log.Printf("Password: %v\n", circuitMultiple)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
