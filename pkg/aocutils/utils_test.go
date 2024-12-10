package aocutils_test

import (
	"aoc2024/pkg/aocutils"
	"log"
	"testing"
)

func TestVectorFrom(t *testing.T) {
	tests := []struct {
		p      aocutils.Position
		o      aocutils.Position
		result aocutils.Vector
	}{
		{
			p:      aocutils.Position{X: 1, Y: 1},
			o:      aocutils.Position{X: 0, Y: 0},
			result: aocutils.Vector{X: 1, Y: 1},
		},
	}

	for _, test := range tests {
		result := test.p.VectorFrom(test.o)

		if result != test.result {
			log.Fatalf("Expected: %v. Got: %v\n", test.result, result)
		}
	}
}
