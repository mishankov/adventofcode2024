package main

import (
	"aoc2024/pkg/aocutils"
	"log"
)

type layoutKind int

const (
	file  layoutKind = 0
	empty layoutKind = 1
)

type layoutItem struct {
	kind layoutKind
	size int
	id   int
}

func solve(data []byte) (int, int) {
	layout := []layoutItem{}
	for index, digitByte := range data {
		digit := aocutils.ToInt(digitByte)
		if index%2 == 0 {
			layout = append(layout, layoutItem{kind: file, size: digit, id: index / 2})
		} else {
			layout = append(layout, layoutItem{kind: empty, size: digit, id: -1})
		}
	}

	layoutInt := []int{}
	emptyPositions := []int{}

	for _, item := range layout {
		for range item.size {
			layoutInt = append(layoutInt, item.id)
			if item.kind == empty {
				emptyPositions = append(emptyPositions, len(layoutInt)-1)
			}
		}
	}

	itemPos := len(layoutInt) - 1
	posIndex := 0
	for posIndex < len(emptyPositions) {
		breakOuter := false
		for {
			item := layoutInt[itemPos]
			pos := emptyPositions[posIndex]

			if itemPos > pos && item != -1 {
				layoutInt[pos] = item
				layoutInt[itemPos] = -2
				itemPos -= 1
				posIndex += 1
				break
			}

			if itemPos <= pos {
				breakOuter = true
			}

			if item == -1 {
				itemPos -= 1
				break
			}

			break
		}

		if breakOuter {
			break
		}
	}

	r1 := 0
	for i, item := range layoutInt {
		if item < 0 {
			break
		}

		r1 += i * item
	}

	return r1, 0
}

func main() {
	bytesData := aocutils.GetFileBytes("data/9")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
