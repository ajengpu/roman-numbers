package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Dictionary struct {
	Regex          string
	VariableNumber int
	Command        string
	ItemType       string
	Answer         string
}

var dictionary = map[string][]Dictionary{
	"param": []Dictionary{
		Dictionary{
			Regex:          "([\\w]+) is (I|V|X|L|C|D|M)",
			VariableNumber: 0,
			Command:        "SaveParam",
			Answer:         "",
		},
	},
	"item": []Dictionary{
		Dictionary{
			Regex:          "([\\w ]+)is ([\\d]+) Credits",
			VariableNumber: 1,
			Command:        "SaveItem",
			ItemType:       "Credits",
			Answer:         "",
		},
	},
	"question": []Dictionary{
		Dictionary{
			Regex:          "how much is ([\\w ]+\\?)",
			VariableNumber: 1,
			Command:        "GetValue",
			Answer:         "{p} is {v}",
		},
		Dictionary{
			Regex:          "how many Credits is ([\\w ]+\\?)",
			VariableNumber: 1,
			Command:        "GetValue",
			ItemType:       "Credits",
			Answer:         "{p} is {v} Credits",
		},
		Dictionary{
			Regex:          "Does ([\\w ]+) has more Credits than ([\\w ]+)\\?",
			VariableNumber: 2,
			Command:        "CompareValue",
			ItemType:       "Credits",
			Answer:         "{p} has {v} Credits than {p}",
		},
		Dictionary{
			Regex:          "Is ([\\w ]+) larger than ([\\w ]+)\\?",
			VariableNumber: 2,
			Command:        "CompareValue",
			Answer:         "{p} is {v} than {p}",
		},
	},
}

func ExecuteCommand(cmd string) string {
	cmdArr := strings.Split(cmd, " ")
	switch cmdArr[len(cmdArr)-1] {
	case "?":
		return GetAnswer("question", cmd, cmdArr)
	case "I", "V", "X", "L", "C", "D", "M":
		return GetAnswer("param", cmd, cmdArr)
	default:
		return GetAnswer("item", cmd, cmdArr)
	}
}

func GetAnswer(cmdType string, cmd string, cmdArr []string) string {
	for _, d := range dictionary[cmdType] {
		if match, _ := regexp.MatchString(d.Regex, cmd); match {
			input, err := d.GetCommandInput(cmdType, cmdArr)
			if err != nil {
				return fmt.Sprintf("%v\n", err)
			}

			output, err := CallCommand(d, input, commands[d.Command])
			if err != nil {
				return fmt.Sprintf("%v\n", err)
			}

			return d.PrintAnswer(input, output)
		}
	}

	return fmt.Sprintf("%v\n", ErrInvalidCommand)
}

func (d Dictionary) GetCommandInput(cmdType string, cmd []string) (CommandInput, error) {
	switch cmdType {
	case "param":
		return d.GetParamCommandInput(cmd)
	case "question":
		return d.GetQuestionCommandInput(cmd)
	default:
		return d.GetItemCommandInput(cmd)
	}
}

func (d Dictionary) GetParamCommandInput(cmd []string) (CommandInput, error) {
	input := CommandInput{}

	input.RomanNumber = cmd[len(cmd)-1]
	input.Param = cmd[0]

	return input, nil
}

func (d Dictionary) GetItemCommandInput(cmd []string) (CommandInput, error) {
	input := CommandInput{}

	variables := [][]string{}
	tempVar := []string{}
	isParam := false

	for _, c := range cmd {
		if val, err := strconv.ParseFloat(c, 32); err == nil {
			input.Value = float32(val)
		} else if _, err := GetParamValue(c); err == nil {
			isParam = true
			tempVar = append(tempVar, c)
		} else if isParam {
			input.ItemName = c
			variables = append(variables, tempVar)
			tempVar = []string{}
			isParam = false
		}
	}

	if len(input.ItemName) < 1 {
		return CommandInput{}, ErrInvalidCommand
	}

	if d.VariableNumber != len(variables) {
		return CommandInput{}, ErrInvalidCommand
	}

	if input.Value == 0 {
		return CommandInput{}, ErrInvalidCommand
	}

	input.Variables = variables

	return input, nil
}

func (d Dictionary) GetQuestionCommandInput(cmd []string) (CommandInput, error) {
	input := CommandInput{}

	variables := [][]string{}
	tempVar := []string{}
	isParam := false

	for _, c := range cmd {
		if _, err := GetParamValue(c); err == nil {
			isParam = true
			tempVar = append(tempVar, c)
		} else if _, err := GetItemValue(d.ItemType, c); err == nil {
			isParam = true
			tempVar = append(tempVar, c)
		} else if isParam {
			variables = append(variables, tempVar)
			tempVar = []string{}
			isParam = false
		}
	}

	if d.VariableNumber != len(variables) {
		return CommandInput{}, ErrInvalidCommand
	}

	input.Variables = variables

	return input, nil
}

func (d Dictionary) PrintAnswer(input CommandInput, output CommandOutput) string {
	answer := d.Answer
	for i := 0; i < d.VariableNumber; i++ {
		paramStr := ""
		for _, param := range input.Variables[i] {
			paramStr += fmt.Sprintf(param + " ")
		}
		paramStr = strings.TrimSuffix(paramStr, " ")
		answer = strings.Replace(answer, "{p}", paramStr, 1)
	}

	valStr := fmt.Sprintf("%v", output.Value)
	answer = strings.Replace(answer, "{v}", valStr, -1)

	if len(answer) > 0 {
		answer += "\n"
	}

	return answer
}
