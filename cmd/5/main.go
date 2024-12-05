package main

import (
	"aoc2024/pkg/aocutils"
	"log"
	"slices"
	"strings"
)

type order struct {
	left, right string
}

type update struct {
	pages     map[string]int
	pagesList []string
}

func (u update) correct(orders []order) bool {
	for _, order := range orders {
		left, okl := u.pages[order.left]
		right, okr := u.pages[order.right]

		if !okl || !okr {
			continue
		}

		if left > right {
			return false
		}
	}

	return true
}

func (u update) sort(orders []order) {
	slices.SortFunc(u.pagesList, func(p1, p2 string) int {
		for _, order := range orders {
			if p1 == order.left && p2 == order.right {
				return -1
			}

			if p1 == order.right && p2 == order.left {
				return 1
			}
		}

		return 0
	})
}

func solve(data []byte) (int, int) {
	lines := aocutils.SplitByteLines(data)
	r1 := 0
	r2 := 0

	orders := []order{}
	for _, line := range lines {
		lineStr := string(line)
		if strings.Contains(lineStr, "|") {
			pages := strings.Split(lineStr, "|")
			orders = append(orders, order{pages[0], pages[1]})
		}

		if strings.Contains(lineStr, ",") {
			pages := strings.Split(lineStr, ",")
			middlePageNumber := len(pages) / 2

			currentUpdate := update{pages: map[string]int{}, pagesList: []string{}}
			for number, page := range pages {
				currentUpdate.pages[page] = number
				currentUpdate.pagesList = append(currentUpdate.pagesList, page)
			}

			if currentUpdate.correct(orders) {
				r1 += aocutils.ToInt(currentUpdate.pagesList[middlePageNumber])
			} else {
				currentUpdate.sort(orders)
				r2 += aocutils.ToInt(currentUpdate.pagesList[middlePageNumber])
			}
		}
	}

	return r1, r2
}

func main() {
	bytesData := aocutils.GetFileBytes("data/5")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
