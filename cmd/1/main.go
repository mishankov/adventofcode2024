package main

import (
	"aoc2024/pkg/aocutils"
	"bytes"
	"log"
	"slices"
)

func main() {
	bytesData := aocutils.GetFileBytes("data/1")
	byteLines := bytes.Split(bytesData, []byte{13, 10})

	left := make([]int, len(byteLines))
	right := make([]int, len(byteLines))
	rightAmounts := make(map[int]int)

	for index, line := range byteLines {
		numberBytes := bytes.Split(line, []byte{32, 32, 32})

		left[index] = aocutils.ToInt(numberBytes[0])
		right[index] = aocutils.ToInt(numberBytes[1])

		rightAmounts[right[index]] += 1
	}

	slices.Sort(left)
	slices.Sort(right)

	resultOne := 0
	resultTwo := 0
	for i := range left {
		resultOne += aocutils.Abs(left[i] - right[i])
		resultTwo += left[i] * rightAmounts[left[i]]
	}

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
