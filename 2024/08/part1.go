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

type point struct {
	x int
	y int
}

func findUniqueFrequencyPoints(grid [][]rune) map[rune][]point {
	uniqueFreq := map[rune][]point{}
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '.' {
				uniqueFreq[grid[i][j]] = append(uniqueFreq[grid[i][j]], point{x: i, y: j})
			}
		}
	}
	return uniqueFreq
}

func calcAnthenaPoint(grid [][]rune, aMain point, a point) (anthenaPoint point, canPlace bool) {
	n := len(grid)
	anthenaPoint = point{x: aMain.x + (aMain.x - a.x), y: aMain.y + (aMain.y - a.y)}
	canPlace = (anthenaPoint.x >= 0 && anthenaPoint.x < n) && (anthenaPoint.y >= 0 && anthenaPoint.y < n)
	return
}

func calcTotalAnthenas(grid [][]rune, uniqueFreq map[rune][]point) int {
	var anthenaPoints = map[point]bool{}
	for _, points := range uniqueFreq {
		pSize := len(points)
		for i := 0; i < pSize; i++ {
			for j := 0; j < pSize; j++ {
				if points[i] != points[j] {
					point, canPlace := calcAnthenaPoint(grid, points[i], points[j])
					if canPlace {
						anthenaPoints[point] = true
					}
				}
			}
		}
	}
	return len(anthenaPoints)
}

func main() {
	grid, err := readFileIntoRuneArray("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	uniqueFreq := findUniqueFrequencyPoints(grid)
	totalAnthenas := calcTotalAnthenas(grid, uniqueFreq)
	fmt.Println("total anthenas: ", totalAnthenas)
}
