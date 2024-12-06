package main

import (
	"aoc2024/pkg/aocutils"
	"log"
)

type cell int

const (
	cellFree    cell = 0
	cellBusy    cell = 1
	cellVisited cell = 2
	cellAddBusy cell = 3
)

type field struct {
	cells [][]cell
}

func (f field) cellAvailable(x, y int) (inbounds bool, available bool) {
	if y >= len(f.cells) || y < 0 || x >= len(f.cells[0]) || x < 0 {
		return false, true
	}

	return true, f.cells[y][x] == cellFree || f.cells[y][x] == cellVisited
}

func (f field) countVisited() int {
	count := 0
	for _, row := range f.cells {
		for _, cell := range row {
			if cell == cellVisited {
				count++
			}
		}
	}
	return count
}

func (f field) visitedCellsPositions() []pos {
	res := []pos{}
	for y, row := range f.cells {
		for x, cell := range row {
			if cell == cellVisited {
				res = append(res, pos{x, y})
			}
		}
	}
	return res
}

func (f field) reset() {
	for y, row := range f.cells {
		for x, cell := range row {
			if cell == cellVisited || cell == cellAddBusy {
				f.cells[y][x] = cellFree
			}
		}
	}
}

type pos struct {
	x, y int
}

type dir int

const (
	up    dir = 0
	down  dir = 1
	left  dir = 2
	right dir = 3
)

var nextDir map[dir]dir = map[dir]dir{
	up:    right,
	right: down,
	down:  left,
	left:  up,
}

type guard struct {
	position      pos
	startPosition pos
	direction     dir
}

func (g guard) moveUntilOut(field field) (loop bool) {
	visitedPos := map[pos]dir{}
	for {

		checkdir, ok := visitedPos[g.position]
		if ok && checkdir == g.direction {
			return true
		}

		nextPosition := pos{}
		switch g.direction {
		case up:
			nextPosition = pos{x: g.position.x, y: g.position.y - 1}
		case right:
			nextPosition = pos{x: g.position.x + 1, y: g.position.y}
		case down:
			nextPosition = pos{x: g.position.x, y: g.position.y + 1}
		case left:
			nextPosition = pos{x: g.position.x - 1, y: g.position.y}
		}

		inbounds, available := field.cellAvailable(nextPosition.x, nextPosition.y)
		if !inbounds {
			field.cells[g.position.y][g.position.x] = cellVisited
			break
		}

		if available {
			field.cells[g.position.y][g.position.x] = cellVisited
			visitedPos[g.position] = g.direction
			g.position = nextPosition
			continue
		}

		if !available {
			g.direction = nextDir[g.direction]
			continue
		}
	}

	return false
}

func solve(data []byte) (int, int) {
	lines := aocutils.SplitByteLines(data)

	field := field{}
	field.cells = make([][]cell, len(lines))

	guard := guard{}
	guard.direction = up

	for y, line := range lines {
		field.cells[y] = make([]cell, len(line))
		for x, cell := range line {
			switch cell {
			case byte('.'):
				field.cells[y][x] = cellFree
			case byte('#'):
				field.cells[y][x] = cellBusy
			case byte('^'):
				field.cells[y][x] = cellFree
				guard.position = pos{x, y}
				guard.startPosition = pos{x, y}
			}
		}
	}

	guard.moveUntilOut(field)
	r1 := field.countVisited()

	r2 := 0
	for _, pos := range field.visitedCellsPositions() {
		if pos.x == guard.startPosition.x && pos.y == guard.position.y {
			continue
		}

		field.reset()

		guard.position = guard.startPosition
		guard.direction = up

		field.cells[pos.y][pos.x] = cellAddBusy

		loop := guard.moveUntilOut(field)
		if loop {
			r2++
		}
	}

	return r1, r2
}

func main() {
	bytesData := aocutils.GetFileBytes("data/6")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
