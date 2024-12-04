package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readLinesToArray(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return lines, nil
}

func getTotalOfMuls(line string) int {
	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := regex.FindAllStringSubmatch(line, -1)
	var total int
	for _, match := range matches {
		if len(match) == 3 {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			total += num1 * num2
		}
	}
	return total
}

func main() {
	lines, err := readLinesToArray("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var total int
	for _, line := range lines {
		total += getTotalOfMuls(line)
	}
	fmt.Println("Total:", total)
}
