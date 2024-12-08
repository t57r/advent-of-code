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

	dataMap := make(map[int][]int)
	var twoDimArray [][]int
	scanner := bufio.NewScanner(file)
	isSecondPart := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			isSecondPart = true
			continue
		}

		if !isSecondPart {
			parts := strings.Split(line, "|")
			if len(parts) != 2 {
				fmt.Println("Invalid line in map section:", line)
				continue
			}
			key, err1 := strconv.Atoi(parts[0])
			value, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Error parsing integers in map section:", line)
				continue
			}
			dataMap[key] = append(dataMap[key], value)
		} else {
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
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("Error reading file:", err)
	}
	return dataMap, twoDimArray, nil
}

func calcIncorrectPages(pageOrderingMap map[int][]int, pageNumbers [][]int) [][]int {
	var incorrectPages [][]int
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
		if !allPagesLower {
			incorrectPages = append(incorrectPages, pageNumbers[i])
		}
	}
	return incorrectPages
}

func contains(slice []int, x int) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == x {
			return true
		}
	}
	return false
}

func sortIncorrectPages(incorrectPages [][]int, pageOrderingMap map[int][]int) [][]int {
	var sortedPages [][]int
	for _, incorrectPage := range incorrectPages {
		slices.SortFunc(incorrectPage, func(a, b int) int {
			aRights, aRightsExist := pageOrderingMap[a]
			bRights, bRightsExist := pageOrderingMap[b]

			if (aRightsExist && !bRightsExist) || contains(aRights, b) {
				return -1
			}
			if (!aRightsExist && bRightsExist) || contains(bRights, a) {
				return 1
			}
			return 0
		})
		sortedPages = append(sortedPages, incorrectPage)
	}
	return sortedPages
}

func main() {
	pageOrderingMap, pageNumbers, err := readFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	incorrectPages := calcIncorrectPages(pageOrderingMap, pageNumbers)
	sortedPages := sortIncorrectPages(incorrectPages, pageOrderingMap)

	var sumOfMiddles int
	for _, row := range sortedPages {
		sumOfMiddles += row[len(row)/2]
	}
	fmt.Println("Sum of middles ", sumOfMiddles)
}
