package day3

import (
	"bufio"
	"fmt"
	"os"
)

func isASCIIDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func getSchematic() [][]rune {
	var schematic [][]rune
	file, err := os.Open("day3/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return schematic
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return schematic
}

func SolveDay3() {
	total := 0
	schematic := getSchematic()
	for i, row := range schematic {
		for j, char := range row {
			fmt.Printf("i: %d, j: %d, char: %c\n", i, j, char)
			if isASCIIDigit(char) {

			}
		}
	}
	fmt.Printf("Answer to Day 3: %d\n", total)
}
