package main

import (
	"aoc2024/pkg/aocutils"
	"log"
	"slices"
)

func highs(current int, pos aocutils.Position, field [][]int) ([]aocutils.Position, int) {
	if current == 9 {
		return []aocutils.Position{pos}, 1
	}

	maxY := len(field) - 1
	maxX := len(field[0]) - 1

	directions := []aocutils.Vector{
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: 0, Y: 1},
		{X: -1, Y: 0},
	}

	hs := []aocutils.Position{}
	rat := 0
	for _, dir := range directions {
		np := pos.Add(dir)

		if !np.IsValid(maxX, maxY) {
			continue
		}

		next := field[np.Y][np.X]
		if next-current != 1 {
			continue
		}

		nhs, nr := highs(field[np.Y][np.X], np, field)
		rat += nr

		for _, nh := range nhs {
			if !slices.Contains(hs, nh) {
				hs = append(hs, nh)
			}
		}
	}

	return hs, rat
}

func solve(data []byte) (int, int) {
	lines := aocutils.SplitByteLines(data)

	field := make([][]int, len(lines))

	for y, line := range lines {
		row := make([]int, len(line))
		for x, cell := range line {
			row[x] = aocutils.ToInt(cell)
		}

		field[y] = row
	}

	r1 := 0
	r2 := 0
	for y, row := range field {
		for x, cell := range row {
			if cell == 0 {
				hs, rat := highs(cell, aocutils.Position{X: x, Y: y}, field)
				r1 += len(hs)
				r2 += rat
			}
		}
	}

	return r1, r2
}

func main() {
	bytesData := aocutils.GetFileBytes("data/10")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
