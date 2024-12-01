package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readNumbersToHeaps(filename string) (*IntHeap, *IntHeap, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	leftHeap := &IntHeap{}
	rightHeap := &IntHeap{}
	heap.Init(leftHeap)
	heap.Init(rightHeap)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		if len(values) >= 2 {
			val1, err1 := strconv.Atoi(values[0])
			val2, err2 := strconv.Atoi(values[1])

			if err1 == nil && err2 == nil {
				heap.Push(leftHeap, val1)
				heap.Push(rightHeap, val2)
			} else {
				return nil, nil, fmt.Errorf("error parsing line: %v (val1: %v, val2: %v)", line, err1, err2)
			}
		} else {
			return nil, nil, fmt.Errorf("invalid line format (not enough columns): %v", line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	return leftHeap, rightHeap, nil
}

func main() {
	leftHeap, rightHeap, err := readNumbersToHeaps("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var totalDistance int
	for leftHeap.Len() > 0 && rightHeap.Len() > 0 {
		leftValue := heap.Pop(leftHeap).(int)
		rightValue := heap.Pop(rightHeap).(int)
		distance := absInt(leftValue - rightValue)
		totalDistance += distance
	}
	fmt.Println("Total distance:", totalDistance)
}
