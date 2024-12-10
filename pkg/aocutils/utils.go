package aocutils

import (
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
)

type Vector struct {
	X, Y int
}

type Position struct {
	X, Y int
}

func (p Position) VectorFrom(other Position) Vector {
	return Vector{X: p.X - other.X, Y: p.Y - other.Y}
}

func (p Position) Add(vec Vector) Position {
	return Position{X: p.X + vec.X, Y: p.Y + vec.Y}
}

func (p Position) IsValid(maxX, maxY int) bool {
	return p.X >= 0 && p.X <= maxX && p.Y >= 0 && p.Y <= maxY
}

func GetFileBytes(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file content:", err)
	}

	return data
}

func SplitByteLines(data []byte) [][]byte {
	return bytes.Split(bytes.ReplaceAll(data, []byte{13, 10}, []byte{10}), []byte{10})
}

type absInput interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Abs[V absInput](i V) V {
	if i < 0 {
		return -i
	}
	return i
}

type toIntInput interface {
	string | []byte | byte
}

func ToInt[V toIntInput](input V) int {
	s := string(input)

	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error converting %q to int: %v\n", input, err)
	}
	return num
}
