package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readNumbersToArrays(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var leftNumbers, rightNumbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		if len(values) >= 2 {
			val1, err1 := strconv.Atoi(values[0])
			val2, err2 := strconv.Atoi(values[1])

			if err1 == nil && err2 == nil {
				leftNumbers = append(leftNumbers, val1)
				rightNumbers = append(rightNumbers, val2)
			} else {
				return nil, nil, fmt.Errorf("error parsing line: %v (val1: %v, val2: %v)", line, err1, err2)
			}
		} else {
			return nil, nil, fmt.Errorf("invalid line format (not enough columns): %v", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return leftNumbers, rightNumbers, nil
}

func main() {
	leftNumbers, rightNumbers, err := readNumbersToArrays("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	counts := make(map[int]int)
	for rnIndex := range rightNumbers {
		counts[rightNumbers[rnIndex]]++
	}

	var totalScore int
	for lnIndex := range leftNumbers {
		totalScore += leftNumbers[lnIndex] * counts[leftNumbers[lnIndex]]
	}
	fmt.Println("totalScore:", totalScore)
}
