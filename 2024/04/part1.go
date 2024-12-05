package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFileIntoRuneArray(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var text [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		text = append(text, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return text, nil
}

func calcHorizontalXmasCount(text [][]rune) int {
	var xMasCount int
	for i := 0; i < len(text); i++ {
		for j := 0; j < len(text[i])-3; j++ {
			if (text[i][j] == 'X' && text[i][j+1] == 'M' && text[i][j+2] == 'A' && text[i][j+3] == 'S') || (text[i][j] == 'S' && text[i][j+1] == 'A' && text[i][j+2] == 'M' && text[i][j+3] == 'X') {
				xMasCount++
			}
		}
	}
	return xMasCount
}

func calcVerticalXmasCount(text [][]rune) int {
	var xMasCount int
	for i := 0; i < len(text)-3; i++ {
		for j := 0; j < len(text[i]); j++ {
			if (text[i][j] == 'X' && text[i+1][j] == 'M' && text[i+2][j] == 'A' && text[i+3][j] == 'S') || (text[i][j] == 'S' && text[i+1][j] == 'A' && text[i+2][j] == 'M' && text[i+3][j] == 'X') {
				xMasCount++
			}
		}
	}
	return xMasCount
}

func calcDiagonalXmasCount(text [][]rune) int {
	var xMasCount int
	for i := 0; i < len(text)-3; i++ {
		for j := 0; j < len(text[i])-3; j++ {
			if (text[i][j] == 'X' && text[i+1][j+1] == 'M' && text[i+2][j+2] == 'A' && text[i+3][j+3] == 'S') || (text[i][j] == 'S' && text[i+1][j+1] == 'A' && text[i+2][j+2] == 'M' && text[i+3][j+3] == 'X') {
				xMasCount++
			}
		}
	}
	return xMasCount
}

func calcReversedDiagonalXmasCount(text [][]rune) int {
	var xMasCount int
	for i := 0; i < len(text)-3; i++ {
		for j := 3; j < len(text[i]); j++ {
			if (text[i][j] == 'X' && text[i+1][j-1] == 'M' && text[i+2][j-2] == 'A' && text[i+3][j-3] == 'S') || (text[i][j] == 'S' && text[i+1][j-1] == 'A' && text[i+2][j-2] == 'M' && text[i+3][j-3] == 'X') {
				xMasCount++
			}
		}
	}
	return xMasCount
}

func main() {
	text, err := readFileIntoRuneArray("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalCount := calcHorizontalXmasCount(text)
	totalCount += calcVerticalXmasCount(text)
	totalCount += calcDiagonalXmasCount(text)
	totalCount += calcReversedDiagonalXmasCount(text)
	fmt.Println("XMAS occurs", totalCount)
}
