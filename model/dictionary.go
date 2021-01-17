package model

import (
	"fmt"
	"strconv"
	"strings"
)

func ExecuteCommad(cmd string) string {
	cmdArr := strings.Split(cmd, " ")
	switch cmdArr[len(cmdArr)-1] {
	case "?":
		return GetValueCommad(cmdArr)
	case "Credits":
		return SaveItemCommand(cmdArr)
	default:
		if _, err := GetRomanNumber(cmdArr[len(cmdArr)-1]); err != nil {
			return fmt.Sprintf("%v\n", ErrInvalidCommand)

		}
		return SaveParamCommand(cmdArr)
	}
}

func SaveParamCommand(cmd []string) string {
	if len(cmd) != 3 || cmd[1] != "is" {
		return fmt.Sprintf("%v\n", ErrInvalidCommand)
	}

	UpsertParam(cmd[0], cmd[len(cmd)-1])
	return ""
}

func SaveItemCommand(cmd []string) string {
	total, err := strconv.ParseFloat(cmd[len(cmd)-2], 32)
	if err != nil || cmd[len(cmd)-3] != "is" {
		return fmt.Sprintf("%v\n", ErrInvalidCommand)
	}

	romanNum, err := TransalateParamToRoman(cmd[:len(cmd)-4])
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	value, err := GetItemValueFromStatement(romanNum, float32(total))
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	UpsertItem(cmd[len(cmd)-4], value)
	return ""
}

func GetValueCommad(cmd []string) string {
	if cmd[2] == "Credits" {
		return GetCreditCommand(cmd)
	}

	if cmd[0] != "how" || cmd[1] != "much" || cmd[2] != "is" {
		return fmt.Sprintf("%v\n", ErrInvalidCommand)
	}

	romanNum, err := TransalateParamToRoman(cmd[3 : len(cmd)-1])
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	value, err := GetDecimalFromRoman(romanNum)
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	res := ""
	for _, param := range cmd[3 : len(cmd)-1] {
		res += fmt.Sprintf(param + " ")
	}
	return fmt.Sprintf(res+"is %v\n", value)
}

func GetCreditCommand(cmd []string) string {
	itemValue, err := GetItemValue(cmd[len(cmd)-2])
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	if cmd[0] != "how" || cmd[1] != "many" || cmd[3] != "is" {
		return fmt.Sprintf("%v\n", ErrInvalidCommand)
	}

	romanNum, err := TransalateParamToRoman(cmd[4 : len(cmd)-2])
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	value, err := GetDecimalFromRoman(romanNum)
	if err != nil {
		return fmt.Sprintf("%v\n", err)
	}

	res := ""
	for _, param := range cmd[4 : len(cmd)-1] {
		res += fmt.Sprintf(param + " ")
	}

	return fmt.Sprintf(res+"is %v Credits\n", float32(value)*itemValue)
}
