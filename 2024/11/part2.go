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

func blink(stones map[int]int) map[int]int {
	newStones := map[int]int{}
	for stone, _ := range stones {
		if stone == 0 {
			newStones[1] += stones[0]
			continue
		}

		stoneStr := strconv.Itoa(stone)
		stoneDigitLength := len(stoneStr)
		if stoneDigitLength%2 == 0 {
			left, _ := strconv.Atoi(stoneStr[:stoneDigitLength/2])
			right, _ := strconv.Atoi(stoneStr[stoneDigitLength/2:])
			newStones[left] += stones[stone]
			newStones[right] += stones[stone]
			continue
		}

		newStones[stone*2024] += stones[stone]
	}
	return newStones
}

func blinkTimes(stones map[int]int, times int) int {
	blinked := stones
	for i := 0; i < times; i++ {
		blinked = blink(blinked)
	}
	count := 0
	for _, v := range blinked {
		count += v
	}
	return count
}

func main() {
	intArray, err := readFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	stones := map[int]int{}
	for _, n := range intArray {
		stones[n] += 1
	}

	stonesNum := blinkTimes(stones, 75)
	fmt.Println(stonesNum)
}
