package main

import (
	"aoc2024/internal/utils"
	"bytes"
	"log"
	"strconv"
)

func main() {
	bytesData := utils.GetFileBytes("data/2")
	byteLines := bytes.Split(bytesData, []byte{13, 10})

	resultOne := 0
	resultTwo := 0
	for _, line := range byteLines {
		numberBytes := bytes.Split(line, []byte{32})

		previousNumber, err := strconv.Atoi(string(numberBytes[0]))
		if err != nil {
			log.Fatal("Error converting number to string:", err)
		}
		dir := 0
		isSafeStrict := true
		for _, numberByte := range numberBytes[1:] {
			number, err := strconv.Atoi(string(numberByte))
			if err != nil {
				log.Fatal("Error converting number to string:", err)
			}

			delta := number - previousNumber
			var currentDir int
			if delta < 0 {
				currentDir = -1
			} else {
				currentDir = 1
			}

			if !(utils.Abs(delta) >= 1 && utils.Abs(delta) <= 3) {
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

		if isSafeStrict {
			resultOne++
		}

	}

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
