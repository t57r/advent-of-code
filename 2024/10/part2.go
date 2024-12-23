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

type Coordinate struct {
	x, y int
}

var directions = []Coordinate{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

func isValidMove(matrix [][]rune, current, next Coordinate) bool {
	rows, cols := len(matrix), len(matrix[0])
	if next.x < 0 || next.x >= rows || next.y < 0 || next.y >= cols {
		return false // Out of bounds
	}
	if matrix[next.x][next.y] == '.' {
		return false // Unreachable
	}
	currValue := int(matrix[current.x][current.y] - '0')
	nextValue := int(matrix[next.x][next.y] - '0')
	if nextValue != currValue+1 {
		return false // Too high
	}
	return true
}

func findRoutes(matrix [][]rune, current Coordinate) int {
	if matrix[current.x][current.y] == '9' {
		return 1
	}

	routes := 0
	for _, dir := range directions {
		next := Coordinate{current.x + dir.x, current.y + dir.y}
		if isValidMove(matrix, current, next) {
			routes += findRoutes(matrix, next)
		}
	}

	return routes
}

func findAllPoints(matrix [][]rune, pointValue rune) []Coordinate {
	var points []Coordinate
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == pointValue {
				points = append(points, Coordinate{x: i, y: j})
			}
		}
	}
	return points
}

func countRoutes(matrix [][]rune) int {
	startingPoints := findAllPoints(matrix, '0')
	totalRoutes := 0
	for _, start := range startingPoints {
		totalRoutes += findRoutes(matrix, start)
	}
	return totalRoutes
}

func main() {
	matrix, err := readFileIntoRuneArray("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalRoutes := countRoutes(matrix)
	fmt.Println(totalRoutes)
}
