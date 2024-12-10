package main

import (
	"aoc2024/pkg/aocutils"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aocutils.GetFileBytes("../../data/4_test")
	expectedOne := 18
	expectedTwo := 9

	resultOne, resultTwo := solve(input)

	if resultOne != expectedOne {
		t.Fatalf("Result expected to be %v, got %v", expectedOne, resultOne)
	}

	if resultTwo != expectedTwo {
		t.Fatalf("Result expected to be %v, got %v", expectedTwo, resultTwo)
	}
}
