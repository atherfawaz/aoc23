package day1

import (
	"bufio"
	"fmt"
	"os"
)

func isASCIIDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func SolveDay1() {

	file, err := os.Open("day1/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	start := 0
	last_ix := 0
	first := -1
	last := -1

	for scanner.Scan() {
		line := scanner.Text()
		first = -1
		last = -1
		start = 0
		last_ix = len(line) - 1
		for start <= last_ix {
			if isASCIIDigit(line[start]) {
				if first == -1 {
					first = int(line[start] - '0')
				} else {
					last = int(line[start] - '0')
				}
			}
			start++
		}
		if first != -1 {
			total += first * 10
		}
		if last != -1 {
			total += last
		} else {
			total += first
		}
	}
	fmt.Printf("Answer to Day 1: %d\n", total)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
