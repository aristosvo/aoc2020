package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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
	var sum1, sum3, sum5, sum7, sum1uneven int
	y := 1
	for scanner.Scan() {
		trees := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}
		sum1+= tree(trees, y, 1, false)
		sum3+= tree(trees, y, 3, false)
		sum5+= tree(trees, y, 5, false)
		sum7+= tree(trees, y, 7, false)
		sum1uneven+= tree(trees, y, 0.5, true)
		y++

	}
	fmt.Printf("Solution A = %d\n", sum3)
	fmt.Printf("1: %d\n", sum1)
	fmt.Printf("3: %d\n", sum3)
	fmt.Printf("5: %d\n", sum5)
	fmt.Printf("7: %d\n", sum7)
	fmt.Printf("1uneven: %d\n", sum1uneven)
	fmt.Printf("Solution B = %d\n", sum1 * sum3 * sum5 * sum7 * sum1uneven)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func tree(trees string, y int, step float64, unevenOnly bool) int {
	if unevenOnly && y % 2 == 0 {
		return 0
	}
	stepSize := 1.0
	if unevenOnly {
		stepSize = 0.5
	}
	length := len(trees)
	position := int((float64(y) - stepSize)*step) % length 

	if rune(trees[position]) == '#' {
		return 1
	}
	return 0
}
