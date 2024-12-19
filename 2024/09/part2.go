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
	rightIndex := len(fileSystem) - 1
	for {
		// calc latest file size
		fileID := fileSystem[rightIndex]
		fileSize := 1
		for rightIndex > 0 && fileSystem[rightIndex] == fileSystem[rightIndex-1] {
			rightIndex--
			fileSize++
		}

		// seek for the free space that could fit fileSize
		freeSpaceFound := false
		freeSpaceIndex := -1
		freeSpaceSize := 0
		for leftIndex := 0; leftIndex <= rightIndex; leftIndex++ {
			if fileSystem[leftIndex] == -1 {
				if freeSpaceSize == 0 {
					freeSpaceIndex = leftIndex
				}
				freeSpaceSize++
			} else {
				if freeSpaceSize >= fileSize {
					freeSpaceFound = true
					break
				}
				freeSpaceSize = 0
				freeSpaceIndex = -1
			}
		}

		// free space found, swap file with the free space
		if freeSpaceFound {
			for i := 0; i < fileSize; i++ {
				fileSystem[freeSpaceIndex+i], fileSystem[rightIndex+i] = fileSystem[rightIndex+i], fileSystem[freeSpaceIndex+i]
			}
		}

		// iterate to the next file, skip empty spaces
		nextFileID := fileID - 1
		if nextFileID == 0 {
			// we have already traversed all file system
			break
		}
		for fileSystem[rightIndex] != nextFileID {
			rightIndex--
		}

	}
	return fileSystem
}

func calcChecksum(fileSystem []int) int {
	checksum := 0
	for i := 0; i < len(fileSystem); i++ {
		if fileSystem[i] != -1 {
			checksum += i * fileSystem[i]
		}
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
