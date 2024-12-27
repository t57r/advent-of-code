package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) ([]int, error) {
	buff, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	content := string(buff)
	tokens := strings.Fields(content)
	intArr := []int{}
	for _, token := range tokens {
		tokenNumb, _ := strconv.Atoi(token)
		intArr = append(intArr, tokenNumb)
	}
	return intArr, nil
}

func blink(stones []int) []int {
	newStones := []int{}
	for _, stone := range stones {
		if stone == 0 {
			newStones = append(newStones, 1)
			continue
		}

		stoneStr := strconv.Itoa(stone)
		stoneDigitLength := len(stoneStr)
		if stoneDigitLength%2 == 0 {
			left, _ := strconv.Atoi(stoneStr[:stoneDigitLength/2])
			right, _ := strconv.Atoi(stoneStr[stoneDigitLength/2:])
			newStones = append(newStones, left, right)
			continue
		}

		newStones = append(newStones, stone*2024)
	}
	return newStones
}

func blinkTimes(stones []int, times int) int {
	blinked := stones
	for i := 0; i < times; i++ {
		blinked = blink(blinked)
	}
	return len(blinked)
}

func main() {
	intArray, err := readFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	stonesNum := blinkTimes(intArray, 75)
	fmt.Println(stonesNum)
}
