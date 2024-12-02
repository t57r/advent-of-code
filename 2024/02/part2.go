package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToIntSlice(slice []string) ([]int, error) {
	result := make([]int, 0, len(slice))
	for _, value := range slice {
		converted, err := strconv.Atoi(value)
		if err != nil {
			return result, fmt.Errorf("Error parsing line %v", err)
		}
		result = append(result, converted)
	}
	return result, nil
}

func absInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func isSafe(slice []int) bool {
	direction := 0 // -1 descending, 1 accending
	for i := 0; i < len(slice)-1; i++ {
		curr := slice[i]
		next := slice[i+1]
		diff := absInt(curr - next)
		if diff == 0 || diff > 3 {
			return false
		}
		stepDirection := 0
		if next > curr {
			stepDirection = 1
		} else {
			stepDirection = -1
		}
		if direction == 0 {
			// setting the direction for the first time
			direction = stepDirection
		} else {
			// check if the direction is unchanged
			if stepDirection != direction {
				return false
			}
		}
	}
	return true
}

func removeOneItem(slice []int, indexToRemove int) []int {
	result := make([]int, 0, len(slice)-1)
	for index, value := range slice {
		if index == indexToRemove {
			continue
		}
		result = append(result, value)
	}
	return result
}

func isSafeWithOneError(slice []int) bool {
	for i := 0; i < len(slice); i++ {
		if isSafe(removeOneItem(slice, i)) {
			return true
		}
	}
	return false
}

func calcSafeReports(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var safeReportsCounter int
	for scanner.Scan() {
		line := scanner.Text()
		values, err := convertToIntSlice(strings.Fields(line))
		if err != nil {
			return 0, fmt.Errorf("Error parsing line %v", err)
		}
		if isSafeWithOneError(values) {
			safeReportsCounter++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return safeReportsCounter, nil
}

func main() {
	safeReportsCounter, err := calcSafeReports("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("safeReportsCounter:", safeReportsCounter)
}
