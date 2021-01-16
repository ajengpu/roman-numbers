package model

import (
	"fmt"
	"strconv"
	"strings"
)

func ExecuteCommad(cmd string) {
	cmdArr := strings.Split(cmd, " ")
	switch cmdArr[len(cmdArr)-1] {
	case "?":
		GetValueCommad(cmdArr)
		return
	case "Credits":
		SaveItemCommand(cmdArr)
		return
	default:
		if _, err := GetRomanNumber(cmdArr[len(cmdArr)-1]); err != nil {
			fmt.Printf("%v\n", ErrInvalidCommand)
			return
		}
		SaveParamCommand(cmdArr)
		return
	}
}

func SaveParamCommand(cmd []string) {
	if len(cmd) != 3 || cmd[1] != "is" {
		fmt.Printf("%v", ErrInvalidCommand)
		return
	}

	UpsertParam(cmd[0], cmd[len(cmd)-1])
}

func SaveItemCommand(cmd []string) {
	total, err := strconv.ParseFloat(cmd[len(cmd)-2], 32)
	if err != nil || cmd[len(cmd)-3] != "is" {
		fmt.Printf("%v", ErrInvalidCommand)
		return
	}

	romanNum, err := TransalateParamToRoman(cmd[:len(cmd)-4])
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	value, err := GetItemValueFromStatement(romanNum, float32(total))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	UpsertItem(cmd[len(cmd)-4], value)
}

func GetValueCommad(cmd []string) {
	if cmd[2] == "Credits" {
		GetCreditCommand(cmd)
		return
	}

	if cmd[0] != "how" || cmd[1] != "much" || cmd[2] != "is" {
		fmt.Printf("%v\n", ErrInvalidCommand)
		return
	}

	romanNum, err := TransalateParamToRoman(cmd[3 : len(cmd)-1])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	value, err := GetDecimalFromRoman(romanNum)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	for _, param := range cmd[3 : len(cmd)-1] {
		fmt.Printf(param + " ")
	}
	fmt.Printf("is %v\n", value)
}

func GetCreditCommand(cmd []string) {
	itemValue, err := GetItemValue(cmd[len(cmd)-2])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if cmd[0] != "how" || cmd[1] != "many" || cmd[3] != "is" {
		fmt.Printf("%v\n", ErrInvalidCommand)
		return
	}

	romanNum, err := TransalateParamToRoman(cmd[4 : len(cmd)-2])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	value, err := GetDecimalFromRoman(romanNum)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	for _, param := range cmd[4 : len(cmd)-1] {
		fmt.Printf(param + " ")
	}

	fmt.Printf("is %v Credits\n", float32(value)*itemValue)
}
