package main

import (
	"aoc2024/pkg/aocutils"
	"bytes"
	"log"
	"strconv"
)

type cache struct {
	data map[int]map[int]int
	hits int
}

func (c *cache) set(iteration int, stone int, data int) {
	_, ok := c.data[iteration]
	if !ok {
		c.data[iteration] = map[int]int{}
	}

	c.data[iteration][stone] = data
}

func (c *cache) get(iteration int, stone int) (int, bool) {
	iterationCache := c.data[iteration]
	if iterationCache != nil {
		stoneCache, ok := iterationCache[stone]
		if ok {
			c.hits++
			return stoneCache, true
		}
	}

	return 0, false
}

func (c *cache) reset() {
	c.data = map[int]map[int]int{}
	c.hits = 0
}

func (c cache) entriesAmount() int {
	r := 0
	for _, stoneCache := range c.data {
		r += len(stoneCache)
	}

	return r
}

var CACHE cache = cache{data: map[int]map[int]int{}, hits: 0}

func processStones(stone int, iteration int, depth int) int {
	if iteration == depth {
		return 1
	}

	iteration++

	c, ok := CACHE.get(iteration, stone)
	if ok {
		return c
	}

	switch {
	case stone == 0:
		result := processStones(1, iteration, depth)
		CACHE.set(iteration, stone, result)

		return result
	case len(strconv.Itoa(stone))%2 == 0:
		numberAsString := strconv.Itoa(stone)
		middleIndex := len(numberAsString) / 2

		left := processStones(aocutils.ToInt(numberAsString[:middleIndex]), iteration, depth)
		right := processStones(aocutils.ToInt(numberAsString[middleIndex:]), iteration, depth)

		result := left + right
		CACHE.set(iteration, stone, result)

		return result
	default:
		result := processStones(stone*2024, iteration, depth)
		CACHE.set(iteration, stone, result)

		return result
	}
}

func solve(data []byte) (int, int) {
	splitedData := bytes.Split(data, []byte{32})

	numbers := make([]int, len(splitedData))
	for i, d := range splitedData {
		numbers[i] = aocutils.ToInt(d)
	}

	r1 := 0
	for _, num := range numbers {
		r1 += processStones(num, 0, 25)
	}
	log.Println("1. Cache hits:", CACHE.hits, "Cache entries:", CACHE.entriesAmount())

	CACHE.reset()

	r2 := 0
	for _, num := range numbers {
		r2 += processStones(num, 0, 75)
	}
	log.Println("2. Cache hits:", CACHE.hits, "Cache entries:", CACHE.entriesAmount())

	return r1, r2
}

func main() {
	bytesData := aocutils.GetFileBytes("data/11")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}
