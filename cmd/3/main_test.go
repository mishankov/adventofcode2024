package main

import (
	"aoc2024/pkg/aocutils"
	"testing"
)

func TestSolveFirst(t *testing.T) {
	input := aocutils.GetFileBytes("../../data/3.1_test")
	expected := 161

	resultOne, _ := solve(input)

	if resultOne != expected {
		t.Fatalf("Result expected to be %v, got %v", resultOne, expected)
	}
}

func TestSolveSecond(t *testing.T) {
	input := aocutils.GetFileBytes("../../data/3.2_test")
	expected := 48

	_, resultTwo := solve(input)

	if resultTwo != expected {
		t.Fatalf("Result expected to be %v, got %v", resultTwo, expected)
	}
}
