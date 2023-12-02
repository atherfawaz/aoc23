package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getGameID(str string) int {
	splitResult := strings.Split(str, " ")
	if len(splitResult) == 0 {
		fmt.Println("Error: Empty split result")
		return -1
	}

	num, err := strconv.Atoi(splitResult[len(splitResult)-1])
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return -1
	}

	return num
}

func SolveDay2() {
	RMAX := 12
	GMAX := 13
	BMAX := 14
	file, err := os.Open("day2/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		result := strings.Split(line, ":")
		gameId := getGameID(result[0])
		gameSets := strings.Split(result[1], ";")
		possible := true

		for _, set := range gameSets {
			pieces := strings.Split(strings.TrimSpace(set), ",")
			for _, piece := range pieces {
				g := strings.Split(strings.TrimSpace(piece), " ")
				pCount, err := strconv.Atoi(g[0])
				if err != nil {
					fmt.Println("Error converting to int:", err)
					continue
				}
				switch g[1] {
				case "green":
					if pCount > GMAX {
						possible = false
					}
				case "red":
					if pCount > RMAX {
						possible = false
					}
				case "blue":
					if pCount > BMAX {
						possible = false
					}
				}
			}
			if !possible {
				break
			}
		}
		if possible {
			total += gameId
		}
	}
	fmt.Printf("Answer to Day 2: %d\n", total)
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
