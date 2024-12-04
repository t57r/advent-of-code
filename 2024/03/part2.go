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

func getTotalOfMuls2(line string, isMulEnabled bool) (int, bool) {
	regex := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don't\(\))`)
	matches := regex.FindAllStringSubmatch(line, -1)
	var total int
	for _, match := range matches {
		if match[1] != "" {
			if isMulEnabled {
				num1, _ := strconv.Atoi(match[2])
				num2, _ := strconv.Atoi(match[3])
				total += num1 * num2
			}
		} else if match[4] != "" {
			// do() matches
			isMulEnabled = true
		} else if match[5] != "" {
			// don't() matches
			isMulEnabled = false
		}
	}
	return total, isMulEnabled
}

func main() {
	lines, err := readLinesToArray("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var total int
	isMulEnabled := true
	for _, line := range lines {
		curTotal, curMulEnabled := getTotalOfMuls2(line, isMulEnabled)
		total += curTotal
		isMulEnabled = curMulEnabled
	}
	fmt.Println("Total:", total)
}
