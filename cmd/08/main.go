package main

import (
	"aoc2024/pkg/aocutils"
	"log"
)

func solve(data []byte) (int, int) {
	lines := aocutils.SplitByteLines(data)
	nodes := map[rune][]aocutils.Position{}

	for y, line := range lines {
		for x, cell := range []rune(string(line)) {
			if cell != '.' {
				nodes[cell] = append(nodes[cell], aocutils.Position{X: x, Y: y})
			}
		}
	}

	maxY := len(lines) - 1
	maxX := len(lines[0]) - 1

	antinodePrositions := []aocutils.Position{}
	for _, positions := range nodes {
		for i, first := range positions {
			for j, second := range positions {
				if i == j {
					continue
				}

				x := first.X + (first.X - second.X)
				y := first.Y + (first.Y - second.Y)

				if x < 0 || x > maxX || y < 0 || y > maxY {
					continue
				}

				continueOuter := false
				for _, pos := range antinodePrositions {
					if pos.X == x && pos.Y == y {
						continueOuter = true
						break
					}
				}
				if continueOuter {
					continue
				}

				antinodePrositions = append(antinodePrositions, aocutils.Position{X: x, Y: y})
			}
		}
	}

	antinodePrositions2 := []aocutils.Position{}
	for _, positions := range nodes {
		for i, first := range positions {
			for j, second := range positions {
				if i == j {
					continue
				}

				vectorAntenas := second.VectorFrom(first)

				for y, line := range lines {
					for x := range []rune(string(line)) {
						pos := aocutils.Position{X: x, Y: y}
						vectorPoint := second.VectorFrom(pos)

						if float32(vectorPoint.X)/float32(vectorAntenas.X) == float32(vectorPoint.Y)/float32(vectorAntenas.Y) {
							continueOuter := false
							for _, pos := range antinodePrositions2 {
								if pos.X == x && pos.Y == y {
									continueOuter = true
									break
								}
							}
							if continueOuter {
								continue
							}
							antinodePrositions2 = append(antinodePrositions2, pos)
						}
					}
				}
			}
		}
	}

	return len(antinodePrositions), len(antinodePrositions2)
}

func main() {
	bytesData := aocutils.GetFileBytes("data/8")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
