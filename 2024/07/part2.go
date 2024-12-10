package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type numbers struct {
	left  int64
	right []int64
}

func readFile(filename string) ([]numbers, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var dataMap []numbers
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Split(line, ":")
		leftNumber, _ := strconv.ParseInt(parts[0], 10, 64)
		strNums := strings.Split(strings.TrimSpace(parts[1]), " ")
		var rightNumbers []int64
		for _, strNum := range strNums {
			num, _ := strconv.ParseInt(strNum, 10, 64)
			rightNumbers = append(rightNumbers, num)
		}
		dataMap = append(dataMap, numbers{
			left:  leftNumber,
			right: rightNumbers,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("Error reading file: %v", err)
	}

	return dataMap, nil
}

func evaluateExpression(nums []int64, ops []string) int64 {
	result := nums[0]
	for i, op := range ops {
		if op == "+" {
			result += nums[i+1]
		} else if op == "*" {
			result *= nums[i+1]
		} else if op == "|" {
			lh := strconv.FormatInt(result, 10)
			rh := strconv.FormatInt(nums[i+1], 10)
			result, _ = strconv.ParseInt(lh+rh, 10, 64)
		}
	}
	return result
}

func generateResults(nums []int64, operators []string, currentOps []string, index int, results *[]int64) {
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

func isResultMatch(leftNumber int64, rightNumbers []int64) bool {
	operators := []string{"+", "*", "|"}

	var results []int64
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

	var sumOfMatchedKeys int64
	for _, numbers := range dataMap {
		if isResultMatch(numbers.left, numbers.right) {
			sumOfMatchedKeys += numbers.left
		}
	}
	fmt.Println(sumOfMatchedKeys)
}
