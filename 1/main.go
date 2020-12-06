package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var filename = flag.String("file", "input.txt", "file to analyze")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var output []int
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		for _, oldExpense := range output {
			if oldExpense + expense == 2020 {
				total := oldExpense * expense
				fmt.Printf("Result first quest: %d\n", total)
			}
		}

		for i, oldestExpense := range output {
			for _, olderExpense := range output[0:i] {
				if oldestExpense + olderExpense + expense == 2020 {
					total := olderExpense * oldestExpense * expense
					fmt.Printf("Result second quest: %d\n", total)
				}
			}
		}



		output = append(output, expense)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}