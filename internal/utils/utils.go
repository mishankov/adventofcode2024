package utils

import (
	"io"
	"log"
	"os"
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

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
