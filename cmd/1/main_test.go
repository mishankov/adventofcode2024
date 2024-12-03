package main

import (
	"aoc2024/pkg/aocutils"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aocutils.GetFileBytes("../../data/1_test")
	expectedOne := 11
	expectedTwo := 31

	resultOne, resultTwo := solve(input)

	if resultOne != expectedOne {
		t.Fatalf("Result expected to be %q, got %q", resultOne, expectedOne)
	}

	if resultTwo != expectedTwo {
		t.Fatalf("Result expected to be %q, got %q", resultTwo, expectedTwo)
	}
}
