package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToIntArray(arr []string) []int {
	intArray := []int{}
	for _, str := range arr {
		if str == "" {
			continue
		}
		intValue, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			fmt.Println("Error parsing string:", err)
			return intArray
		}
		intArray = append(intArray, intValue)
	}
	return intArray
}

func getTimesAndDistances() (times []int, distances []int, err error) {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var splits []string

	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid format in line 1: %s", line)
		}

		splits = strings.Split(parts[1], " ")
		times = convertToIntArray(splits)
	} else {
		return nil, nil, fmt.Errorf("error reading line 1: %w", scanner.Err())
	}

	if scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid format in line 2: %s", line)
		}
		splits = strings.Split(parts[1], " ")
		distances = convertToIntArray(splits)
	} else {
		return nil, nil, fmt.Errorf("file contains only one line")
	}

	return times, distances, nil
}
func SolveDay6() {
	total := 1
	times, distances, _ := getTimesAndDistances()
	ways := []int{}

	for i, time := range times {
		currWay := 0
		// distance = speed * time
		for j := 1; j < time; j++ {
			travelDistance := j * (time - j)
			if travelDistance > distances[i] {
				currWay++
			}
		}
		ways = append(ways, currWay)
	}
	for i := 0; i < len(ways); i++ {
		total = total * ways[i]
	}
	fmt.Printf("Answer to Day 6: %d\n", total)
}
