package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string) (map[int][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	dataMap := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ":")
		leftNumber, _ := strconv.Atoi(parts[0])
		strNums := strings.Split(strings.TrimSpace(parts[1]), " ")
		var rightNumbers []int
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			rightNumbers = append(rightNumbers, num)
		}
		dataMap[leftNumber] = append(dataMap[leftNumber], rightNumbers...)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	return dataMap, nil
}

func evaluateExpression(nums []int, ops []string) int {
	result := nums[0]
	for i, op := range ops {
		if op == "+" {
			result += nums[i+1]
		} else if op == "*" {
			result *= nums[i+1]
		}
	}
	return result
}

func generateResults(nums []int, operators []string, currentOps []string, index int, results *[]int) {
	if index == len(nums)-1 {
		*results = append(*results, evaluateExpression(nums, currentOps))
		return
	}

	for _, op := range operators {
		currentOps = append(currentOps, op)
		generateResults(nums, operators, currentOps, index+1, results)
		currentOps = currentOps[:len(currentOps)-1]
	}
}

func isResultMatch(leftNumber int, rightNumbers []int) bool {
	operators := []string{"+", "*"}

	var results []int
	generateResults(rightNumbers, operators, []string{}, 0, &results)

	for _, res := range results {
		if res == leftNumber {
			return true
		}
	}
	return false
}

func main() {
	dataMap, err := readFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var sumOfMarchedKeys int
	for key, value := range dataMap {
		if isResultMatch(key, value) {
			sumOfMarchedKeys += key
			fmt.Println(key, " => ", value)
		}
	}
	fmt.Println("Sum of matched keys", sumOfMarchedKeys)
}
