package main

import (
	"aoc2024/pkg/aocutils"
	"log"
	"strconv"
)

func generateVariants(levels [][]byte) [][][]byte {
	result := make([][][]byte, 0)
	for i := range len(levels) {
		subResult := make([][]byte, 0)
		for j, level := range levels {
			if i != j {
				subResult = append(subResult, level)
			}
		}
		result = append(result, subResult)
	}

	return result
}

func checkSafe(levels [][]byte) bool {
	if len(levels) < 1 {
		return false
	}

	previousNumber, err := strconv.Atoi(string(levels[0]))
	if err != nil {
		log.Fatal("Error converting number to string:", err)
	}
	dir := 0
	isSafeStrict := true
	for _, numberByte := range levels[1:] {
		number := aocutils.ToInt(numberByte)

		delta := number - previousNumber
		var currentDir int
		if delta < 0 {
			currentDir = -1
		} else {
			currentDir = 1
		}

		if !(aocutils.Abs(delta) >= 1 && aocutils.Abs(delta) <= 3) {
			isSafeStrict = false
			break
		}

		if dir == 0 {
			dir = currentDir
		} else if currentDir != dir {
			isSafeStrict = false
			break
		}

		previousNumber = number
	}

	return isSafeStrict
}
