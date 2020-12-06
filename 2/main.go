package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := flag.String("file", "input.txt", "file to analyze")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumA := 0
	sumB := 0
	for scanner.Scan() {
		fullLine := scanner.Text()
		s := strings.Fields(fullLine)
		minMax := strings.Split(s[0], "-")
		min, err := strconv.Atoi(minMax[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			log.Fatal(err)
		}
		count := 0
		checkedChar := rune(s[1][0])
		for _, i := range s[2] {
			if i == checkedChar {
				count++
			}
		}
		if count >= min && count <= max {
			sumA += 1
		}

		firstPosition := rune(s[2][min-1])
		lastPosition := rune(s[2][max-1])
		if (firstPosition == checkedChar) != (lastPosition == checkedChar) {
			sumB += 1
		}
	}

	fmt.Printf("Solution A = %d\n", sumA)
	fmt.Printf("Solution B = %d\n", sumB)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
