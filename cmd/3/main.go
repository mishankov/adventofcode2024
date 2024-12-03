package main

import (
	"aoc2024/pkg/aocutils"
	"log"
	"regexp"
	"strings"
)

func main() {
	bytesData := aocutils.GetFileBytes("data/3")
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

			left := aocutils.ToInt(digitMatches[0])
			right := aocutils.ToInt(digitMatches[1])

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
