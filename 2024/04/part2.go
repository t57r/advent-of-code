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

func calcXmasCount(text [][]rune) int {
	var xMasCount int
	for i := 0; i < len(text)-2; i++ {
		for j := 0; j < len(text[i])-2; j++ {
			diagonalMatched := (text[i][j] == 'M' && text[i+1][j+1] == 'A' && text[i+2][j+2] == 'S') || (text[i][j] == 'S' && text[i+1][j+1] == 'A' && text[i+2][j+2] == 'M')
			reversedDiagonalMatched := (text[i][j+2] == 'M' && text[i+1][j+1] == 'A' && text[i+2][j] == 'S') || (text[i][j+2] == 'S' && text[i+1][j+1] == 'A' && text[i+2][j] == 'M')
			if diagonalMatched && reversedDiagonalMatched {
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

	fmt.Println("XMAS occurs", calcXmasCount(text))
}
