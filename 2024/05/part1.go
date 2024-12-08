package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(filename string) (map[int][]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Variables to store results
	dataMap := make(map[int][]int)
	var twoDimArray [][]int

	// Read file line by line
	scanner := bufio.NewScanner(file)
	isSecondPart := false

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			// Switch to parsing the second part after the empty line
			isSecondPart = true
			continue
		}

		if !isSecondPart {
			// First part: Parse the map
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				fmt.Println("Invalid line in map section:", line)
				continue
			}

			// Parse key and value
			key, err1 := strconv.Atoi(parts[0])
			value, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing integers in map section:", line)
				continue
			}

			// Append value to the map
			dataMap[key] = append(dataMap[key], value)
		} else {
			// Second part: Parse the 2D array
			strNums := strings.Split(line, ",")
			var row []int
			for _, strNum := range strNums {
				num, err := strconv.Atoi(strings.TrimSpace(strNum))
				if err != nil {
					fmt.Println("Error parsing integer in 2D array section:", line)
					continue
				}
				row = append(row, num)
			}
			twoDimArray = append(twoDimArray, row)
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("Error reading file:", err)
	}

	return dataMap, twoDimArray, nil
}

func main() {
	pageOrderingMap, pageNumbers, err := readFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var correctPages [][]int
	for i := 0; i < len(pageNumbers); i++ {
		allPagesLower := true
		for j := 0; j < len(pageNumbers[i])-1; j++ {
			key := pageNumbers[i][j]
			keyPages := pageOrderingMap[key]
			for z := j + 1; z < len(pageNumbers[i]); z++ {
				if !slices.Contains(keyPages, pageNumbers[i][z]) {
					allPagesLower = false
					break
				}
			}
			if !allPagesLower {
				break
			}
		}
		if allPagesLower {
			correctPages = append(correctPages, pageNumbers[i])
		}
	}

	var sumOfMiddles int
	for _, row := range correctPages {
		sumOfMiddles += row[len(row)/2]
	}
	fmt.Println("Sum of middles ", sumOfMiddles)
}
