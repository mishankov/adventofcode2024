package main

import (
	"aoc2024/pkg/aocutils"
	"bytes"
	"log"
	"slices"
	"strconv"
)

func solve(data []byte) (int, int) {
	splitedData := bytes.Split(data, []byte{32})

	numbers := make([]int, len(splitedData))
	for i, d := range splitedData {
		numbers[i] = aocutils.ToInt(d)
	}

	for range 25 {
		index := 0
		for index < len(numbers) {
			switch {
			case numbers[index] == 0:
				numbers[index] = 1
				index++
			case len(strconv.Itoa(numbers[index]))%2 == 0:
				numberAsString := strconv.Itoa(numbers[index])
				middleIndex := len(numberAsString) / 2

				left := numberAsString[:middleIndex]
				right := numberAsString[middleIndex:]

				numbers[index] = aocutils.ToInt(left)
				numbers = slices.Insert(numbers, index+1, aocutils.ToInt(right))

				index++
				index++
			default:
				numbers[index] *= 2024
				index++
			}
		}
	}

	return len(numbers), 0
}

func main() {
	bytesData := aocutils.GetFileBytes("data/11")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
