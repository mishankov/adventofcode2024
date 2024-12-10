package main

import (
	"aoc2024/pkg/aocutils"
	"log"
	"strconv"
	"strings"
)

func canMatch(numbers []int, currentResult uint64, target uint64) bool {
	if currentResult == 0 {
		return canMatch(numbers[1:], uint64(numbers[0]), target)
	}

	if len(numbers) == 1 {
		return canMatch([]int{}, currentResult+uint64(numbers[0]), target) || canMatch([]int{}, currentResult*uint64(numbers[0]), target)
	}

	if len(numbers) == 0 {
		return currentResult == target
	}

	return canMatch(numbers[1:], currentResult+uint64(numbers[0]), target) || canMatch(numbers[1:], currentResult*uint64(numbers[0]), target)
}

func concat(first, second uint64) uint64 {
	res, err := strconv.ParseUint(strconv.FormatUint(first, 10)+strconv.FormatUint(second, 10), 10, 64)
	if err != nil {
		log.Fatal("Error concating:", err)
	}

	return res
}

func canMatch2(numbers []int, currentResult uint64, target uint64) bool {
	if currentResult == 0 {
		return canMatch2(numbers[1:], uint64(numbers[0]), target)
	}

	if len(numbers) == 1 {
		return canMatch2([]int{}, currentResult+uint64(numbers[0]), target) || canMatch2([]int{}, currentResult*uint64(numbers[0]), target) || canMatch2([]int{}, concat(currentResult, uint64(numbers[0])), target)
	}

	if len(numbers) == 0 {
		return currentResult == target
	}

	return canMatch2(numbers[1:], currentResult+uint64(numbers[0]), target) || canMatch2(numbers[1:], currentResult*uint64(numbers[0]), target) || canMatch2(numbers[1:], concat(currentResult, uint64(numbers[0])), target)
}

func solve(data []byte) (uint64, uint64) {
	lines := aocutils.SplitByteLines(data)

	var r1 uint64 = 0
	var r2 uint64 = 0
	for _, line := range lines {
		lineStr := string(line)
		target := aocutils.ToInt(strings.Split(lineStr, ":")[0])

		numbers := []int{}
		for _, numStr := range strings.Split(strings.Split(lineStr, ":")[1], " ") {
			if numStr != "" {
				numbers = append(numbers, aocutils.ToInt(numStr))
			}
		}

		if canMatch(numbers, 0, uint64(target)) {
			r1 += uint64(target)
		}

		if canMatch2(numbers, 0, uint64(target)) {
			r2 += uint64(target)
		}
	}

	return r1, r2
}

func main() {
	bytesData := aocutils.GetFileBytes("data/7")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
