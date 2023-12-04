package day4

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findAllNumbers(input string) []int {
	re := regexp.MustCompile(`\b\d+\b`)
	matches := re.FindAllString(input, -1)
	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			numbers = append(numbers, num)
		}
	}

	return numbers
}

func getHashSetOfNumbers(numbers []int) util.HashSetInt {
	hashSet := make(util.HashSetInt)
	for _, val := range numbers {
		hashSet.Add(val)
	}
	return hashSet
}

func SolveDay4() {
	total := 0
	file, err := os.Open("day4/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runningTotal := 0
		card := strings.Split(line, "|")
		myNumbers := findAllNumbers(strings.TrimSpace(card[1]))
		winningNumbers := getHashSetOfNumbers(findAllNumbers(strings.Split(strings.TrimSpace(card[0]), ":")[1]))

		for _, num := range myNumbers {
			if winningNumbers.Contains(num) {
				if runningTotal == 0 {
					runningTotal++
				} else {
					runningTotal = runningTotal * 2
				}
			}
		}
		total += runningTotal
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Answer to Day 4: %d\n", total)
}
