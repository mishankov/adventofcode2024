package main

import (
	"aoc2024/pkg/aocutils"
	"log"
	"slices"
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

func countChecksum(input []int) int {
	r := 0
	for i, item := range input {
		if item < 0 {
			break
		}

		r += i * item
	}

	return r
}

func countChecksum2(input []int) int {
	r := 0
	for i, item := range input {
		if item < 0 {
			continue
		}

		r += i * item
	}

	return r
}

func layoutItemsToIntSlice(lis []layoutItem) ([]int, []int) {
	layoutInt := []int{}
	emptyPositions := []int{}

	for _, item := range lis {
		for range item.size {
			layoutInt = append(layoutInt, item.id)
			if item.kind == empty {
				emptyPositions = append(emptyPositions, len(layoutInt)-1)
			}
		}
	}

	return layoutInt, emptyPositions
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

	layoutInt, emptyPositions := layoutItemsToIntSlice(layout)

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

	r1 := countChecksum(layoutInt)

	itemPos = len(layout) - 1
	itemPosDecrement := 0
	for itemPos >= 0 {
		item := layout[itemPos]
		if item.kind != file {
			itemPosDecrement += 1
			itemPos = len(layout) - 1 - itemPosDecrement
			continue
		}

		for emptyPos, emptyItem := range layout {
			if emptyItem.kind != empty {
				continue
			}

			if emptyPos >= itemPos {
				continue
			}

			if emptyItem.size < item.size {
				continue
			}

			layout[emptyPos] = item
			layout = slices.Insert(layout, emptyPos+1, layoutItem{kind: empty, size: emptyItem.size - item.size, id: -1})
			layout[itemPos+1] = layoutItem{kind: empty, size: item.size, id: -1}

			break
		}

		itemPosDecrement += 1
		itemPos = len(layout) - 1 - itemPosDecrement
	}

	layoutInt, _ = layoutItemsToIntSlice(layout)
	r2 := countChecksum2(layoutInt)

	return r1, r2
}

func main() {
	bytesData := aocutils.GetFileBytes("data/9")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
