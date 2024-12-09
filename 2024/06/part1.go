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

const (
	MOVE_UP    = iota
	MOVE_RIGHT = iota
	MOVE_DOWN  = iota
	MOVE_LEFT  = iota
)

type Coordinate struct {
	x int
	y int
}

func findStartingPoint(puzzleMap [][]rune) (c Coordinate, err error) {
	n := len(puzzleMap)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if puzzleMap[i][j] == '^' {
				return Coordinate{x: i, y: j}, nil
			}
		}
	}
	return Coordinate{}, fmt.Errorf("Starting point not found")
}

type PuzzleMap struct {
	gameMap   [][]rune // assume the map is square (size x size)
	size      int
	cursor    Coordinate
	direction int
	seen      map[Coordinate]bool
}

func (pm *PuzzleMap) SwitchDirection() {
	var nextDirection int
	switch pm.direction {
	case MOVE_UP:
		nextDirection = MOVE_RIGHT
	case MOVE_RIGHT:
		nextDirection = MOVE_DOWN
	case MOVE_DOWN:
		nextDirection = MOVE_LEFT
	case MOVE_LEFT:
		nextDirection = MOVE_UP
	}
	pm.direction = nextDirection
}

func (pm *PuzzleMap) MakeNextMove() (isFinished, isObstacle bool) {
	var deltaX, deltaY int
	switch pm.direction {
	case MOVE_UP:
		deltaX, deltaY = -1, 0
	case MOVE_RIGHT:
		deltaX, deltaY = 0, 1
	case MOVE_DOWN:
		deltaX, deltaY = 1, 0
	case MOVE_LEFT:
		deltaX, deltaY = 0, -1
	}

	nextX, nextY := pm.cursor.x+deltaX, pm.cursor.y+deltaY
	if nextX < 0 || nextX >= pm.size || nextY < 0 || nextY >= pm.size {
		isFinished = true
		return
	}

	if pm.gameMap[nextX][nextY] == '#' {
		isObstacle = true
		return
	}

	if pm.gameMap[nextX][nextY] != 'X' {
		pm.gameMap[nextX][nextY] = 'X'
		pm.seen[Coordinate{x: nextX, y: nextY}] = true
	}

	pm.cursor.x = nextX
	pm.cursor.y = nextY

	return false, false
}

func main() {
	gameMap, err := readFileIntoRuneArray("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	startingPoint, err := findStartingPoint(gameMap)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	gameMap[startingPoint.x][startingPoint.y] = 'X'

	puzzleMap := PuzzleMap{
		gameMap:   gameMap,
		size:      len(gameMap),
		cursor:    startingPoint,
		direction: MOVE_UP,
		seen:      map[Coordinate]bool{startingPoint: true},
	}

	for {
		isFinished, isObstacle := puzzleMap.MakeNextMove()
		if isFinished {
			break
		}
		if isObstacle {
			puzzleMap.SwitchDirection()
		}
	}

	fmt.Println("Filled moves:", len(puzzleMap.seen))

}
