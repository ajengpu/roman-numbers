package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	m "github.com/ajengpu/roman-numbers/model"
)

var jsonPath string

type TestCase struct {
	Input  []string `json:"input"`
	Output []string `json:"output"`
}

type JsonTest struct {
	TestData []TestCase `json:"test_data"`
}

type TestCases []TestCase

func GetTestCases(jsonPath string) (TestCases, error) {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		return TestCases{}, err
	}
	defer jsonFile.Close()

	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return TestCases{}, err
	}

	var testCases JsonTest
	json.Unmarshal(byteData, &testCases)

	return testCases.TestData, nil
}

func (testCases TestCases) Run() {
	for i, testCase := range testCases {
		fmt.Printf("Start Test Case : %v\n", i+1)
		fmt.Printf("\n--INPUT--\n")

		var sysOutput []string
		for _, input := range testCase.Input {
			res := strings.TrimSuffix(m.ExecuteCommad(input), "\n")
			if res != "" {
				sysOutput = append(sysOutput, res)
			}
			fmt.Println(input)
		}

		fmt.Printf("\n--EXPECTED OUTPUT--\n")
		for _, output := range testCase.Output {
			fmt.Println(output)
		}

		fmt.Printf("\n--SYSTEM OUTPUT--\n")
		flag := 0
		for i := 0; i < len(sysOutput); i++ {
			if testCase.Output[i] != sysOutput[i] {
				fmt.Println(sysOutput[i] + "(X)")
				flag = 1
			} else {
				fmt.Println(sysOutput[i] + "(V)")
			}
		}

		if flag == 0 {
			fmt.Printf("\n--CORRECT--\n\n")
		} else {
			fmt.Printf("\n--INCORRECT--\n\n")
		}
	}
}
