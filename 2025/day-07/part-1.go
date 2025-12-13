package main

import (
	"bufio"
	"log"
	"os"
)

// func printLine(bytes *[]byte, beams *[]bool) {
// 	for i := range *beams {
// 		if (*beams)[i] {
// 			fmt.Print("|")
// 		} else {
// 			fmt.Printf("%c", (*bytes)[i])
// 		}
// 	}
//
// 	fmt.Print("\n")
// }

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	bytes := scanner.Bytes()
	beams := make([]bool, len(bytes))

	for i, char := range bytes {
		if char == 'S' {
			beams[i] = true
		}
	}

	// printLine(&bytes, &beams)

	splits := 0

	for scanner.Scan() {
		bytes = scanner.Bytes()

		for i, char := range bytes {
			if char == '^' && beams[i] {
				splits++
				beams[i-1] = true
				beams[i] = false
				beams[i+1] = true
			}
		}

		// printLine(&bytes, &beams)
	}

	log.Printf("Password: %v\n", splits)

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}
