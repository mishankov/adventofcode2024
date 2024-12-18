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

func TestSolveAsync(t *testing.T) {
	input := aocutils.GetFileBytes("../../data/2_test")
	expectedOne := 2
	expectedTwo := 4

	resultOne, resultTwo := solveAsync(input)

	if resultOne != expectedOne {
		t.Fatalf("Result expected to be %v, got %v", expectedOne, resultOne)
	}

	if resultTwo != expectedTwo {
		t.Fatalf("Result expected to be %v, got %v", expectedTwo, resultTwo)
	}
}

func BenchmarkSolve(b *testing.B) {
	input := aocutils.GetFileBytes("../../data/2_test")
	b.ResetTimer()

	for range b.N {
		solve(input)
	}
}

func BenchmarkSolveAsync(b *testing.B) {
	input := aocutils.GetFileBytes("../../data/2_test")
	b.ResetTimer()

	for range b.N {
		solveAsync(input)
	}
}
