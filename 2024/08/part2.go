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

func placeAnthenas(grid [][]rune, anthenaPoints map[point]bool, aMain point, a point) {
	delta := point{x: aMain.x - a.x, y: aMain.y - a.y}
	placeAnthenasWithDelta(grid, anthenaPoints, delta, aMain, a)

	delta = point{x: a.x - aMain.x, y: a.y - aMain.y}
	placeAnthenasWithDelta(grid, anthenaPoints, delta, aMain, a)
}

func placeAnthenasWithDelta(grid [][]rune, anthenaPoints map[point]bool, delta point, aMain point, a point) {
	n := len(grid)
	anthenaPoint := point{x: aMain.x + delta.x, y: aMain.y + delta.y}
	for {
		canPlace := (anthenaPoint.x >= 0 && anthenaPoint.x < n) && (anthenaPoint.y >= 0 && anthenaPoint.y < n)
		if canPlace {
			anthenaPoints[anthenaPoint] = true
			anthenaPoint = point{x: anthenaPoint.x + delta.x, y: anthenaPoint.y + delta.y}
		} else {
			break
		}
	}
}

func calcTotalAnthenas(grid [][]rune, uniqueFreq map[rune][]point) int {
	var anthenaPoints = map[point]bool{}
	for _, points := range uniqueFreq {
		pSize := len(points)
		for i := 0; i < pSize; i++ {
			for j := 0; j < pSize; j++ {
				if points[i] != points[j] {
					placeAnthenas(grid, anthenaPoints, points[i], points[j])
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
