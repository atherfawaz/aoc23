package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func isPartNumber(i int, height int, width int, k int, schematic [][]rune) bool {

	checkTop := func() bool {
		return i != 0 && !isASCIIDigit(schematic[i-1][k]) && schematic[i-1][k] != '.'
	}

	checkTopLeft := func() bool {
		return i != 0 && k != 0 && !isASCIIDigit(schematic[i-1][k-1]) && schematic[i-1][k-1] != '.'
	}

	checkTopRight := func() bool {
		return i != 0 && width-k > 1 && !isASCIIDigit(schematic[i-1][k+1]) && schematic[i-1][k+1] != '.'
	}

	checkBottomLeft := func() bool {
		return k != 0 && height-i > 1 && !isASCIIDigit(schematic[i+1][k-1]) && schematic[i+1][k-1] != '.'
	}

	checkBottomRight := func() bool {
		return width-k > 1 && height-i > 1 && !isASCIIDigit(schematic[i+1][k+1]) && schematic[i+1][k+1] != '.'
	}

	checkBottom := func() bool {
		return height-i > 1 && !isASCIIDigit(schematic[i+1][k]) && schematic[i+1][k] != '.'
	}

	checkLeft := func() bool {
		return k != 0 && !isASCIIDigit(schematic[i][k-1]) && schematic[i][k-1] != '.'
	}

	checkRight := func() bool {
		return width-k > 1 && !isASCIIDigit(schematic[i][k+1]) && schematic[i][k+1] != '.'
	}

	return checkTop() || checkTopLeft() || checkTopRight() || checkBottomLeft() || checkBottomRight() || checkBottom() || checkLeft() || checkRight()
}

func SolveDay3() {
	total := 0
	schematic := getSchematic()
	height := len(schematic)
	width := len(schematic[0])

	for i, row := range schematic {
		for j, char := range row {
			if isASCIIDigit(char) {
				var charArr []rune
				isPartNum := false
				if j == 0 || (j > 0 && !isASCIIDigit(schematic[i][j-1])) {
					for k := j; k < width && isASCIIDigit(schematic[i][k]); k++ {
						charArr = append(charArr, schematic[i][k])
						if !isPartNum {
							isPartNum = isPartNumber(i, height, width, k, schematic)
						}
					}
				}
				if len(charArr) == 0 {
					continue
				}
				number, err := strconv.Atoi(string(charArr))
				if err != nil {
					fmt.Println("Error converting to int:", err)
					continue
				}
				if isPartNum {
					total += number
				}
			}
		}
	}

	fmt.Printf("Answer to Day 3: %d\n", total)
}
