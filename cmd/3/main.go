package main

import (
	"aoc2024/internal/utils"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	bytesData := utils.GetFileBytes("data/3")
	stringData := string(bytesData)

	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	matches := re.FindAllString(stringData, -1)

	resultOne := 0
	resultTwo := 0
	doMul := true
	for _, match := range matches {
		switch {
		case strings.HasPrefix(match, "mul"):
			re := regexp.MustCompile(`\d+`)
			digitMatches := re.FindAllString(match, -1)

			left, err := strconv.Atoi(digitMatches[0])
			if err != nil {
				log.Fatal("Error converting left number to string:", err)
			}

			right, err := strconv.Atoi(digitMatches[1])
			if err != nil {
				log.Fatal("Error converting right number to string:", err)
			}

			resultOne += left * right

			if doMul {
				resultTwo += left * right
			}
		case match == "do()":
			doMul = true
		case match == "don't()":
			doMul = false
		}

	}

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
