package main

import (
	"aoc2024/pkg/aocutils"
	"bytes"
	"log"
)

func solve(data []byte) (int, int) {
	byteLines := bytes.Split(data, []byte{13, 10})

	resultOne := 0
	resultTwo := 0
	for _, line := range byteLines {
		numberBytes := bytes.Split(line, []byte{32})

		isSafeStrict := checkSafe(numberBytes)

		if isSafeStrict {
			resultOne++
			resultTwo++
		} else {
			variants := generateVariants(numberBytes)
			for _, variant := range variants {
				if checkSafe(variant) {
					resultTwo++
					break
				}
			}
		}
	}

	return resultOne, resultTwo
}

func main() {
	bytesData := aocutils.GetFileBytes("data/2")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
