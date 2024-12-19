package main

import (
	"fmt"
	"os"
)

func readLine(filename string) (string, error) {
	buff, err := os.ReadFile(filename)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	return string(buff), nil
}

func appendRepeat(arr []int, val int, count int) []int {
	for i := 0; i < count; i++ {
		arr = append(arr, val)
	}
	return arr
}

func parseFileSystem(line string) []int {
	fileSystem := []int{}
	isFile := true
	id := 0
	for _, n := range line {
		count := int(n - '0')
		if isFile {
			fileSystem = appendRepeat(fileSystem, id, count)
			id++
		} else {
			fileSystem = appendRepeat(fileSystem, -1, count)
		}
		isFile = !isFile
	}
	return fileSystem
}

func formatFileSystem(fileSystem []int) []int {
	leftIndex := 0
	for fileSystem[leftIndex] != -1 {
		leftIndex++
	}
	rightIndex := len(fileSystem) - 1

	for leftIndex < rightIndex {
		fileSystem[leftIndex], fileSystem[rightIndex] = fileSystem[rightIndex], fileSystem[leftIndex]
		for fileSystem[leftIndex] != -1 {
			leftIndex++
		}
		rightIndex--
		for fileSystem[rightIndex] == -1 {
			rightIndex--
		}
	}
	return fileSystem
}

func calcChecksum(fileSystem []int) int {
	checksum := 0
	for i := 0; fileSystem[i] != -1; i++ {
		checksum += i * fileSystem[i]
	}
	return checksum
}

func main() {
	line, err := readLine("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	blocks := parseFileSystem(line)
	formattedBlocks := formatFileSystem(blocks)
	checksum := calcChecksum(formattedBlocks)
	fmt.Println(checksum)
}
