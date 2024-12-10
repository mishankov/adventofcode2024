package main

import (
	"aoc2024/pkg/aocutils"
	"log"
)

type position struct {
	x int
	y int
}

func (p position) isValid(maxX, maxY int) bool {
	return p.x <= maxX && p.y <= maxY && p.x >= 0 && p.y >= 0
}

type direction struct {
	x int
	y int
}

func findLetter(field [][]rune, letter rune, currentPos position, currentDir direction) []direction {
	maxY := len(field) - 1
	maxX := len(field[0]) - 1

	if letter == 'M' {
		allDirs := []direction{
			{-1, -1},
			{-1, 0},
			{-1, 1},
			{0, -1},
			{0, 0},
			{0, 1},
			{1, -1},
			{1, 0},
			{1, 1},
		}

		results := []direction{}
		for _, dir := range allDirs {
			newPos := position{currentPos.x + dir.x, currentPos.y + dir.y}
			if newPos.isValid(maxX, maxY) && field[newPos.y][newPos.x] == letter {
				results = append(results, dir)
			}
		}

		return results
	}

	newPos := position{currentPos.x + currentDir.x, currentPos.y + currentDir.y}
	if newPos.isValid(maxX, maxY) && field[newPos.y][newPos.x] == letter {
		return []direction{currentDir}
	}

	return nil
}

func InMAS(field [][]rune, pos position) bool {
	maxY := len(field) - 1
	maxX := len(field[0]) - 1

	topLeftPos := position{pos.x - 1, pos.y - 1}
	topRightPos := position{pos.x + 1, pos.y - 1}
	bottomLeftPos := position{pos.x - 1, pos.y + 1}
	bottomRightPos := position{pos.x + 1, pos.y + 1}

	if !(topLeftPos.isValid(maxX, maxY) && topRightPos.isValid(maxX, maxY) && bottomLeftPos.isValid(maxX, maxY) && bottomRightPos.isValid(maxX, maxY)) {
		return false
	}

	topLeftLetter := field[topLeftPos.y][topLeftPos.x]
	topRightLetter := field[topRightPos.y][topRightPos.x]
	bottomLeftLetter := field[bottomLeftPos.y][bottomLeftPos.x]
	bottomRightLetter := field[bottomRightPos.y][bottomRightPos.x]

	if !(topLeftLetter == 'M' && bottomRightLetter == 'S' || topLeftLetter == 'S' && bottomRightLetter == 'M') {
		return false
	}

	if !(topRightLetter == 'M' && bottomLeftLetter == 'S' || topRightLetter == 'S' && bottomLeftLetter == 'M') {
		return false
	}

	return true
}

func solve(data []byte) (int, int) {
	var field [][]rune
	for _, line := range aocutils.SplitByteLines(data) {
		field = append(field, []rune(string(line)))
	}

	r1 := 0
	r2 := 0
	for y, line := range field {
		for x, letter := range line {
			if letter == 'X' {
				mDirs := findLetter(field, 'M', position{x, y}, direction{})
				if mDirs == nil {
					continue
				}

				for _, mDir := range mDirs {
					currentPos := position{x + mDir.x, y + mDir.y}
					aDirs := findLetter(field, 'A', currentPos, mDir)
					if aDirs == nil {
						continue
					}

					for _, aDir := range aDirs {
						currentPos := position{x + aDir.x*2, y + aDir.y*2}
						sDirs := findLetter(field, 'S', currentPos, aDir)
						if sDirs == nil {
							continue
						}

						r1++
					}
				}
			}

			if letter == 'A' {
				if InMAS(field, position{x, y}) {
					r2++
				}
			}
		}
	}

	return r1, r2
}

func main() {
	bytesData := aocutils.GetFileBytes("data/4")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
