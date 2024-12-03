package aocutils

import (
	"io"
	"log"
	"os"
	"strconv"
)

func GetFileBytes(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	bytesData, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file content:", err)
	}

	return bytesData
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
	string | []byte
}

func ToInt[V toIntInput](input V) int {
	s := string(input)

	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error converting %q to int: %v\n", input, err)
	}
	return num
}
