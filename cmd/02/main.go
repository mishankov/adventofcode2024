package main

import (
	"aoc2024/pkg/aocutils"
	"bytes"
	"log"
	"sync"
	"sync/atomic"
)

func solveAsync(data []byte) (int, int) {
	byteLines := aocutils.SplitByteLines(data)

	var wg sync.WaitGroup

	var resultOne atomic.Uint32
	var resultTwo atomic.Uint32
	for _, line := range byteLines {
		numberBytes := bytes.Split(line, []byte{32})

		wg.Add(1)

		go func() {
			defer wg.Done()

			isSafeStrict := checkSafe(numberBytes)

			if isSafeStrict {
				resultOne.Add(1)
				resultTwo.Add(1)
			} else {
				variants := generateVariants(numberBytes)
				for _, variant := range variants {
					if checkSafe(variant) {
						resultTwo.Add(1)
						break
					}
				}
			}
		}()
	}

	wg.Wait()

	return int(resultOne.Load()), int(resultTwo.Load())
}

func solve(data []byte) (int, int) {
	byteLines := aocutils.SplitByteLines(data)

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
