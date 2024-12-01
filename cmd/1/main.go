package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	file, err := os.Open("data/1")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	bytesData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file content:", err)
	}

	byteLines := bytes.Split(bytesData, []byte{13, 10})

	left := make([]int, len(byteLines))
	right := make([]int, len(byteLines))
	rightAmounts := make(map[int]int)

	for index, line := range byteLines {
		numberBytes := bytes.Split(line, []byte{32, 32, 32})

		left[index], err = strconv.Atoi(string(numberBytes[0]))
		if err != nil {
			log.Fatal("Error converting left number:", err)
		}

		right[index], err = strconv.Atoi(string(numberBytes[1]))
		if err != nil {
			log.Fatal("Error converting right number:", err)
		}

		rightAmounts[right[index]] += 1
	}

	slices.Sort(left)
	slices.Sort(right)

	resultOne := 0
	resultTwo := 0
	for i := range left {
		resultOne += abs(left[i] - right[i])
		resultTwo += left[i] * rightAmounts[left[i]]
	}

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
