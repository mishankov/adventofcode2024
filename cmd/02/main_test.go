package main

import (
	"aoc2024/pkg/aocutils"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aocutils.GetFileBytes("../../data/2_test")
	expectedOne := 2
	expectedTwo := 4

	resultOne, resultTwo := solve(input)

	if resultOne != expectedOne {
		t.Fatalf("Result expected to be %v, got %v", expectedOne, resultOne)
	}

	if resultTwo != expectedTwo {
		t.Fatalf("Result expected to be %v, got %v", expectedTwo, resultTwo)
	}
}
