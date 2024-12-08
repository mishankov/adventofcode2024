package main

import (
	"fmt"
	"log"
	"os"
)

func createMain(day string) {
	content := fmt.Sprintf(`package main

import (
	"aoc2024/pkg/aocutils"
	"log"
)

func solve(data []byte) (int, int) {
	return 0, 0
}

func main() {
	bytesData := aocutils.GetFileBytes("data/%v")
	resultOne, resultTwo := solve(bytesData)

	log.Println("Result 1:", resultOne)
	log.Println("Result 2:", resultTwo)
}`, day)

	err := os.WriteFile(fmt.Sprintf("cmd/%v/main.go", day), []byte(content), os.ModePerm)
	if err != nil {
		log.Fatal("Error creating main:", err)
	}
}

func createTest(day string) {
	content := fmt.Sprintf(`package main

import (
	"aoc2024/pkg/aocutils"
	"testing"
)

func TestSolve(t *testing.T) {
	input := aocutils.GetFileBytes("../../data/%v_test")
	expectedOne := 0
	expectedTwo := 0

	resultOne, resultTwo := solve(input)

	if resultOne != expectedOne {
		t.Fatalf("Result expected to be %%v, got %%v", expectedOne, resultOne)
	}

	if resultTwo != expectedTwo {
		t.Fatalf("Result expected to be %%v, got %%v", expectedTwo, resultTwo)
	}
}`, day)

	err := os.WriteFile(fmt.Sprintf("cmd/%v/main_test.go", day), []byte(content), os.ModePerm)
	if err != nil {
		log.Fatal("Error creating test:", err)
	}
}

func createFiles(day string) {
	err := os.WriteFile(fmt.Sprintf("data/%v", day), nil, os.ModePerm)
	if err != nil {
		log.Fatal("Error creating input:", err)
	}

	err = os.WriteFile(fmt.Sprintf("data/%v_test", day), nil, os.ModePerm)
	if err != nil {
		log.Fatal("Error creating test input:", err)
	}
}

func main() {
	day := os.Args[1]

	os.Mkdir(fmt.Sprintf("cmd/%v", day), os.ModePerm)

	createMain(day)
	createTest(day)
	createFiles(day)
}
