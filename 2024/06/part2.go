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
	MOVE_UP    = 0
	MOVE_RIGHT = 1
	MOVE_DOWN  = 2
	MOVE_LEFT  = 3
)

type Coordinate struct {
	x int
	y int
}

type CoordinateWithDirection struct {
	Coordinate
	direction int
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

func (pm *PuzzleMap) GetDirectionDelta() (int, int) {
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
	return deltaX, deltaY
}

func (pm *PuzzleMap) MakeNextMove() (isFinished, isObstacle bool) {
	deltaX, deltaY := pm.GetDirectionDelta()
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

func createPuzzleMap(gameMap [][]rune, startingPoint Coordinate) PuzzleMap {
	return PuzzleMap{
		gameMap:   gameMap,
		size:      len(gameMap),
		cursor:    startingPoint,
		direction: MOVE_UP,
		seen:      map[Coordinate]bool{startingPoint: true},
	}
}

func calcSeenCoordinates(gameMap [][]rune, startingPoint Coordinate) map[Coordinate]bool {
	puzzleMap := createPuzzleMap(gameMap, startingPoint)
	for {
		isFinished, isObstacle := puzzleMap.MakeNextMove()
		if isFinished {
			break
		}
		if isObstacle {
			puzzleMap.SwitchDirection()
		}
	}
	return puzzleMap.seen
}

func deepCopyRuneArray(original [][]rune) [][]rune {
	copied := make([][]rune, len(original))
	for i := range original {
		copied[i] = make([]rune, len(original[i]))
		copy(copied[i], original[i])
	}
	return copied
}

func hasLoop(gameMap [][]rune, startingPoint Coordinate, obstaclePoint Coordinate) bool {
	if gameMap[obstaclePoint.x][obstaclePoint.y] == '#' {
		return false
	}
	gameMap[obstaclePoint.x][obstaclePoint.y] = '#'

	puzzleMap := createPuzzleMap(gameMap, startingPoint)
	seen := map[CoordinateWithDirection]bool{}
	for {
		currPos := CoordinateWithDirection{
			Coordinate: puzzleMap.cursor,
			direction:  puzzleMap.direction,
		}
		_, currPosSeen := seen[currPos]
		if currPosSeen {
			return true
		}
		seen[currPos] = true

		isFinished, isObstacle := puzzleMap.MakeNextMove()
		if isFinished {
			break
		}
		if isObstacle {
			puzzleMap.SwitchDirection()
		}
	}
	return false
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

	seenCoordinates := calcSeenCoordinates(gameMap, startingPoint)
	fmt.Println("Filled moves:", len(seenCoordinates))

	var loopCount int
	for key, _ := range seenCoordinates {
		if key.x == startingPoint.x && key.y == startingPoint.y {
			continue
		}
		if hasLoop(deepCopyRuneArray(gameMap), startingPoint, key) {
			loopCount++
		}
	}
	fmt.Println("Loop count", loopCount)
}
