package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func absInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
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
		values := strings.Fields(line)
		isSafe := true
		direction := 0 // -1 descending, 1 accending
		for i := 0; i < len(values)-1; i++ {
			curr, currErr := strconv.Atoi(values[i])
			next, nextErr := strconv.Atoi(values[i+1])
			if currErr != nil || nextErr != nil {
				return 0, fmt.Errorf("Error parsing line %v %v", currErr, nextErr)
			}
			diff := absInt(curr - next)
			if diff == 0 || diff > 3 {
				isSafe = false
				break
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
					isSafe = false
					break
				}
			}
		}
		if isSafe {
			safeReportsCounter++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return safeReportsCounter, nil
}

func main() {
	safeReportsCounter, err := calcSafeReports("test.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("safeReportsCounter:", safeReportsCounter)
}
