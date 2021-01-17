package main

import (
	"fmt"
	"os"

	t "github.com/ajengpu/roman-numbers/test/test-case"
)

var jsonPath = os.Args[1]

func main() {
	testCases, err := t.GetTestCases(jsonPath)
	if err != nil {
		fmt.Println("Cannot read test case data, please check your format.")
	}

	testCases.Run()
}
